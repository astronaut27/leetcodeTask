package main

func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])+1

	// Initialize dp table with boundary values set to a large number
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0], dp[i][n] = 10000, 10000 // Set boundaries to a large number (since we're looking for the minimum)
	}

	// Initialize the first row of dp
	for j := 1; j < n; j++ {
		dp[0][j] = matrix[0][j-1]
	}

	// Fill the dp table row by row
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = min(min(dp[row-1][col], dp[row-1][col-1]), dp[row-1][col+1]) + matrix[row][col-1]
		}
	}

	// Find the minimum value in the last row of dp
	result := 10000
	for j := 1; j < n; j++ {
		result = min(dp[m-1][j], result)
	}

	return result
}

// Helper function to find the minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
