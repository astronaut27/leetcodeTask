package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0]) // Get dimensions of the matrix
	i, j := 0, n-1                      // Start at the top-right corner of the matrix
	for i < m && j >= 0 {               // Ensure indices are within the bounds
		if matrix[i][j] == target {
			return true // Target found
		}
		if matrix[i][j] > target {
			j-- // Move left
		} else {
			i++ // Move down
		}
	}
	return false // Target not found
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	fmt.Println("Target found:", searchMatrix(matrix, target))
}
