package main
import ("bufio"; "os"; "strconv")
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    if !scanner.Scan() { return }
    V, _ := strconv.Atoi(scanner.Text())
    scanner.Scan(); E, _ := strconv.Atoi(scanner.Text())
    adj := make([][]int, V)
    for i:=0; i<E; i++ {
        scanner.Scan(); u, _ := strconv.Atoi(scanner.Text())
        scanner.Scan(); v, _ := strconv.Atoi(scanner.Text())
        scanner.Scan() // w
        adj[u] = append(adj[u], v)
    }
    q := []int{0}
    vis := make([]bool, V)
    vis[0] = true
    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        for _, nxt := range adj[curr] {
            if !vis[nxt] { vis[nxt]=true; q=append(q, nxt) }
        }
    }
}