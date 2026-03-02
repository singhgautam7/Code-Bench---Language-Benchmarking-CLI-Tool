package main
import ("bufio"; "os"; "strconv"; "container/heap")
type Edge struct { to, weight int }
type Item struct { node, dist int; index int }
type PriorityQueue []*Item
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x any) { n := len(*pq); item := x.(*Item); item.index = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() any { old := *pq; n := len(old); item := old[n-1]; old[n-1] = nil; item.index = -1; *pq = old[0 : n-1]; return item }

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    if !scanner.Scan() { return }
    V, _ := strconv.Atoi(scanner.Text())
    scanner.Scan(); E, _ := strconv.Atoi(scanner.Text())
    adj := make([][]Edge, V)
    for i:=0; i<E; i++ {
        scanner.Scan(); u, _ := strconv.Atoi(scanner.Text())
        scanner.Scan(); v, _ := strconv.Atoi(scanner.Text())
        scanner.Scan(); w, _ := strconv.Atoi(scanner.Text())
        adj[u] = append(adj[u], Edge{v, w})
    }
    dist := make([]int, V)
    for i:=range dist { dist[i] = 1e9 }
    dist[0] = 0
    pq := make(PriorityQueue, 0)
    heap.Push(&pq, &Item{node: 0, dist: 0})
    for pq.Len() > 0 {
        curr := heap.Pop(&pq).(*Item)
        if curr.dist > dist[curr.node] { continue }
        for _, edge := range adj[curr.node] {
            if dist[curr.node] + edge.weight < dist[edge.to] {
                dist[edge.to] = dist[curr.node] + edge.weight
                heap.Push(&pq, &Item{node: edge.to, dist: dist[edge.to]})
            }
        }
    }
}