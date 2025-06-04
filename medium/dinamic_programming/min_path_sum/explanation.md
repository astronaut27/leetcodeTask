# Problem Statement
Given a `m x n` grid filled with non-negative numbers, find a path from the top left to the bottom right, which minimizes the sum of all numbers along its path. You can only move either down or right at any point in time.

## Example
- **Input**: grid = [[1,3,1],[1,5,1],[4,2,1]]
- **Output**: 7
- **Explanation**: The path that minimizes the sum is `1→3→1→1→1`, which totals 7.

# Approach
This problem is a classic example of dynamic programming for grid-based traversal:
1. **Initialization**:
    - Create a 2D array `dp` of the same dimensions as `grid` to store the minimum path sum to each cell.
    - Initialize `dp[0][0]` with the value of `grid[0][0]`, which is the starting point.

2. **Dynamic Programming Table**:
    - For each cell `(row, col)` in the grid, compute the minimum path sum by considering the paths from the top `(row-1, col)` and from the left `(row, col-1)`.
    - For cells in the first row, they can only be reached from the left, and for cells in the first column, they can only be reached from above.
    - Use a helper function `min()` to find the minimum sum for the current cell by adding the cell's value to the minimum of the values from the cells that can be reached from the top or left.

3. **Fill the DP Table**:
    - Loop through each row and each column of the grid. For each cell, determine the minimum path sum considering the possible directions (from top and left).
    - Avoid recomputation for the starting cell `(0,0)`.

4. **Result Extraction**:
    - The value at `dp[m-1][n-1]` (bottom-right corner of the `dp` table) contains the minimum path sum from the top-left to the bottom-right of the grid.

## Complexity Analysis
- **Time Complexity**: O(m * n), where `m` is the number of rows and `n` is the number of columns in the grid. Each cell is processed once.
- **Space Complexity**: O(m * n) due to the additional space used for the `dp` table. This can be optimized to O(n) if only the current and previous rows are stored at any time.

# Code
```go
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
