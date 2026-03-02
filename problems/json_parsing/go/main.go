package main
import ("encoding/json"; "io/ioutil"; "os")
func main() {
    b, _ := ioutil.ReadAll(os.Stdin)
    var v interface{}
    json.Unmarshal(b, &v)
}