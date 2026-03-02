package main
import ("bufio"; "os")
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    var arr []string
    for scanner.Scan() { arr = append(arr, scanner.Text()) }
    if len(arr) == 0 { return }
    target := arr[len(arr)-1]
    for i:=0; i<len(arr)-1; i++ {
        if arr[i] == target { break }
    }
}