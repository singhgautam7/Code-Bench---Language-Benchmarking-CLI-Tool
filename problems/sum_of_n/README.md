# Sum of First N Numbers

## 1️⃣ Overview
Calculates the sum of integers from 1 to N sequentially.

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
1000
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 1,000 (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 1,000,000
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 500,000,000
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Numeric arithmetic throughput and basic loop overhead.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
