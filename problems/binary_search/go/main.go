package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// binarySearch performs binary search on a sorted slice.
// Returns the index of target, or -1 if not found.
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read all input
	var lines []string
	scanner := bufio.NewScanner(reader)
	// Increase buffer for large inputs
	scanner.Buffer(make([]byte, 0), 64*1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 1 {
		fmt.Fprintln(os.Stderr, "Error: expected array and target on stdin")
		os.Exit(1)
	}

	// Parse: all tokens together, last is target
	var allTokens []string
	for _, line := range lines {
		tokens := strings.Fields(line)
		allTokens = append(allTokens, tokens...)
	}

	if len(allTokens) < 2 {
		fmt.Fprintln(os.Stderr, "Error: need at least array + target")
		os.Exit(1)
	}

	// Last token is target, rest are array
	target, err := strconv.Atoi(allTokens[len(allTokens)-1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing target: %v\n", err)
		os.Exit(1)
	}

	arr := make([]int, len(allTokens)-1)
	for i := 0; i < len(allTokens)-1; i++ {
		arr[i], err = strconv.Atoi(allTokens[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing array element %d: %v\n", i, err)
			os.Exit(1)
		}
	}

	result := binarySearch(arr, target)
	fmt.Println(result)
}
