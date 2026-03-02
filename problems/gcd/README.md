# GCD Computation

## 1️⃣ Overview
Determines the greatest common divisor using the Euclidean algorithm iterated.

## 2️⃣ Computational Characteristics
- **Time Complexity**: O(N log(min(A,B)))
- **Space Complexity**: O(1)
- **Pattern**: Iterative
- **Bottleneck**: CPU-bound
- **Allocation Intensity**: Low

## 3️⃣ Input Format
An integer N for repetitions.

**Example Input**:
```text
1000
```

## 4️⃣ Input Sets (Tested Parameters)

### `small`
- **Size Description**: 1,000,000 loops (< 50ms)
- **Intended Purpose**: Quick checks ensuring syntax parses completely without bottlenecks.

### `medium`
- **Size Description**: 15,000,000 loops
- **Intended Purpose**: Moderate workload triggering standard Garbage Collection pauses.

### `large`
- **Size Description**: 300,000,000 loops
- **Intended Purpose**: Heavy capacity limits exploring OOM boundaries and memory allocator faults. Inputs generated dynamically.

## 5️⃣ Benchmark Intent
Repeated Modulus calculation mathematics and tight inner loops.

## 6️⃣ Notes on Fairness
- **Idiomatic Models**: All languages employ language-standard solutions conforming to the test (e.g., using specific structures for graphs, recursion, or loops).
- **No Micro-Optimization**: Solutions avoid extreme unsafe bindings or raw binary hacks.
- **Deterministic**: Every run across every language tests perfectly equivalent dataset distributions shipped in standard input blocks.
- **Fair Sizing**: Algorithm limits restrict behaviors to strictly mapped $O()$ equivalents matching original designs.
