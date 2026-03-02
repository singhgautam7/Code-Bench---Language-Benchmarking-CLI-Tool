# Linear Search

## 1️⃣ Overview
Scans across an entire array to locate a target, representing worst-case O(N) traversal.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(1)
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound
- **Allocation Intensity**: Low

## 3️⃣ Input Format
Space-separated items followed by the target to search.

**Example Input**:
```text
1 2 3 4 5\n5
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 10,000 elements (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 2,000,000 elements
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 10,000,000 elements
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Baseline comparison for contiguous array memory scanning.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
