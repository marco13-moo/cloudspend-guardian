// Command demo executes the first deterministic CloudSpend Guardian vertical slice.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/billing"
	"github.com/marco13-moo/cloudspend-guardian/internal/demo"
	"github.com/marco13-moo/cloudspend-guardian/internal/recommendation"
)

type output struct {
	Synthetic        bool                      `json:"synthetic"`
	SchemaVersion    string                    `json:"schemaVersion"`
	SourceSHA256     string                    `json:"sourceSha256"`
	RecordCount      int                       `json:"recordCount"`
	Reconciled       bool                      `json:"reconciled"`
	ReconciledAmount int64                     `json:"reconciledAmountMinorUnits"`
	Currency         string                    `json:"currency"`
	Evaluation       recommendation.Evaluation `json:"evaluation"`
}

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(destination io.Writer) error {
	report, err := billing.ImportCSV(strings.NewReader(demo.FixtureCSV()))
	if err != nil {
		return fmt.Errorf("import synthetic fixture: %w", err)
	}
	evaluation, err := recommendation.EvaluateIdleCompute(report, time.Date(2026, 9, 15, 12, 0, 0, 0, time.UTC))
	if err != nil {
		return fmt.Errorf("evaluate synthetic fixture: %w", err)
	}

	encoder := json.NewEncoder(destination)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(output{
		Synthetic:        true,
		SchemaVersion:    report.SchemaVersion,
		SourceSHA256:     report.SourceSHA256,
		RecordCount:      report.RecordCount,
		Reconciled:       report.ReconciliationOK,
		ReconciledAmount: report.ReconciledCost.MinorUnits,
		Currency:         report.ReconciledCost.Currency,
		Evaluation:       evaluation,
	}); err != nil {
		return fmt.Errorf("encode demonstration output: %w", err)
	}
	return nil
}
