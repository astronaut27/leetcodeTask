package main

import "fmt"

func findMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	count, maxlen := 0, 0
	store := make(map[int]int)
	store[0] = -1
	for i, v := range nums {
		if v == 1 {
			count++
		} else {
			count--
		}
		if v, ok := store[count]; ok {
			maxlen = max(maxlen, i-v)
		} else {
			store[count] = i
		}
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 1, 0}
	fmt.Println(findMaxLength(nums)) // Output: 2
}
