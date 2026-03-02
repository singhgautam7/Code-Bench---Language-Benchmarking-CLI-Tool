package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/singhgautam7/Code-Bench---Language-Benchmarking-CLI-Tool/internal/benchmark"
	"github.com/singhgautam7/Code-Bench---Language-Benchmarking-CLI-Tool/internal/export"
	"github.com/singhgautam7/Code-Bench---Language-Benchmarking-CLI-Tool/internal/problems"
	"github.com/singhgautam7/Code-Bench---Language-Benchmarking-CLI-Tool/internal/ui"
)

var (
	flagLangs     string
	flagInput     string
	flagRuns      int
	flagWarmup    int
	flagTimeoutMs int
	flagParallel  int
	flagExport    string
	flagVerbose   bool
	flagNoSpinner bool
)

const Version = "0.2.0"

var runCmd = &cobra.Command{
	Use:   "run <problem>",
	Short: "Run benchmarks for a problem across languages",
	Long: `Run benchmarks for a specified problem across one or more languages.
Each language implementation is built and executed inside a Docker container
for isolation and reproducibility.`,
	Args:    cobra.ExactArgs(1),
	RunE:    runBenchmark,
}

func init() {
	runCmd.Flags().StringVar(&flagLangs, "langs", "", "Comma-separated list of languages (default: all from problem.yaml)")
	runCmd.Flags().StringVar(&flagInput, "input", "small", "Input size: small, medium, or large")
	runCmd.Flags().IntVar(&flagRuns, "runs", 5, "Number of measured runs per language")
	runCmd.Flags().IntVar(&flagWarmup, "warmup", 1, "Number of warmup runs per language")
	runCmd.Flags().IntVar(&flagTimeoutMs, "timeout-ms", 0, "Timeout per run in milliseconds (default: from problem.yaml)")
	runCmd.Flags().IntVar(&flagParallel, "parallel", 1, "Max number of concurrent language benchmarks")
	runCmd.Flags().StringVar(&flagExport, "export", "", "Export results to file (.json or .csv)")
	runCmd.Flags().BoolVar(&flagVerbose, "show-all-logs", false, "Show full Docker and command logs (disables progress bars)")
	runCmd.Flags().BoolVar(&flagNoSpinner, "no-spinner", false, "Disable spinner animation")

	rootCmd.AddCommand(runCmd)
}

func runBenchmark(cmd *cobra.Command, args []string) error {
	problemName := args[0]

	// Sync verbose state across UI models
	ui.Verbose = flagVerbose

	// Determine spinner suppression rules
	isCI := os.Getenv("CI") == "true"
	isTTY := term.IsTerminal(int(os.Stdout.Fd()))
	ui.NoSpinner = flagNoSpinner || isCI || !isTTY

	// Load problem spec first (gives useful errors even without Docker)
	spec, err := problems.LoadProblem(problemName)
	if err != nil {
		return err
	}

	// Determine languages
	langs := spec.Languages
	if flagLangs != "" {
		langs = strings.Split(flagLangs, ",")
		for i := range langs {
			langs[i] = strings.TrimSpace(langs[i])
		}
	}

	if len(langs) == 0 {
		return fmt.Errorf("no languages specified")
	}

	// Check Docker availability
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

	// Print minimalist header if not verbose
	if !flagVerbose {
		fmt.Fprintf(os.Stderr, "codebench v%s\n", Version)
		fmt.Fprintf(os.Stderr, "problem: %s\n", spec.Name)
		fmt.Fprintf(os.Stderr, "langs: %s\n", strings.Join(langs, ", "))
		fmt.Fprintf(os.Stderr, "input: %s | runs: %d (warmup: %d)", flagInput, flagRuns, flagWarmup)
		if flagParallel > 1 {
			fmt.Fprintf(os.Stderr, " | parallel: %d", flagParallel)
		}
		if flagExport != "" {
			fmt.Fprintf(os.Stderr, " | export: %s", flagExport)
		}
		fmt.Fprintf(os.Stderr, "\n\n")

		if ui.NoSpinner {
			fmt.Fprintf(os.Stderr, "Benchmarking, please wait...\n")
		}
	}

	// Run benchmarks
	config := benchmark.RunConfig{
		Problem:   spec,
		Languages: langs,
		InputSize: flagInput,
		Runs:      flagRuns,
		Warmup:    flagWarmup,
		TimeoutMs: timeoutMs,
		Parallel:  flagParallel,
		Verbose:   flagVerbose,
	}

	ui.StartSpinner(" Benchmarking, please wait...")

	benchStart := time.Now()
	metadata, results, err := benchmark.RunBenchmark(config)
	benchDuration := time.Since(benchStart)

	ui.StopSpinner()

	if err != nil {
		return err
	}

	// Sort: OK results by median ascending, then failures at the bottom
	sort.SliceStable(results, func(i, j int) bool {
		iOK := results[i].Status == benchmark.StatusOK
		jOK := results[j].Status == benchmark.StatusOK
		if iOK && jOK {
			return results[i].Stats.Median < results[j].Stats.Median
		}
		if iOK && !jOK {
			return true // OK comes before failures
		}
		return false
	})

	// Display results
	fmt.Fprintln(os.Stderr)
	renderSummaryLines(results)
	fmt.Fprintln(os.Stderr)
	if len(results) > 0 {
		renderPerformanceTable(results)
	}

	// Handle Export
	if flagExport != "" {
		ext := strings.ToLower(filepath.Ext(flagExport))
		report := export.Report{
			Metadata: metadata,
			Results:  results,
		}

		var exportErr error

		if ext == ".json" {
			exportErr = export.WriteJSON(report, flagExport)
		} else if ext == ".csv" {
			exportErr = export.WriteCSV(report, flagExport)
		} else {
			fmt.Fprintf(os.Stderr, "\n⚠ Unsupported export format %q. Please use .json or .csv\n", ext)
			return nil
		}

		if exportErr != nil {
			fmt.Fprintf(os.Stderr, "\n✗ Failed to export results to %s: %v\n", flagExport, exportErr)
		} else {
			fmt.Fprintf(os.Stderr, "\n✓ Successfully exported results to %s\n", flagExport)
		}
	}

	// Print total duration
	fmt.Fprintf(os.Stdout, "\ntotal duration: %s\n", formatTotalDuration(benchDuration))

	return nil
}

