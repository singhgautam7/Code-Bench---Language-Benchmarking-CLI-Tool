package main
import ("io/ioutil"; "os"; "strconv"; "strings")
func main() {
    b, _ := ioutil.ReadAll(os.Stdin)
    n, _ := strconv.Atoi(strings.TrimSpace(string(b)))
    if n < 2 { return }
    sieve := make([]bool, n+1)
    count := 0
    for p:=2; p<=n; p++ {
        if !sieve[p] {
            count++
            for i:=p*p; i<=n; i+=p { sieve[i] = true }
        }
    }
}