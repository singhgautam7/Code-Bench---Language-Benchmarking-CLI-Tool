package export

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// WriteCSV serializes the benchmark results and metadata into a flattened CSV format.
func WriteCSV(report Report, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create export file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write Headers
	headers := []string{
		"Problem", "Input", "Timestamp", "OS", "Arch", "CPUs", "DockerVersion",
		"Language", "Version", "Status", "CompileTime(ms)",
		"MedianTime(ms)", "AvgTime(ms)", "MinTime(ms)", "MaxTime(ms)",
		"MemMedian(MB)", "MemAvg(MB)", "MemMin(MB)", "MemMax(MB)",
		"TotalRuns", "Successes", "Failures", "Timeouts",
	}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}

	m := report.Metadata

	// Write Rows
	for _, r := range report.Results {
		row := []string{
			m.Problem,
			m.Input,
			m.Timestamp,
			m.HostOS,
			m.HostArch,
			strconv.Itoa(m.HostCPUs),
			m.DockerVersion,

			r.Language,
			r.Version,
			r.Status,
			fmt.Sprintf("%.2f", ms(r.CompileTime)),

			fmt.Sprintf("%.2f", ms(r.Stats.Median)),
			fmt.Sprintf("%.2f", ms(r.Stats.Avg)),
			fmt.Sprintf("%.2f", ms(r.Stats.Min)),
			fmt.Sprintf("%.2f", ms(r.Stats.Max)),

			fmt.Sprintf("%.2f", r.Stats.MemMedian),
			fmt.Sprintf("%.2f", r.Stats.MemAvg),
			fmt.Sprintf("%.2f", r.Stats.MemMin),
			fmt.Sprintf("%.2f", r.Stats.MemMax),

			strconv.Itoa(r.TotalRuns),
			strconv.Itoa(r.Successes),
			strconv.Itoa(r.Failures),
			strconv.Itoa(r.Timeouts),
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	return nil
}

func ms(d time.Duration) float64 {
	return float64(d.Microseconds()) / 1000.0
}
