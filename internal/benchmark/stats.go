package benchmark

import (
	"sort"
	"time"
)

// Stats holds computed statistics for a set of benchmark runs.
type Stats struct {
	Median time.Duration
	Avg    time.Duration
	Min    time.Duration
	Max    time.Duration
}

// BenchmarkResult holds the complete results for a single language benchmark.
type BenchmarkResult struct {
	Language    string
	CompileTime time.Duration
	Stats       Stats
	TotalRuns   int
	Successes   int
	Failures    int
	Timeouts    int
}

// ComputeStats calculates min, max, average, and median from a slice of durations.
// The input slice must have at least one element. Returns zero Stats for empty input.
func ComputeStats(times []time.Duration) Stats {
	if len(times) == 0 {
		return Stats{}
	}

	// Sort a copy to avoid mutating the input
	sorted := make([]time.Duration, len(times))
	copy(sorted, times)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	// Min and Max
	min := sorted[0]
	max := sorted[len(sorted)-1]

	// Average
	var total time.Duration
	for _, t := range sorted {
		total += t
	}
	avg := total / time.Duration(len(sorted))

	// Median: for odd count, take middle element; for even count, average of two middle elements
	var median time.Duration
	n := len(sorted)
	if n%2 == 1 {
		median = sorted[n/2]
	} else {
		median = (sorted[n/2-1] + sorted[n/2]) / 2
	}

	return Stats{
		Median: median,
		Avg:    avg,
		Min:    min,
		Max:    max,
	}
}
