# Binary Search

## 1️⃣ Overview
Finds the position of a target value within a sorted array using divide-and-conquer strategy.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(log N)
- **Space Complexity**: O(1)
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound (cache misses)
- **Allocation Intensity**: Low

## 3️⃣ Input Format
Space-separated sorted integers followed by a newline, then the target integer.

**Example Input**:
```text
1 3 5 7 9\n5
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 1,000 elements (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 100,000 elements
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 5,000,000 elements
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Measures array indexing efficiency and CPU branch prediction.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
