# Factorial (Iterative)

## 1️⃣ Overview
Calculates the factorial of N iteratively.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(1)
- **Pattern**: Iterative
- **Bottleneck**: CPU-bound
- **Allocation Intensity**: Low

## 3️⃣ Input Format
A single integer N.

**Example Input**:
```text
5
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 100 (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 10,000
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 250,000
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Integer arithmetic throughput. Note: due to lack of arbitrary precision in some systems, tests wrapping behavior.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
