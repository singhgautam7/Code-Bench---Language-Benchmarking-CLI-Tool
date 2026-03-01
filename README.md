# CodeBench

**A production-grade CLI benchmarking tool that compares algorithm implementations across programming languages using Docker containers.**

CodeBench measures compile time and execution time separately, supports warmup runs, enforces timeouts, and produces statistical summaries — all in isolated, reproducible Docker environments.

## Why CodeBench?

- **Fair comparison**: Every language runs in its own Docker container — same environment, every time
- **Accurate measurement**: Separates compile time from execution time; warmup runs eliminate cold-start noise
- **Statistical rigor**: Reports median, average, min, and max across multiple runs
- **Reproducible**: Docker guarantees identical environments across machines
- **Extensible**: Add new problems or languages by dropping files into the `problems/` directory

## Phase 1 Features

- Run benchmarks for a single problem across multiple languages
- Dockerized build and execution per language
- Configurable warmup and measured runs
- Separate compile time vs execution time measurement
- Median, Min, Max, Average statistics
- Timeout enforcement per run
- stdin-based input handling
- Deterministic input selection (small/medium/large)
- Sorted table output

---

## Prerequisites

- **Go** 1.22 or later
- **Docker** installed and running

## Installation

```bash
# Clone the repository
git clone https://github.com/gautam/codebench.git
cd codebench

# Build the binary
go build -o codebench .
```

## Usage

### Run Benchmarks

```bash
codebench run <problem> [flags]
```

### Flags

| Flag           | Default            | Description                                  |
|----------------|--------------------|----------------------------------------------|
| `--langs`      | all (from YAML)    | Comma-separated list of languages            |
| `--input`      | `small`            | Input size: `small`, `medium`, or `large`    |
| `--runs`       | `5`                | Number of measured runs per language          |
| `--warmup`     | `1`                | Number of warmup runs (excluded from stats)  |
| `--timeout-ms` | from `problem.yaml`| Timeout per run in milliseconds              |

### Examples

```bash
# Run all languages with defaults
codebench run binary_search

# Specific languages, medium input, 10 runs
codebench run binary_search --langs python,go,cpp --input medium --runs 10

# Custom timeout and warmup
codebench run binary_search --langs go,cpp --warmup 3 --timeout-ms 5000
```

### Sample Output

```
╔══════════════════════════════════════════╗
║          CodeBench — Phase 1             ║
╚══════════════════════════════════════════╝
  Problem:   binary_search
  Languages: python, go, cpp
  Input:     small
  Runs:      5 (warmup: 1)
  Timeout:   2000ms

╭──────────┬──────────────┬─────────────┬──────────┬──────────┬──────────┬───────╮
│ Language │ Compile (ms) │ Median (ms) │ Avg (ms) │ Min (ms) │ Max (ms) │ Runs  │
├──────────┼──────────────┼─────────────┼──────────┼──────────┼──────────┼───────┤
│ cpp      │      5230.00 │      120.50 │   125.30 │   118.20 │   140.10 │ 5/5   │
│ go       │      3450.00 │      135.80 │   140.20 │   130.50 │   155.00 │ 5/5   │
│ python   │       850.00 │      580.30 │   595.40 │   570.10 │   630.50 │ 5/5   │
╰──────────┴──────────────┴─────────────┴──────────┴──────────┴──────────┴───────╯
```

---

## Adding New Problems

### 1. Create the problem directory

```
problems/
  your_problem/
    problem.yaml
    inputs/
      small.txt
      medium.txt
      large.txt
    python/
      main.py
      Dockerfile
    go/
      main.go
      Dockerfile
    cpp/
      main.cpp
      Dockerfile
```

### 2. Define `problem.yaml`

```yaml
name: your_problem
input_mode: stdin
timeout_ms: 2000
languages:
  - python
  - go
  - cpp
```

### 3. Create input files

Place input data in `inputs/small.txt`, `inputs/medium.txt`, and `inputs/large.txt`. These are piped to programs via stdin.

### 4. Implement solutions

Each language needs a source file and a `Dockerfile`. The program must:
- Read input from **stdin**
- Print the result to **stdout**

### 5. Run it

```bash
codebench run your_problem
```

---

## Project Structure

```
CodeBench/
  main.go                       # Entry point
  cmd/
    root.go                     # Cobra root command
    run.go                      # Run subcommand + table rendering
  internal/
    benchmark/
      docker.go                 # Docker build & run via os/exec
      runner.go                 # Benchmark orchestration
      stats.go                  # Statistical computations
    problems/
      loader.go                 # Problem YAML loader
      types.go                  # ProblemSpec type definition
  problems/
    binary_search/              # Example problem
      problem.yaml
      inputs/
      python/
      go/
      cpp/
```

---

## Maintenance Rule

> **The README must always be updated whenever new commands, flags, features, or architectural changes are introduced.
> It serves as the single source of truth for usage and capabilities.**

---

## License

MIT
