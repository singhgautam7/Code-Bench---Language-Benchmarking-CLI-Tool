package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/gautam/codebench/internal/benchmark"
	"github.com/gautam/codebench/internal/problems"
)

var (
	flagLangs     string
	flagInput     string
	flagRuns      int
	flagWarmup    int
	flagTimeoutMs int
)

var runCmd = &cobra.Command{
	Use:   "run <problem>",
	Short: "Run benchmarks for a problem across languages",
	Long: `Run benchmarks for a specified problem across one or more languages.
Each language implementation is built and executed inside a Docker container
for isolation and reproducibility.`,
	Args: cobra.ExactArgs(1),
	RunE: runBenchmark,
}

func init() {
	runCmd.Flags().StringVar(&flagLangs, "langs", "", "Comma-separated list of languages (default: all from problem.yaml)")
	runCmd.Flags().StringVar(&flagInput, "input", "small", "Input size: small, medium, or large")
	runCmd.Flags().IntVar(&flagRuns, "runs", 5, "Number of measured runs per language")
	runCmd.Flags().IntVar(&flagWarmup, "warmup", 1, "Number of warmup runs per language")
	runCmd.Flags().IntVar(&flagTimeoutMs, "timeout-ms", 0, "Timeout per run in milliseconds (default: from problem.yaml)")

	rootCmd.AddCommand(runCmd)
}

func runBenchmark(cmd *cobra.Command, args []string) error {
	problemName := args[0]

	// Load problem spec first (gives useful errors even without Docker)
	spec, err := problems.LoadProblem(problemName)
	if err != nil {
		return err
	}

	// Determine languages
	langs := spec.Languages
	if flagLangs != "" {
		langs = strings.Split(flagLangs, ",")
		// Trim whitespace
		for i := range langs {
			langs[i] = strings.TrimSpace(langs[i])
		}
	}

	if len(langs) == 0 {
		return fmt.Errorf("no languages specified")
	}

	// Check Docker availability (after validation so config errors show first)
	if err := benchmark.CheckDockerAvailable(); err != nil {
		return err
	}

	// Determine timeout
	timeoutMs := spec.TimeoutMs
	if flagTimeoutMs > 0 {
		timeoutMs = flagTimeoutMs
	}
	if timeoutMs <= 0 {
		timeoutMs = 2000 // default fallback
	}

	// Print configuration
	fmt.Fprintf(os.Stderr, "╔══════════════════════════════════════════╗\n")
	fmt.Fprintf(os.Stderr, "║          CodeBench — Phase 1             ║\n")
	fmt.Fprintf(os.Stderr, "╚══════════════════════════════════════════╝\n")
	fmt.Fprintf(os.Stderr, "  Problem:   %s\n", spec.Name)
	fmt.Fprintf(os.Stderr, "  Languages: %s\n", strings.Join(langs, ", "))
	fmt.Fprintf(os.Stderr, "  Input:     %s\n", flagInput)
	fmt.Fprintf(os.Stderr, "  Runs:      %d (warmup: %d)\n", flagRuns, flagWarmup)
	fmt.Fprintf(os.Stderr, "  Timeout:   %dms\n", timeoutMs)

	// Run benchmarks
	config := benchmark.RunConfig{
		Problem:   spec,
		Languages: langs,
		InputSize: flagInput,
		Runs:      flagRuns,
		Warmup:    flagWarmup,
		TimeoutMs: timeoutMs,
	}

	results, err := benchmark.RunBenchmark(config)
	if err != nil {
		return err
	}

	// Sort results by median execution time (ascending)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Stats.Median < results[j].Stats.Median
	})

	// Display results table
	fmt.Fprintln(os.Stderr)
	printResultsTable(results)

	return nil
}

// printResultsTable renders benchmark results as a formatted table.
func printResultsTable(results []benchmark.BenchmarkResult) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.SetTitle("Benchmark Results (sorted by median execution time)")

	t.AppendHeader(table.Row{
		"Language", "Compile (ms)", "Median (ms)", "Avg (ms)", "Min (ms)", "Max (ms)", "Runs",
	})

	for _, r := range results {
		status := fmt.Sprintf("%d/%d", r.Successes, r.TotalRuns)
		if r.Timeouts > 0 {
			status += fmt.Sprintf(" (%d timeout)", r.Timeouts)
		}
		if r.Failures > 0 {
			status += fmt.Sprintf(" (%d failed)", r.Failures)
		}

		t.AppendRow(table.Row{
			r.Language,
			fmtMs(r.CompileTime),
			fmtMs(r.Stats.Median),
			fmtMs(r.Stats.Avg),
			fmtMs(r.Stats.Min),
			fmtMs(r.Stats.Max),
			status,
		})
	}

	t.SetStyle(table.StyleRounded)
	t.Style().Title.Align = text.AlignCenter
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 2, Align: text.AlignRight},
		{Number: 3, Align: text.AlignRight},
		{Number: 4, Align: text.AlignRight},
		{Number: 5, Align: text.AlignRight},
		{Number: 6, Align: text.AlignRight},
	})

	t.Render()
}

// fmtMs formats a duration as milliseconds with 2 decimal places.
func fmtMs(d time.Duration) string {
	return fmt.Sprintf("%.2f", float64(d.Microseconds())/1000.0)
}
