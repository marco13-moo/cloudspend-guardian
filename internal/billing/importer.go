// Package billing imports provider records into a deliberately small,
// FOCUS-compatible canonical model while preserving source provenance.
package billing

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/marco13-moo/cloudspend-guardian/internal/domain"
)

const SchemaVersion = "focus-subset/v1"

var requiredColumns = []string{
	"BillingPeriodStart", "BillingPeriodEnd", "ChargePeriodStart", "ChargePeriodEnd",
	"ProviderName", "BillingAccountId", "ResourceId", "ServiceName", "BilledCost",
	"Currency", "UsageQuantity", "UsageUnit", "CPUUtilizationP95",
	"MemoryUtilizationP95", "SLOErrorBudgetRemaining",
}

// Record contains the canonical fields required by the first vertical slice.
// It is a documented subset of FOCUS, not a claim of complete specification support.
type Record struct {
	ID                      string
	BillingPeriodStart      time.Time
	BillingPeriodEnd        time.Time
	ChargePeriodStart       time.Time
	ChargePeriodEnd         time.Time
	ProviderName            string
	BillingAccountID        string
	ResourceID              string
	ServiceName             string
	BilledCost              domain.Money
	UsageQuantity           float64
	UsageUnit               string
	CPUUtilizationP95       float64
	MemoryUtilizationP95    float64
	SLOErrorBudgetRemaining float64
}

// ImportReport is the authoritative result only after reconciliation succeeds.
type ImportReport struct {
	SchemaVersion    string
	SourceSHA256     string
	Records          []Record
	RecordCount      int
	ReconciledCost   domain.Money
	ReconciliationOK bool
}

// ImportCSV parses, validates, deduplicates, and reconciles a canonical CSV artifact.
func ImportCSV(reader io.Reader) (ImportReport, error) {
	raw, err := io.ReadAll(reader)
	if err != nil {
		return ImportReport{}, fmt.Errorf("read source artifact: %w", err)
	}
	if len(raw) == 0 {
		return ImportReport{}, errors.New("source artifact is empty")
	}

	digest := sha256.Sum256(raw)
	csvReader := csv.NewReader(bytes.NewReader(raw))
	csvReader.ReuseRecord = false
	header, err := csvReader.Read()
	if err != nil {
		return ImportReport{}, fmt.Errorf("read header: %w", err)
	}
	columns, err := validateHeader(header)
	if err != nil {
		return ImportReport{}, err
	}

	report := ImportReport{SchemaVersion: SchemaVersion, SourceSHA256: hex.EncodeToString(digest[:])}
	seen := make(map[string]struct{})
	var currency string

	for line := 2; ; line++ {
		row, readErr := csvReader.Read()
		if errors.Is(readErr, io.EOF) {
			break
		}
		if readErr != nil {
			return ImportReport{}, fmt.Errorf("read line %d: %w", line, readErr)
		}
		record, parseErr := parseRecord(row, columns)
		if parseErr != nil {
			return ImportReport{}, fmt.Errorf("line %d: %w", line, parseErr)
		}
		if _, duplicate := seen[record.ID]; duplicate {
			return ImportReport{}, fmt.Errorf("line %d: duplicate canonical record %s", line, record.ID)
		}
		seen[record.ID] = struct{}{}
		if currency == "" {
			currency = record.BilledCost.Currency
		}
		if record.BilledCost.Currency != currency {
			return ImportReport{}, fmt.Errorf("line %d: mixed currencies require an explicit conversion boundary", line)
		}
		if record.BilledCost.MinorUnits > math.MaxInt64-report.ReconciledCost.MinorUnits {
			return ImportReport{}, fmt.Errorf("line %d: reconciled cost exceeds supported integer range", line)
		}
		report.Records = append(report.Records, record)
		report.ReconciledCost.MinorUnits += record.BilledCost.MinorUnits
	}

	if len(report.Records) == 0 {
		return ImportReport{}, errors.New("source artifact contains no billing records")
	}
	report.RecordCount = len(report.Records)
	report.ReconciledCost.Currency = currency
	report.ReconciliationOK = true
	return report, nil
}

func validateHeader(header []string) (map[string]int, error) {
	columns := make(map[string]int, len(header))
	for index, name := range header {
		columns[strings.TrimSpace(name)] = index
	}
	for _, required := range requiredColumns {
		if _, exists := columns[required]; !exists {
			return nil, fmt.Errorf("required column %q is missing", required)
		}
	}
	return columns, nil
}

