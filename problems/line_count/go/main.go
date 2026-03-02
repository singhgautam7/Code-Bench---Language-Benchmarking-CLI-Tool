package main
import ("bufio"; "os")
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    c := 0
    for scanner.Scan() { c++ }
}