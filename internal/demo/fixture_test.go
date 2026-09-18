package demo

import (
	"crypto/sha256"
	"testing"
)

func TestFixtureCSVIsDeterministic(t *testing.T) {
	first := sha256.Sum256([]byte(FixtureCSV()))
	second := sha256.Sum256([]byte(FixtureCSV()))
	if first != second {
		t.Fatal("synthetic fixture changed between identical generations")
	}
}