func parseRecord(row []string, columns map[string]int) (Record, error) {
	value := func(name string) (string, error) {
		index := columns[name]
		if index >= len(row) {
			return "", fmt.Errorf("column %q is absent from row", name)
		}
		result := strings.TrimSpace(row[index])
		if result == "" {
			return "", fmt.Errorf("column %q is empty", name)
		}
		return result, nil
	}
	parseTime := func(name string) (time.Time, error) {
		raw, err := value(name)
		if err != nil {
			return time.Time{}, err
		}
		parsed, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			return time.Time{}, fmt.Errorf("column %q: %w", name, err)
		}
		return parsed.UTC(), nil
	}
	parseFloat := func(name string) (float64, error) {
		raw, err := value(name)
		if err != nil {
			return 0, err
		}
		parsed, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return 0, fmt.Errorf("column %q: %w", name, err)
		}
		return parsed, nil
	}

	billingStart, err := parseTime("BillingPeriodStart")
	if err != nil {
		return Record{}, err
	}
	billingEnd, err := parseTime("BillingPeriodEnd")
	if err != nil {
		return Record{}, err
	}
	chargeStart, err := parseTime("ChargePeriodStart")
	if err != nil {
		return Record{}, err
	}
	chargeEnd, err := parseTime("ChargePeriodEnd")
	if err != nil {
		return Record{}, err
	}
	if !billingStart.Before(billingEnd) || !chargeStart.Before(chargeEnd) {
		return Record{}, errors.New("period start must precede period end")
	}

	provider, err := value("ProviderName")
	if err != nil {
		return Record{}, err
	}
	account, err := value("BillingAccountId")
	if err != nil {
		return Record{}, err
	}
	resource, err := value("ResourceId")
	if err != nil {
		return Record{}, err
	}
	service, err := value("ServiceName")
	if err != nil {
		return Record{}, err
	}
	costRaw, err := value("BilledCost")
	if err != nil {
		return Record{}, err
	}
	currency, err := value("Currency")
	if err != nil {
		return Record{}, err
	}
	cost, err := parseMinorUnits(costRaw, currency)
	if err != nil {
		return Record{}, fmt.Errorf("column %q: %w", "BilledCost", err)
	}
	usage, err := parseFloat("UsageQuantity")
	if err != nil {
		return Record{}, err
	}
	usageUnit, err := value("UsageUnit")
	if err != nil {
		return Record{}, err
	}
	cpu, err := parseFloat("CPUUtilizationP95")
	if err != nil {
		return Record{}, err
	}
	memory, err := parseFloat("MemoryUtilizationP95")
	if err != nil {
		return Record{}, err
	}
	errorBudget, err := parseFloat("SLOErrorBudgetRemaining")
	if err != nil {
		return Record{}, err
	}
	if usage < 0 || cpu < 0 || cpu > 100 || memory < 0 || memory > 100 || errorBudget < 0 || errorBudget > 1 {
		return Record{}, errors.New("usage and telemetry values are outside their valid ranges")
	}

	idMaterial := strings.Join([]string{provider, account, resource, service, chargeStart.Format(time.RFC3339), chargeEnd.Format(time.RFC3339)}, "\x00")
	idDigest := sha256.Sum256([]byte(idMaterial))
	return Record{
		ID:                      hex.EncodeToString(idDigest[:]),
		BillingPeriodStart:      billingStart,
		BillingPeriodEnd:        billingEnd,
		ChargePeriodStart:       chargeStart,
		ChargePeriodEnd:         chargeEnd,
		ProviderName:            provider,
		BillingAccountID:        account,
		ResourceID:              resource,
		ServiceName:             service,
		BilledCost:              cost,
		UsageQuantity:           usage,
		UsageUnit:               usageUnit,
		CPUUtilizationP95:       cpu,
		MemoryUtilizationP95:    memory,
		SLOErrorBudgetRemaining: errorBudget,
	}, nil
}

func parseMinorUnits(raw, currency string) (domain.Money, error) {
	if len(currency) != 3 || strings.ToUpper(currency) != currency {
		return domain.Money{}, fmt.Errorf("currency must be an uppercase ISO 4217 code")
	}
	parts := strings.Split(raw, ".")
	if len(parts) > 2 || len(parts) == 0 || parts[0] == "" {
		return domain.Money{}, fmt.Errorf("invalid decimal amount %q", raw)
	}
	whole, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil || whole < 0 {
		return domain.Money{}, fmt.Errorf("invalid non-negative decimal amount %q", raw)
	}
	fraction := ""
	if len(parts) == 2 {
		fraction = parts[1]
	}
	if len(fraction) > 2 {
		return domain.Money{}, fmt.Errorf("amount %q exceeds minor-unit precision", raw)
	}
	fraction += strings.Repeat("0", 2-len(fraction))
	minor, err := strconv.ParseInt(fraction, 10, 64)
	if err != nil {
		return domain.Money{}, fmt.Errorf("invalid decimal amount %q", raw)
	}
	if whole > (math.MaxInt64-minor)/100 {
		return domain.Money{}, fmt.Errorf("amount %q exceeds supported integer range", raw)
	}
	return domain.Money{Currency: currency, MinorUnits: whole*100 + minor}, nil
}