// formatTotalDuration returns milliseconds if < 1s, else nicely formatted seconds
func formatTotalDuration(d time.Duration) string {
	ms := float64(d.Microseconds()) / 1000.0
	if ms >= 1000 {
		return fmt.Sprintf("%.2fs", ms/1000.0)
	}
	return fmt.Sprintf("%.2fms", ms)
}

// ANSI Escape Codes
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorMagenta = "\033[35m"
)

func colorizeStatus(status string) string {
	switch status {
	case benchmark.StatusOK:
		return ColorGreen + status + ColorReset
	case benchmark.StatusFailed, benchmark.StatusNoSuccessfulRuns:
		return ColorRed + status + ColorReset
	case benchmark.StatusTimeout:
		return ColorYellow + status + ColorReset
	case benchmark.StatusBuildError:
		return ColorMagenta + status + ColorReset
	default:
		return status
	}
}

// renderSummaryLines prints out the metadata for each language before the main performance table.
func renderSummaryLines(results []benchmark.BenchmarkResult) {
	// Find the longest language name for alignment
	maxLangLen := 0
	for _, r := range results {
		if len(r.Language) > maxLangLen {
			maxLangLen = len(r.Language)
		}
	}

	for _, r := range results {
		var compileStr string
		if r.CompileTime > 0 {
			compileStr = fmtMs(r.CompileTime) + " ms"
		} else {
			compileStr = "N/A"
		}

		version := r.Version
		if version == "" {
			version = "unknown"
		}

		padLang := fmt.Sprintf("%-*s", maxLangLen, r.Language)
		colorStatus := colorizeStatus(r.Status)

		fmt.Fprintf(os.Stdout, "%s: %d/%d runs | %s | version: %s | compile: %s\n",
			padLang, r.Successes, r.TotalRuns, colorStatus, version, compileStr)
	}
}

// renderPerformanceTable renders benchmark performance metrics purely, omitting summary metadata.
func renderPerformanceTable(results []benchmark.BenchmarkResult) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.SetTitle("Benchmark Results (sorted by median execution time)")

	t.AppendHeader(table.Row{
		"Language", "Median(ms)", "Avg(ms)", "Min(ms)", "Max(ms)", "Mem Median(MB)", "Mem-Avg(MB)",
	})

	winnerSet := false

	for _, r := range results {
		var medianFmt, avgFmt, minFmt, maxFmt, memMedFmt, memAvgFmt string

		if r.Status == benchmark.StatusOK {
			if !winnerSet {
				medianFmt = ColorGreen + fmtMs(r.Stats.Median) + ColorReset
				winnerSet = true
			} else {
				medianFmt = fmtMs(r.Stats.Median)
			}
			avgFmt = fmtMs(r.Stats.Avg)
			minFmt = fmtMs(r.Stats.Min)
			maxFmt = fmtMs(r.Stats.Max)
			memMedFmt = fmtMem(r.Stats.MemMedian)
			memAvgFmt = fmtMem(r.Stats.MemAvg)
		} else {
			medianFmt = "N/A"
			avgFmt = "N/A"
			minFmt = "N/A"
			maxFmt = "N/A"
			memMedFmt = "N/A"
			memAvgFmt = "N/A"
		}

		t.AppendRow(table.Row{
			r.Language,
			medianFmt,
			avgFmt,
			minFmt,
			maxFmt,
			memMedFmt,
			memAvgFmt,
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
		{Number: 7, Align: text.AlignRight},
	})

	t.Render()
}

// fmtMs formats a duration as milliseconds with 2 decimal places.
func fmtMs(d time.Duration) string {
	return fmt.Sprintf("%.2f", float64(d.Microseconds())/1000.0)
}

func fmtMem(mb float64) string {
	if mb == 0 {
		return "N/A"
	}
	return fmt.Sprintf("%.2f", mb)
}
