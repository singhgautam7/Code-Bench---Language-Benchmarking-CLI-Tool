package main
import ("io/ioutil"; "strconv"; "strings"; "os")
func main() {
    b, _ := ioutil.ReadAll(os.Stdin)
    n, _ := strconv.Atoi(strings.TrimSpace(string(b)))
    h := make(map[string]int)
    for i:=0; i<n; i++ { h[strconv.Itoa(i)] = i }
}