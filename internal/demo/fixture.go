// Package demo provides deterministic, explicitly synthetic demonstration data.
package demo

import (
	"fmt"
	"strings"
	"time"
)

const csvHeader = "BillingPeriodStart,BillingPeriodEnd,ChargePeriodStart,ChargePeriodEnd,ProviderName,BillingAccountId,ResourceId,ServiceName,BilledCost,Currency,UsageQuantity,UsageUnit,CPUUtilizationP95,MemoryUtilizationP95,SLOErrorBudgetRemaining\n"

// FixtureCSV returns a stable dataset with one underutilized and one healthy resource.
func FixtureCSV() string {
	var fixture strings.Builder
	fixture.WriteString(csvHeader)
	start := time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	for day := 0; day < 14; day++ {
		chargeStart := start.AddDate(0, 0, day)
		chargeEnd := chargeStart.AddDate(0, 0, 1)
		writeRow(&fixture, chargeStart, chargeEnd, "arn:synthetic:ec2:idle-1", "2.50", 4.2, 8.1, 0.95)
		writeRow(&fixture, chargeStart, chargeEnd, "arn:synthetic:ec2:active-1", "5.00", 62.0, 71.0, 0.90)
	}
	return fixture.String()
}

func writeRow(fixture *strings.Builder, start, end time.Time, resourceID, cost string, cpu, memory, errorBudget float64) {
	fmt.Fprintf(
		fixture,
		"2026-09-01T00:00:00Z,2026-10-01T00:00:00Z,%s,%s,AWS,synthetic-account,%s,Amazon EC2,%s,EUR,24,hours,%.2f,%.2f,%.2f\n",
		start.Format(time.RFC3339), end.Format(time.RFC3339), resourceID, cost, cpu, memory, errorBudget,
	)
}
