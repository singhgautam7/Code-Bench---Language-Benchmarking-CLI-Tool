# Prime Count (Sieve of Eratosthenes)

## 1️⃣ Overview
Calculates boolean prime limits up to an integer N using Sieve of Eratosthenes.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N log log N)
- **Space Complexity**: O(N)
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
An integer N.

**Example Input**:
```text
10
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: N = 100,000 (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: N = 20,000,000
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: N = 200,000,000
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Data structure stress tests, assessing massive sequential boolean array writes.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
