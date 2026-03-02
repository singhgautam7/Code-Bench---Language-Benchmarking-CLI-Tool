package main
import ("bufio"; "os"; "strconv")
func mergesort(arr []int) []int {
    if len(arr) <= 1 { return arr }
    mid := len(arr)/2
    L := mergesort(arr[:mid])
    R := mergesort(arr[mid:])
    res := make([]int, 0, len(arr))
    i, j := 0, 0
    for i < len(L) && j < len(R) {
        if L[i] < R[j] { res = append(res, L[i]); i++ } else { res = append(res, R[j]); j++ }
    }
    res = append(res, L[i:]...)
    res = append(res, R[j:]...)
    return res
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var nums []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}
	mergesort(nums)
}