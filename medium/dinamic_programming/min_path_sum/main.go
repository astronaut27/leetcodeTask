package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 && col == 0 {
				continue
			}
			ans := 100000 // Large number to simulate infinity
			if row > 0 {
				ans = min(ans, dp[row-1][col])
			}
			if col > 0 {
				ans = min(ans, dp[row][col-1])
			}
			dp[row][col] = grid[row][col] + ans
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid)) // Output: 7
}
