// Package jobs models durable asynchronous workflow state.
package jobs

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type State string

const (
	Pending    State = "pending"
	Running    State = "running"
	Succeeded  State = "succeeded"
	DeadLetter State = "dead_letter"
)

type FailureClass string

const (
	Transient FailureClass = "transient"
	RateLimit FailureClass = "rate_limited"
	Permanent FailureClass = "permanent"
	Conflict  FailureClass = "conflict"
	Operator  FailureClass = "operator_actionable"
)

type Job struct {
	ID             string
	Kind           string
	IdempotencyKey string
	Payload        []byte
	State          State
	Attempts       int
	MaxAttempts    int
	LeaseOwner     string
	LeaseUntil     time.Time
	NextAttemptAt  time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Failure        string
	History        []Transition
}

type Transition struct {
	From    State
	To      State
	At      time.Time
	Actor   string
	Reason  string
	Attempt int
}

type Store struct {
	mu     sync.Mutex
	jobs   map[string]Job
	byKey  map[string]string
	nextID int
}

func NewStore() *Store {
	return &Store{jobs: make(map[string]Job), byKey: make(map[string]string)}
}

func (s *Store) Enqueue(kind, key string, payload []byte, maxAttempts int, now time.Time) (Job, bool, error) {
	if kind == "" || key == "" || maxAttempts < 1 || now.IsZero() {
		return Job{}, false, errors.New("kind, idempotency key, positive max attempts, and timestamp are required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if id, ok := s.byKey[key]; ok {
		return clone(s.jobs[id]), true, nil
	}
	s.nextID++
	id := fmt.Sprintf("job-%d", s.nextID)
	job := Job{ID: id, Kind: kind, IdempotencyKey: key, Payload: append([]byte(nil), payload...), State: Pending, MaxAttempts: maxAttempts, CreatedAt: now.UTC(), UpdatedAt: now.UTC()}
	s.jobs[id], s.byKey[key] = job, id
	return clone(job), false, nil
}

func (s *Store) Claim(worker string, now time.Time, lease time.Duration) (Job, bool, error) {
	if worker == "" || now.IsZero() || lease <= 0 {
		return Job{}, false, errors.New("worker, timestamp, and positive lease are required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, job := range s.jobs {
		if job.State == Running && !job.LeaseUntil.After(now) {
			job.State, job.LeaseOwner, job.LeaseUntil = Pending, "", time.Time{}
			job.UpdatedAt = now.UTC()
			s.jobs[id] = job
		}
	}
	for id, job := range s.jobs {
		if job.State != Pending || job.NextAttemptAt.After(now) {
			continue
		}
		job.State, job.LeaseOwner, job.LeaseUntil = Running, worker, now.UTC().Add(lease)
		job.Attempts++
		job.UpdatedAt = now.UTC()
		job.History = append(job.History, Transition{From: Pending, To: Running, At: now.UTC(), Actor: worker, Attempt: job.Attempts})
		s.jobs[id] = job
		return clone(job), true, nil
	}
	return Job{}, false, nil
}

func (s *Store) Complete(id, worker string, now time.Time) (Job, error) {
	return s.transition(id, worker, now, Succeeded, "")
}

func (s *Store) Fail(id, worker string, now time.Time, class FailureClass, reason string, backoff time.Duration) (Job, error) {
	if reason == "" || backoff < 0 {
		return Job{}, errors.New("failure reason and non-negative backoff are required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	job, ok := s.jobs[id]
	if !ok || job.State != Running || job.LeaseOwner != worker {
		return Job{}, errors.New("job is not leased by worker")
	}
	job.Failure = string(class) + ": " + reason
	if class == Permanent || class == Conflict || class == Operator || job.Attempts >= job.MaxAttempts {
		job.State = DeadLetter
		job.History = append(job.History, Transition{From: Running, To: DeadLetter, At: now.UTC(), Actor: worker, Reason: job.Failure, Attempt: job.Attempts})
	} else {
		job.State, job.LeaseOwner, job.LeaseUntil = Pending, "", time.Time{}
		job.NextAttemptAt = now.UTC().Add(backoff)
		job.History = append(job.History, Transition{From: Running, To: Pending, At: now.UTC(), Actor: worker, Reason: job.Failure, Attempt: job.Attempts})
	}
	job.UpdatedAt = now.UTC()
	s.jobs[id] = job
	return clone(job), nil
}

func (s *Store) Get(id string) (Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	job, ok := s.jobs[id]
	return clone(job), ok
}

func (s *Store) transition(id, worker string, now time.Time, target State, reason string) (Job, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	job, ok := s.jobs[id]
	if !ok || job.State != Running || job.LeaseOwner != worker {
		return Job{}, errors.New("job is not leased by worker")
	}
	job.State, job.LeaseOwner, job.LeaseUntil, job.UpdatedAt = target, "", time.Time{}, now.UTC()
	job.History = append(job.History, Transition{From: Running, To: target, At: now.UTC(), Actor: worker, Reason: reason, Attempt: job.Attempts})
	s.jobs[id] = job
	return clone(job), nil
}

func clone(job Job) Job {
	job.Payload = append([]byte(nil), job.Payload...)
	job.History = append([]Transition(nil), job.History...)
	return job
}
