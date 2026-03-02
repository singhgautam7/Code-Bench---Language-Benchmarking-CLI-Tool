package main
import ("io/ioutil"; "strconv"; "strings"; "os")
func gcd(a, b int) int {
    for b != 0 { t := b; b = a % b; a = t }
    return a
}
func main() {
    b, _ := ioutil.ReadAll(os.Stdin)
    n, _ := strconv.Atoi(strings.TrimSpace(string(b)))
    for i:=1; i<=n; i++ { gcd(1000000007, i) }
}