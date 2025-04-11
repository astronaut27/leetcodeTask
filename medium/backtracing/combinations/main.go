package main

import "fmt"

func combine(n int, k int) [][]int {
	var ans [][]int
	backtrack([]int{}, 1, n, k, &ans)
	return ans
}

func backtrack(curr []int, index int, n int, k int, ans *[][]int) {
	if len(curr) == k { // Base case: length k is reached
		*ans = append(*ans, append([]int{}, curr...))
		return
	}
	for i := index; i <= n; i++ {
		curr = append(curr, i)
		backtrack(curr, i+1, n, k, ans)
		curr = curr[:len(curr)-1] // Backtrack
	}
}

func main() {
	n := 4
	k := 2
	result := combine(n, k)
	fmt.Println("All combinations:", result)
}
