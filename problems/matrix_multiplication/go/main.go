package main
import ("bufio"; "os"; "strconv")
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    if !scanner.Scan() { return }
    n, _ := strconv.Atoi(scanner.Text())
    A := make([][]int, n)
    B := make([][]int, n)
    for i:=0; i<n; i++ {
        A[i] = make([]int, n)
        for j:=0; j<n; j++ { scanner.Scan(); A[i][j], _ = strconv.Atoi(scanner.Text()) }
    }
    for i:=0; i<n; i++ {
        B[i] = make([]int, n)
        for j:=0; j<n; j++ { scanner.Scan(); B[i][j], _ = strconv.Atoi(scanner.Text()) }
    }
    C := make([][]int, n)
    for i:=0; i<n; i++ { C[i] = make([]int, n) }
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            for k:=0; k<n; k++ { C[i][j] += A[i][k] * B[k][j] }
        }
    }
}