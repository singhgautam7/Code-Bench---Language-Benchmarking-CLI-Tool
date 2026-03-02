# Longest Substring Without Repeating Characters

## 1️⃣ Overview
Finds the length of the longest substring without repeating characters.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(min(N, M)) where M is charset
- **Pattern**: Iterative
- **Bottleneck**: CPU-bound
- **Allocation Intensity**: Medium

## 3️⃣ Input Format
A single random string.

**Example Input**:
```text
abcabcbb
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 10,000 characters (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 500,000 characters
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 20,000,000 characters
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Hash set insertions/deletions and sliding window index tracking.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
