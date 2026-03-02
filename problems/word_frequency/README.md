# Word Frequency Counter

## 1️⃣ Overview
Counts identical tokens across a large string source mapping occurrence metrics.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(U) where U is unique words
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
Space-separated dictionary words.

**Example Input**:
```text
apple banana apple date
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 10,000 words (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 1,000,000 words
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 10,000,000 words
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
String hashing logic, associative map insertions, and GC capabilities processing string objects.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
