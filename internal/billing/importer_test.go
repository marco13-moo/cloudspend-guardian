package billing

import (
	"strings"
	"testing"
)

const header = "BillingPeriodStart,BillingPeriodEnd,ChargePeriodStart,ChargePeriodEnd,ProviderName,BillingAccountId,ResourceId,ServiceName,BilledCost,Currency,UsageQuantity,UsageUnit,CPUUtilizationP95,MemoryUtilizationP95,SLOErrorBudgetRemaining\n"

const row = "2026-09-01T00:00:00Z,2026-10-01T00:00:00Z,2026-09-01T00:00:00Z,2026-09-02T00:00:00Z,AWS,synthetic-account,arn:synthetic:ec2:idle-1,Amazon EC2,2.50,EUR,24,hours,4.2,8.1,0.95\n"

func TestImportCSVReconcilesDeterministically(t *testing.T) {
	first, err := ImportCSV(strings.NewReader(header + row))
	if err != nil {
		t.Fatalf("unexpected import error: %v", err)
	}
	second, err := ImportCSV(strings.NewReader(header + row))
	if err != nil {
		t.Fatalf("unexpected replay error: %v", err)
	}
	if !first.ReconciliationOK || first.RecordCount != 1 || first.ReconciledCost.MinorUnits != 250 {
		t.Fatalf("unexpected report: %+v", first)
	}
	if first.SourceSHA256 != second.SourceSHA256 || first.Records[0].ID != second.Records[0].ID {
		t.Fatal("identical input did not produce deterministic provenance")
	}
}

func TestImportCSVRejectsDuplicateCanonicalRecord(t *testing.T) {
	_, err := ImportCSV(strings.NewReader(header + row + row))
	if err == nil || !strings.Contains(err.Error(), "duplicate canonical record") {
		t.Fatalf("expected duplicate error, got %v", err)
	}
}

func TestImportCSVRejectsLossyMoney(t *testing.T) {
	invalid := strings.Replace(row, "2.50,EUR", "2.505,EUR", 1)
	_, err := ImportCSV(strings.NewReader(header + invalid))
	if err == nil || !strings.Contains(err.Error(), "minor-unit precision") {
		t.Fatalf("expected precision error, got %v", err)
	}
}

func TestImportCSVRejectsMissingColumn(t *testing.T) {
	_, err := ImportCSV(strings.NewReader("ProviderName\nAWS\n"))
	if err == nil || !strings.Contains(err.Error(), "required column") {
		t.Fatalf("expected missing-column error, got %v", err)
	}
}

func TestImportCSVRejectsOutOfRangeTelemetry(t *testing.T) {
	invalid := strings.Replace(row, ",4.2,8.1,0.95", ",104.2,8.1,0.95", 1)
	_, err := ImportCSV(strings.NewReader(header + invalid))
	if err == nil || !strings.Contains(err.Error(), "outside their valid ranges") {
		t.Fatalf("expected telemetry range error, got %v", err)
	}
}

func TestImportCSVRejectsCostOverflow(t *testing.T) {
	invalid := strings.Replace(row, "2.50,EUR", "92233720368547759.00,EUR", 1)
	_, err := ImportCSV(strings.NewReader(header + invalid))
	if err == nil || !strings.Contains(err.Error(), "integer range") {
		t.Fatalf("expected integer-range error, got %v", err)
	}
}
