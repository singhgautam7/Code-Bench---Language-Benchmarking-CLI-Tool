# Matrix Multiplication

## 1️⃣ Overview
Multiplies two N x N matrices sequentially.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N^3)
- **Space Complexity**: O(N^2)
- **Pattern**: Iterative
- **Bottleneck**: CPU-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
Integer N representing size, followed by NxN elements for Matrix A, then NxN elements for Matrix B.

**Example Input**:
```text
2\n1 2\n3 4\n1 2\n3 4
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 100x100 matrices (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 300x300 matrices
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 600x600 matrices
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Tests deep multidimensional array access loops, CPU cache locality, and nested branch efficiency.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
