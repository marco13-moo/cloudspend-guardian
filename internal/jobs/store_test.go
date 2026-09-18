package jobs

import (
	"testing"
	"time"
)

func TestEnqueueIsIdempotentAndLeaseCanBeReclaimed(t *testing.T) {
	store := NewStore()
	now := time.Date(2026, 9, 18, 10, 0, 0, 0, time.UTC)
	first, duplicate, err := store.Enqueue("ingest", "artifact-1", []byte("payload"), 3, now)
	if err != nil {
		t.Fatal(err)
	}
	second, wasDuplicate, err := store.Enqueue("ingest", "artifact-1", []byte("different"), 3, now.Add(time.Minute))
	if err != nil || !wasDuplicate || duplicate || second.ID != first.ID {
		t.Fatalf("expected idempotent enqueue: %+v %v %v", second, wasDuplicate, duplicate)
	}
	claimed, ok, err := store.Claim("worker-1", now, time.Minute)
	if err != nil || !ok || claimed.Attempts != 1 {
		t.Fatalf("expected first claim: %+v %v %v", claimed, ok, err)
	}
	reclaimed, ok, err := store.Claim("worker-2", now.Add(2*time.Minute), time.Minute)
	if err != nil || !ok || reclaimed.LeaseOwner != "worker-2" || reclaimed.Attempts != 2 {
		t.Fatalf("expected lease reclaim: %+v %v %v", reclaimed, ok, err)
	}
}

func TestFailuresRetryThenDeadLetterWithHistory(t *testing.T) {
	store := NewStore()
	now := time.Date(2026, 9, 18, 10, 0, 0, 0, time.UTC)
	job, _, _ := store.Enqueue("analyze", "key", nil, 2, now)
	claimed, _, _ := store.Claim("worker", now, time.Minute)
	failed, err := store.Fail(claimed.ID, "worker", now, Transient, "dependency timeout", time.Second)
	if err != nil || failed.State != Pending || failed.NextAttemptAt.IsZero() {
		t.Fatalf("expected retryable failure: %+v %v", failed, err)
	}
	claimed, _, _ = store.Claim("worker", now.Add(2*time.Second), time.Minute)
	dead, err := store.Fail(job.ID, "worker", now.Add(2*time.Second), Permanent, "malformed payload", 0)
	if err != nil || dead.State != DeadLetter || len(dead.History) != 4 {
		t.Fatalf("expected dead letter history: %+v %v", dead, err)
	}
}

func TestCompletionRequiresCurrentLease(t *testing.T) {
	store := NewStore()
	now := time.Date(2026, 9, 18, 10, 0, 0, 0, time.UTC)
	job, _, _ := store.Enqueue("ingest", "key", nil, 1, now)
	if _, err := store.Complete(job.ID, "worker", now); err == nil {
		t.Fatal("expected completion without lease to fail")
	}
}
