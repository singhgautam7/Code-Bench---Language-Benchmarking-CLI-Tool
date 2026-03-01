package benchmark

import (
	"fmt"
	"os"
	"time"

	"github.com/gautam/codebench/internal/problems"
)

// RunConfig holds the configuration for a benchmark run.
type RunConfig struct {
	Problem   *problems.ProblemSpec
	Languages []string
	InputSize string
	Runs      int
	Warmup    int
	TimeoutMs int
}

// RunBenchmark executes the full benchmark pipeline for all specified languages:
// Docker build → warmup runs → measured runs → compute statistics.
func RunBenchmark(config RunConfig) ([]BenchmarkResult, error) {
	inputPath, err := problems.GetInputPath(config.Problem.Name, config.InputSize)
	if err != nil {
		return nil, err
	}

	var results []BenchmarkResult

	for _, lang := range config.Languages {
		fmt.Fprintf(os.Stderr, "\n━━━ %s ━━━\n", lang)

		result, err := benchmarkLanguage(config, lang, inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  ✗ Skipping %s: %v\n", lang, err)
			continue
		}

		results = append(results, *result)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no successful benchmarks completed")
	}

	return results, nil
}

// benchmarkLanguage runs the complete benchmark for a single language.
func benchmarkLanguage(config RunConfig, lang, inputPath string) (*BenchmarkResult, error) {
	// Resolve language directory
	langDir, err := problems.GetLanguageDir(config.Problem.Name, lang)
	if err != nil {
		return nil, err
	}

	// Step 1: Docker Build
	fmt.Fprintf(os.Stderr, "  ▸ Building Docker image...\n")
	buildTime, err := BuildImage(config.Problem.Name, lang, langDir)
	if err != nil {
		return nil, fmt.Errorf("build failed: %w", err)
	}
	fmt.Fprintf(os.Stderr, "  ✓ Build completed in %s\n", formatDuration(buildTime))

	imageName := fmt.Sprintf("codebench/%s/%s", config.Problem.Name, lang)

	// Step 2: Warmup Runs
	if config.Warmup > 0 {
		fmt.Fprintf(os.Stderr, "  ▸ Running %d warmup run(s)...\n", config.Warmup)
		for i := 0; i < config.Warmup; i++ {
			_, _, err := RunContainer(imageName, inputPath, config.TimeoutMs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "    ⚠ Warmup run %d failed: %v\n", i+1, err)
			}
		}
	}

	// Step 3: Measured Runs
	fmt.Fprintf(os.Stderr, "  ▸ Running %d measured run(s)...\n", config.Runs)
	var successTimes []time.Duration
	failures := 0
	timeouts := 0

	for i := 0; i < config.Runs; i++ {
		execTime, _, err := RunContainer(imageName, inputPath, config.TimeoutMs)
		if err != nil {
			if isTimeout(err) {
				timeouts++
				fmt.Fprintf(os.Stderr, "    ⚠ Run %d: TIMEOUT\n", i+1)
			} else {
				failures++
				fmt.Fprintf(os.Stderr, "    ⚠ Run %d failed: %v\n", i+1, err)
			}
			continue
		}
		successTimes = append(successTimes, execTime)
		fmt.Fprintf(os.Stderr, "    Run %d: %s\n", i+1, formatDuration(execTime))
	}

	// Step 4: Compute Statistics
	if len(successTimes) == 0 {
		return nil, fmt.Errorf("all %d runs failed (failures: %d, timeouts: %d)", config.Runs, failures, timeouts)
	}

	stats := ComputeStats(successTimes)

	return &BenchmarkResult{
		Language:    lang,
		CompileTime: buildTime,
		Stats:       stats,
		TotalRuns:   config.Runs,
		Successes:   len(successTimes),
		Failures:    failures,
		Timeouts:    timeouts,
	}, nil
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
