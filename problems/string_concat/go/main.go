package main
import ("fmt"; "io/ioutil"; "os"; "strings"; "strconv")
func main() {
    b, _ := ioutil.ReadAll(os.Stdin)
    p := strings.Fields(string(b))
    if len(p) == 2 {
        n, _ := strconv.Atoi(p[1])
        var sb strings.Builder
        for i:=0; i<n; i++ { sb.WriteString(p[0]) }
    }
}