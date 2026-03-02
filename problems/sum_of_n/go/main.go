package main
import ("fmt"; "io/ioutil"; "strconv"; "strings"; "os")
func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	n, _ := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	s := int64(0)
	for i:=int64(1); i<=n; i++ { s += i }
}