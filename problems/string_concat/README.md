# String Concatenation

## 1️⃣ Overview
Appends a single string token to a buffer thousands of times.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(N)
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
Character to append and N repetitions.

**Example Input**:
```text
A 100
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: A x 100,000 (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: A x 1,500,000
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: B x 30,000,000
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Continuous memory reallocation, StringBuilder speeds vs raw immutable string limits.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
