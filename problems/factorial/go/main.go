package main
import ("io/ioutil"; "strconv"; "strings"; "math/big"; "os")
func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	n, _ := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	s := big.NewInt(1)
	for i:=int64(1); i<=n; i++ { s.Mul(s, big.NewInt(i)) }
}