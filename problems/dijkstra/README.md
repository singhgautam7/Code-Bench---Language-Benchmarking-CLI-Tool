# Dijkstra's Algorithm

## 1️⃣ Overview
Calculates the shortest path tree across a weighted directional graph.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O((V + E) log V)
- **Space Complexity**: O(V)
- **Pattern**: Iterative
- **Bottleneck**: CPU-bound
- **Allocation Intensity**: High

## 3️⃣ Input Format
V E (vertices/edges) followed by E triplets of (u, v, weight).

**Example Input**:
```text
3 2\n0 1 5\n1 2 10
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 1,000 Vertices, 5,000 Edges (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 50,000 Vertices, 200,000 Edges
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 200,000 Vertices, 2,000,000 Edges
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Priority queue data structure behavior, generic class abstractions, heap operations.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
