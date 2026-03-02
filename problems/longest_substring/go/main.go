package main
import ("io/ioutil"; "os"; "strings")
func main() {
    b, _ := ioutil.ReadAll(os.Stdin)
    s := strings.TrimSpace(string(b))
    chars := make(map[byte]bool)
    l, res := 0, 0
    for r:=0; r<len(s); r++ {
        for chars[s[r]] { delete(chars, s[l]); l++ }
        chars[s[r]] = true
        if r-l+1 > res { res = r-l+1 }
    }
}