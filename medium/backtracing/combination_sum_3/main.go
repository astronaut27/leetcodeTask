package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	backtrack([]int{}, 1, n, k, &ans)
	return ans
}

func backtrack(curr []int, index int, n int, k int, ans *[][]int) {
	if len(curr) == k && sum(curr) == n { // Base case: length k is reached
		*ans = append(*ans, append([]int{}, curr...))
		return
	}
	for i := index; i < 10; i++ {
		curr = append(curr, i)
		backtrack(curr, i+1, n, k, ans)
		curr = curr[:len(curr)-1] // Backtrack
	}
}

func sum(curr []int) int {
	var ans int
	for i := range curr {
		ans += curr[i]
	}
	return ans
}
