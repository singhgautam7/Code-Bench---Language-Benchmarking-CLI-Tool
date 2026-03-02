# JSON Parsing

## 1️⃣ Overview
Loads and deserializes dynamic JSON payload structures into language-native models.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(N)
- **Pattern**: Iterative
- **Bottleneck**: CPU-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
A JSON payload array of objects.

**Example Input**:
```text
[{"id":1, "val":2}]
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 5,000 objects (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 150,000 objects
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 2,000,000 objects
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Parsing limits, Reflection/Standard Library efficiency, memory footprint sizes of dynamic structures.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
