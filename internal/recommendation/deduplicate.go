package recommendation

import (
	"fmt"
	"sort"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

// Candidate identifies a provider or native recommendation as one opportunity.
type Candidate struct {
	Organization   string
	ResourceID     string
	ActionClass    string
	TargetState    string
	WindowStart    time.Time
	WindowEnd      time.Time
	Source         string
	Recommendation domain.Recommendation
}

// Opportunity is one canonical action with traceable corroborating sources.
type Opportunity struct {
	Key            string
	Recommendation domain.Recommendation
	Sources        []string
	SourceIDs      []string
}

func Deduplicate(candidates []Candidate) ([]Opportunity, error) {
	groups := make(map[string][]Candidate)
	for _, candidate := range candidates {
		if candidate.Organization == "" || candidate.ResourceID == "" || candidate.ActionClass == "" || candidate.TargetState == "" || candidate.Source == "" || !candidate.WindowStart.Before(candidate.WindowEnd) {
			return nil, fmt.Errorf("candidate has incomplete opportunity identity")
		}
		key := fmt.Sprintf("%s\x00%s\x00%s\x00%s\x00%s\x00%s", candidate.Organization, candidate.ResourceID, candidate.ActionClass, candidate.TargetState, candidate.WindowStart.UTC().Format(time.RFC3339), candidate.WindowEnd.UTC().Format(time.RFC3339))
		groups[key] = append(groups[key], candidate)
	}
	keys := make([]string, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]Opportunity, 0, len(keys))
	for _, key := range keys {
		group := groups[key]
		merged := group[0].Recommendation
		merged.Evidence = append([]domain.Evidence(nil), merged.Evidence...)
		sources, ids := make([]string, 0, len(group)), make([]string, 0, len(group))
		for index, candidate := range group {
			sources = append(sources, candidate.Source)
			ids = append(ids, candidate.Recommendation.ID)
			if index == 0 {
				continue
			}
			if candidate.Recommendation.ProjectedMonthlySavings.MinorUnits < merged.ProjectedMonthlySavings.MinorUnits {
				merged.ProjectedMonthlySavings = candidate.Recommendation.ProjectedMonthlySavings
			}
			merged.Evidence = append(merged.Evidence, candidate.Recommendation.Evidence...)
		}
		result = append(result, Opportunity{Key: key, Recommendation: merged, Sources: sources, SourceIDs: ids})
	}
	return result, nil
}
