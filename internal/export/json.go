package export

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/singhgautam7/Code-Bench---Language-Benchmarking-CLI-Tool/internal/benchmark"
)

// Report structure encapsulating the full benchmark payload
type Report struct {
	Metadata benchmark.RunMetadata       `json:"metadata"`
	Results  []benchmark.BenchmarkResult `json:"results"`
}

// WriteJSON serializes the benchmark report to a JSON file.
func WriteJSON(report Report, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create export file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}
