# Palindrome Check

## 1️⃣ Overview
Checks if a given string reads the same forwards and backwards.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(1)
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound
- **Allocation Intensity**: Low

## 3️⃣ Input Format
A single continuous string.

**Example Input**:
```text
racecar
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 10K characters (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 500K characters
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 30M characters
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
String access speeds and memory bandwidth for continuous scanning.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
