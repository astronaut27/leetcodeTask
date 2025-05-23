package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}
	fmt.Println("output:", minCostClimbingStairs(cost)) // Вывод: 15
}

// index 0, 1, 2 ,
// dp    0, 0,
// cost 10,15,20
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
