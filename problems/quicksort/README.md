# QuickSort

## 1️⃣ Overview
Sorts an array by selecting a pivot and recursively splitting the array into smaller items and larger items.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N log N)
- **Space Complexity**: O(log N) stack
- **Pattern**: Recursive
- **Bottleneck**: CPU-bound
- **Allocation Intensity**: Medium

## 3️⃣ Input Format
Space-separated integers.

**Example Input**:
```text
3 1 4 1 5 9 2
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 10,000 elements (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 300,000 elements
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 2,000,000 elements
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Recursion overhead, pointer/index swaps, and divide-and-conquer paradigms.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
