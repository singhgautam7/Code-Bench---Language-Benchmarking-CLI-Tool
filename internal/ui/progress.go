package ui

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var Verbose bool
var globalSpinner *spinner.Spinner

// StartSpinner instantiates and runs a background spinner until stopped.
// It is a no-op if the CLI is running in verbose mode.
func StartSpinner(suffix string) {
	if Verbose {
		return
	}

	// Ensure we don't start multiple spinners
	if globalSpinner != nil && globalSpinner.Active() {
		return
	}

	globalSpinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	globalSpinner.Suffix = suffix
	globalSpinner.Writer = os.Stderr
	globalSpinner.Start()
}

// StopSpinner safely halts and clears the global spinner instance.
func StopSpinner() {
	if globalSpinner != nil && globalSpinner.Active() {
		globalSpinner.Stop()
	}
}
