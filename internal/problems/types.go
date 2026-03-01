package problems

// ProblemSpec represents the configuration for a benchmark problem,
// as defined in the problem's problem.yaml file.
type ProblemSpec struct {
	Name      string   `yaml:"name"`
	InputMode string   `yaml:"input_mode"`
	TimeoutMs int      `yaml:"timeout_ms"`
	Languages []string `yaml:"languages"`
}
