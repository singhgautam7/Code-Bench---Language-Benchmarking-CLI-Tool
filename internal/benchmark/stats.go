package benchmark

import (
	"sort"
	"time"
)

// Status constants for benchmark results.
const (
	StatusOK               = "OK"
	StatusFailed           = "FAILED"
	StatusTimeout          = "TIMEOUT"
	StatusBuildError       = "BUILD ERROR"
	StatusNoSuccessfulRuns = "NO SUCCESSFUL RUNS"
)

// Stats holds computed statistics for a set of benchmark runs.
type Stats struct {
	Median    time.Duration
	Avg       time.Duration
	Min       time.Duration
	Max       time.Duration
	MemMedian float64 // in MB
	MemAvg    float64 // in MB
	MemMin    float64 // in MB
	MemMax    float64 // in MB
}

// BenchmarkResult holds the complete results for a single language benchmark.
type BenchmarkResult struct {
	Language    string
	Version     string
	Status      string
	CompileTime time.Duration
	Stats       Stats
	TotalRuns   int
	Successes   int
	Failures    int
	Timeouts    int
}

// ComputeStats calculates min, max, average, and median for both time and memory.
// It assumes times and mems are parallel slices of the same length.
func ComputeStats(times []time.Duration, mems []float64) Stats {
	if len(times) == 0 {
		return Stats{}
	}

	// Calculate time stats
	tSorted := make([]time.Duration, len(times))
	copy(tSorted, times)
	sort.Slice(tSorted, func(i, j int) bool { return tSorted[i] < tSorted[j] })

	var tTotal time.Duration
	for _, t := range tSorted {
		tTotal += t
	}
	tAvg := tTotal / time.Duration(len(tSorted))

	var tMedian time.Duration
	n := len(tSorted)
	if n%2 == 1 {
		tMedian = tSorted[n/2]
	} else {
		tMedian = (tSorted[n/2-1] + tSorted[n/2]) / 2
	}

	// Calculate mem stats if available
	var mMedian, mAvg, mMin, mMax float64
	if len(mems) > 0 && len(mems) == len(times) {
		mSorted := make([]float64, len(mems))
		copy(mSorted, mems)
		sort.Float64s(mSorted)

		mMin = mSorted[0]
		mMax = mSorted[len(mSorted)-1]

		var mTotal float64
		for _, m := range mSorted {
			mTotal += m
		}
		mAvg = mTotal / float64(len(mSorted))

		if n%2 == 1 {
			mMedian = mSorted[n/2]
		} else {
			mMedian = (mSorted[n/2-1] + mSorted[n/2]) / 2.0
		}
	}

	return Stats{
		Median:    tMedian,
		Avg:       tAvg,
		Min:       tSorted[0],
		Max:       tSorted[len(tSorted)-1],
		MemMedian: mMedian,
		MemAvg:    mAvg,
		MemMin:    mMin,
		MemMax:    mMax,
	}
}
