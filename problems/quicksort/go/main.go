package main

import (
	"bufio"
	"os"
	"strconv"
)

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := len(arr) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quicksort(arr[:left])
	quicksort(arr[left+1:])
	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var nums []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}
	quicksort(nums)
}
