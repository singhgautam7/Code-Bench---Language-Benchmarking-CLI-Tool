package benchmark

import (
	"fmt"
	"sync"
	"time"

	"github.com/singhgautam7/Code-Bench---Language-Benchmarking-CLI-Tool/internal/problems"
)

// RunConfig holds the configuration for a benchmark run.
type RunConfig struct {
	Problem   *problems.ProblemSpec
	Languages []string
	InputSize string
	Runs      int
	Warmup    int
	TimeoutMs int
	Parallel  int
	Verbose   bool
}

// RunBenchmark executes the full benchmark pipeline for all specified languages.
// It supports parallel execution across languages, limited by config.Parallel.
func RunBenchmark(config RunConfig) (RunMetadata, []BenchmarkResult, error) {
	inputPath, err := problems.GetInputPath(config.Problem.Name, config.InputSize)
	if err != nil {
		return RunMetadata{}, nil, err
	}

	metadata := CaptureMetadata(config)

	var results []BenchmarkResult
	var resultsMutex sync.Mutex

	// Default to sequential if Parallel is not set
	parallels := config.Parallel
	if parallels <= 0 {
		parallels = 1
	}

	sem := make(chan struct{}, parallels)
	var wg sync.WaitGroup

	for _, lang := range config.Languages {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			sem <- struct{}{}        // Acquire token
			defer func() { <-sem }() // Release token

			result := benchmarkLanguage(config, l, inputPath)

			resultsMutex.Lock()
			results = append(results, result)
			resultsMutex.Unlock()
		}(lang)
	}

	wg.Wait()
	return metadata, results, nil
}

// benchmarkLanguage runs the complete benchmark for a single language sequentially.
func benchmarkLanguage(config RunConfig, lang, inputPath string) BenchmarkResult {
	result := BenchmarkResult{
		Language:  lang,
		TotalRuns: config.Runs,
	}

	langDir, err := problems.GetLanguageDir(config.Problem.Name, lang)
	if err != nil {
		result.Status = StatusBuildError
		return result
	}

	buildTime, err := BuildImage(config.Problem.Name, lang, langDir, config.Verbose)
	result.CompileTime = buildTime

	if err != nil {
		result.Status = StatusBuildError
		return result
	}

	imageName := fmt.Sprintf("codebench/%s/%s", config.Problem.Name, lang)
	result.Version = GetLanguageVersion(imageName, lang)

	// Warmup Runs
	if config.Warmup > 0 {
		for i := 0; i < config.Warmup; i++ {
			// Discard warmup returns
			RunContainer(imageName, inputPath, config.TimeoutMs, config.Verbose)
		}
	}

	// Measured Runs
	var successTimes []time.Duration
	var successMems []float64

	for i := 0; i < config.Runs; i++ {
		execTime, memMB, _, err := RunContainer(imageName, inputPath, config.TimeoutMs, config.Verbose)
		if err != nil {
			if isTimeout(err) {
				result.Timeouts++
			} else {
				result.Failures++
			}
			continue
		}
		successTimes = append(successTimes, execTime)
		successMems = append(successMems, memMB)
	}

	result.Successes = len(successTimes)

	if len(successTimes) == 0 {
		if result.Timeouts > 0 && result.Failures == 0 {
			result.Status = StatusTimeout
		} else if result.Failures > 0 && result.Timeouts == 0 {
			result.Status = StatusFailed
		} else {
			result.Status = StatusNoSuccessfulRuns
		}
		return result
	}

	result.Stats = ComputeStats(successTimes, successMems)
	result.Status = StatusOK

	return result
}

// isTimeout checks if an error is a timeout error.
func isTimeout(err error) bool {
	return err != nil && len(err.Error()) >= 7 && err.Error()[:7] == "TIMEOUT"
}

// formatDuration formats a duration as a human-readable string with millisecond precision.
func formatDuration(d time.Duration) string {
	ms := float64(d.Microseconds()) / 1000.0
	if ms >= 1000 {
		return fmt.Sprintf("%.2fs", ms/1000.0)
	}
	return fmt.Sprintf("%.2fms", ms)
}
