package benchmark

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// RunMetadata holds the environment and context information for a benchmark run.
type RunMetadata struct {
	Timestamp     string `json:"timestamp"`
	HostOS        string `json:"os"`
	HostArch      string `json:"arch"`
	HostCPUs      int    `json:"cpu"`
	DockerVersion string `json:"docker_version"`
	Problem       string `json:"problem"`
	Input         string `json:"input"`
	Runs          int    `json:"runs"`
	Warmup        int    `json:"warmup"`
	Parallel      int    `json:"parallel"`
}

// CaptureMetadata gathers environment info for reproducibility.
func CaptureMetadata(config RunConfig) RunMetadata {
	return RunMetadata{
		Timestamp:     time.Now().Format(time.RFC3339),
		HostOS:        runtime.GOOS,
		HostArch:      runtime.GOARCH,
		HostCPUs:      runtime.NumCPU(),
		DockerVersion: getDockerVersion(),
		Problem:       config.Problem.Name,
		Input:         config.InputSize,
		Runs:          config.Runs,
		Warmup:        config.Warmup,
		Parallel:      config.Parallel,
	}
}

// getDockerVersion queries the Docker daemon for its server version.
func getDockerVersion() string {
	cmd := exec.Command("docker", "version", "--format", "{{.Server.Version}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "unknown"
	}
	return strings.TrimSpace(out.String())
}
