package main
import ("bufio"; "os"; "strconv")
var adj [][]int
var vis []bool
func dfs(u int) {
    vis[u] = true
    for _, v := range adj[u] { if !vis[v] { dfs(v) } }
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    if !scanner.Scan() { return }
    V, _ := strconv.Atoi(scanner.Text())
    scanner.Scan(); E, _ := strconv.Atoi(scanner.Text())
    adj = make([][]int, V)
    for i:=0; i<E; i++ {
        scanner.Scan(); u, _ := strconv.Atoi(scanner.Text())
        scanner.Scan(); v, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        adj[u] = append(adj[u], v)
    }
    vis = make([]bool, V)
    dfs(0)
}