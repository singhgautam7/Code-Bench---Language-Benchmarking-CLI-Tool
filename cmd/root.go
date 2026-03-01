package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codebench",
	Short: "CodeBench — benchmark algorithm implementations across languages",
	Long: `CodeBench is a CLI tool that benchmarks algorithm implementations
across multiple programming languages using Docker containers for
isolation and reproducibility.

It measures compile time and execution time separately, supports
warmup runs, and produces statistical summaries (median, avg, min, max).`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
