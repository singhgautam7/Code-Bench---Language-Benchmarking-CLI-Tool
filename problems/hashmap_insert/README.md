# HashMap Insert

## 1️⃣ Overview
Generates a map sequentially inserting numeric string keys and numeric IDs.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N)
- **Space Complexity**: O(N)
- **Pattern**: Iterative
- **Bottleneck**: Memory-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
An integer N dictating the insert loop.

**Example Input**:
```text
1000
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 100,000 loops (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 2,000,000 loops
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 40,000,000 loops
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Core map memory sizing, underlying hash collision avoidance, object allocations.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
