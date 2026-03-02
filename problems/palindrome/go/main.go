package main
import ("io/ioutil"; "strings"; "os")
func main() {
    b, _ := ioutil.ReadAll(os.Stdin)
    s := strings.TrimSpace(string(b))
    for i:=0; i<len(s)/2; i++ {
      if s[i] != s[len(s)-1-i] { return }
    }
}