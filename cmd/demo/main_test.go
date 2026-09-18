package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestRunProducesReconciledSyntheticDemonstration(t *testing.T) {
	var destination bytes.Buffer
	if err := run(&destination); err != nil {
		t.Fatalf("run demonstration: %v", err)
	}

	var result output
	if err := json.Unmarshal(destination.Bytes(), &result); err != nil {
		t.Fatalf("decode demonstration output: %v", err)
	}
	if !result.Synthetic || !result.Reconciled || result.RecordCount != 28 {
		t.Fatalf("unexpected demonstration metadata: %+v", result)
	}
	if len(result.Evaluation.Recommendations) != 1 || len(result.Evaluation.Abstentions) != 1 {
		t.Fatalf("unexpected demonstration evaluation: %+v", result.Evaluation)
	}
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, bytes.ErrTooLarge
}

func TestRunReportsEncodingFailure(t *testing.T) {
	if err := run(failingWriter{}); err == nil {
		t.Fatal("expected output failure to be reported")
	}
}
