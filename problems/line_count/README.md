# File Line Count

## 1️⃣ Overview
Splits a continuous string buffer over standard Unix linebreaks.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(N)
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
Strings separated by newlines.

**Example Input**:
```text
hello\nworld\nagain
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 10,000 lines (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 5,000,000 lines
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 5,000,000 lines
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
String memory boundaries, native IO tokenization algorithms without custom logic offsets.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
