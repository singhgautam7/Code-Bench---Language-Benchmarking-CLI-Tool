package benchmark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)



// BuildImage builds a Docker image for the given problem/language combination.
// Returns the build duration and any error encountered.
func BuildImage(problem, language, contextDir string, verbose bool) (time.Duration, error) {
	imageName := fmt.Sprintf("codebench/%s/%s", problem, language)

	cmd := exec.Command("docker", "build", "-t", imageName, contextDir)

	var stderr bytes.Buffer
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = &bytes.Buffer{}
		cmd.Stderr = &stderr
	}

	start := time.Now()
	if err := cmd.Run(); err != nil {
		if !verbose {
			// On failure, print the hidden buffer so the user knows what broke
			fmt.Fprintf(os.Stderr, "\n%s", stderr.String())
		}
		return time.Since(start), fmt.Errorf("docker build failed for %s/%s: %w", problem, language, err)
	}
	buildTime := time.Since(start)

	return buildTime, nil
}

// getImageCmd inspects the docker image to find its default Cmd
func getImageCmd(imageName string) ([]string, error) {
	cmd := exec.Command("docker", "inspect", "--format='{{json .Config.Cmd}}'", imageName)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	outStr := strings.TrimSpace(string(out))
	if outStr == "" || outStr == "null" {
		return nil, fmt.Errorf("no command found for image %s", imageName)
	}

	// Unquote the string if docker inspect wrapped it in quotes
	outStr = strings.Trim(outStr, "'\"")

	var cmdSlice []string
	if err := json.Unmarshal([]byte(outStr), &cmdSlice); err != nil {
		return nil, fmt.Errorf("failed to parse image Cmd: %w", err)
	}

	return cmdSlice, nil
}

// memRegex parses the resident set size output from GNU time -v
var memRegex = regexp.MustCompile(`Maximum resident set size \(kbytes\):\s+(\d+)`)

// RunContainer runs a Docker container with the given image, piping the input file
// via stdin. It intercepts the image execution to use /usr/bin/time -v for memory tracking.
// Returns execution time, max memory usage in MB, exit code, and error.
func RunContainer(imageName, inputPath string, timeoutMs int, verbose bool) (time.Duration, float64, int, error) {
	timeout := time.Duration(timeoutMs) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Get original cmd from image to wrap it
	origCmd, err := getImageCmd(imageName)
	if err != nil {
		return 0, 0, -1, err
	}

	// Prepare time override args
	runArgs := []string{"run", "--rm", "-i", "--entrypoint", "/usr/bin/time"}
	runArgs = append(runArgs, imageName, "-v")
	runArgs = append(runArgs, origCmd...)

	cmd := exec.CommandContext(ctx, "docker", runArgs...)

	// Pipe input file to stdin
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return 0, 0, -1, fmt.Errorf("failed to open input file %q: %w", inputPath, err)
	}
	defer inputFile.Close()
	cmd.Stdin = inputFile

	// Routing setup
	var stdout bytes.Buffer
	var memBuffer bytes.Buffer // time -v writes to stderr exactly

	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = io.MultiWriter(os.Stderr, &memBuffer)
	} else {
		cmd.Stdout = &stdout
		cmd.Stderr = &memBuffer
	}

	start := time.Now()
	runErr := cmd.Run()
	execTime := time.Since(start)

	// Check if context deadline was exceeded (timeout)
	if ctx.Err() == context.DeadlineExceeded {
		return execTime, 0, -1, fmt.Errorf("TIMEOUT: execution exceeded %dms", timeoutMs)
	}

	// Parse memory regardless of minor errors, as long as it didn't timeout
	var memMB float64
	memMatch := memRegex.FindStringSubmatch(memBuffer.String())
	if len(memMatch) == 2 {
		if kbytes, parseErr := strconv.ParseFloat(memMatch[1], 64); parseErr == nil {
			memMB = kbytes / 1024.0
		}
	}

	if runErr != nil {
		if !verbose && !isTimeout(runErr) {
			// Dump hidden crash reason to error stack if not specifically a timeout.
			fmt.Fprintf(os.Stderr, "\n%s\n%s", stdout.String(), memBuffer.String())
		}
		// Try to extract exit code
		if exitErr, ok := runErr.(*exec.ExitError); ok {
			return execTime, memMB, exitErr.ExitCode(), fmt.Errorf("container exited with code %d", exitErr.ExitCode())
		}
		return execTime, memMB, -1, fmt.Errorf("container execution failed: %w", runErr)
	}

	return execTime, memMB, 0, nil
}

// GetLanguageVersion reads the pre-baked version string from /version.txt inside the container image.
// Returns "unavailable" if detection fails.
func GetLanguageVersion(imageName, language string) string {
	cmd := exec.Command("docker", "run", "--rm", "--entrypoint", "cat", imageName, "/version.txt")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "unavailable"
	}

	output := strings.TrimSpace(stdout.String())
	if output == "" {
		return "unavailable"
	}

	// Take first line only
	if idx := strings.IndexByte(output, '\n'); idx != -1 {
		output = output[:idx]
	}

	return output
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
