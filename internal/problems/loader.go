package problems

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadProblem reads and parses the problem.yaml for the given problem name.
// It looks for problems/<problemName>/problem.yaml relative to the working directory.
func LoadProblem(problemName string) (*ProblemSpec, error) {
	problemDir := filepath.Join("problems", problemName)

	// Check problem directory exists
	info, err := os.Stat(problemDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("problem %q not found: directory %q does not exist", problemName, problemDir)
		}
		return nil, fmt.Errorf("error accessing problem directory: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("problem path %q is not a directory", problemDir)
	}

	// Read problem.yaml
	yamlPath := filepath.Join(problemDir, "problem.yaml")
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("problem.yaml not found in %q", problemDir)
		}
		return nil, fmt.Errorf("error reading problem.yaml: %w", err)
	}

	var spec ProblemSpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("error parsing problem.yaml: %w", err)
	}

	if spec.Name == "" {
		spec.Name = problemName
	}

	return &spec, nil
}

// GetInputPath returns the absolute path to the input file for the given
// problem and input size (small, medium, large).
func GetInputPath(problemName, inputSize string) (string, error) {
	validSizes := map[string]bool{"small": true, "medium": true, "large": true}
	if !validSizes[inputSize] {
		return "", fmt.Errorf("invalid input size %q: must be one of small, medium, large", inputSize)
	}

	inputFile := filepath.Join("problems", problemName, "inputs", inputSize+".txt")

	absPath, err := filepath.Abs(inputFile)
	if err != nil {
		return "", fmt.Errorf("error resolving input path: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("input file %q not found", inputFile)
		}
		return "", fmt.Errorf("error accessing input file: %w", err)
	}

	return absPath, nil
}

// GetLanguageDir returns the path to the language implementation directory.
func GetLanguageDir(problemName, language string) (string, error) {
	langDir := filepath.Join("problems", problemName, language)

	info, err := os.Stat(langDir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("language %q implementation not found for problem %q", language, problemName)
		}
		return "", fmt.Errorf("error accessing language directory: %w", err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("language path %q is not a directory", langDir)
	}

	absPath, err := filepath.Abs(langDir)
	if err != nil {
		return "", fmt.Errorf("error resolving language directory path: %w", err)
	}

	return absPath, nil
}
