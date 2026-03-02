package main
import ("io/ioutil"; "strconv"; "strings"; "os")
func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	n, _ := strconv.Atoi(strings.TrimSpace(string(b)))
	a, b_val := 0, 1
	for i:=0; i<n; i++ {
        t := a
        a = b_val
        b_val = t + b_val
    }
}