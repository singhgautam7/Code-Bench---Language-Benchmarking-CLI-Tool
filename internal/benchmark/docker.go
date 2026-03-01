package benchmark

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// BuildImage builds a Docker image for the given problem/language combination.
// Returns the build duration and any error encountered.
func BuildImage(problem, language, contextDir string) (time.Duration, error) {
	imageName := fmt.Sprintf("codebench/%s/%s", problem, language)

	cmd := exec.Command("docker", "build", "-t", imageName, contextDir)
	cmd.Stdout = os.Stderr // Show build output on stderr for progress
	cmd.Stderr = os.Stderr

	start := time.Now()
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("docker build failed for %s/%s: %w", problem, language, err)
	}
	buildTime := time.Since(start)

	return buildTime, nil
}

// RunContainer runs a Docker container with the given image, piping the input file
// via stdin. Enforces the specified timeout. Returns execution time, exit code, and error.
func RunContainer(imageName, inputPath string, timeoutMs int) (time.Duration, int, error) {
	timeout := time.Duration(timeoutMs) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "-i", imageName)

	// Pipe input file to stdin
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return 0, -1, fmt.Errorf("failed to open input file %q: %w", inputPath, err)
	}
	defer inputFile.Close()
	cmd.Stdin = inputFile

	// Discard stdout/stderr from the container during benchmarking
	cmd.Stdout = nil
	cmd.Stderr = nil

	start := time.Now()
	runErr := cmd.Run()
	execTime := time.Since(start)

	// Check if context deadline was exceeded (timeout)
	if ctx.Err() == context.DeadlineExceeded {
		return execTime, -1, fmt.Errorf("TIMEOUT: execution exceeded %dms", timeoutMs)
	}

	if runErr != nil {
		// Try to extract exit code
		if exitErr, ok := runErr.(*exec.ExitError); ok {
			return execTime, exitErr.ExitCode(), fmt.Errorf("container exited with code %d", exitErr.ExitCode())
		}
		return execTime, -1, fmt.Errorf("container execution failed: %w", runErr)
	}

	return execTime, 0, nil
}

// CheckDockerAvailable verifies that Docker is installed and accessible.
func CheckDockerAvailable() error {
	cmd := exec.Command("docker", "info")
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker is not available: %w\nPlease ensure Docker is installed and running", err)
	}
	return nil
}
