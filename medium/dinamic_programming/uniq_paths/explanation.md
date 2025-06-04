# Problem Statement
Given a grid with dimensions `m x n`, calculate the number of unique paths from the top-left corner to the bottom-right corner of the grid. You can only move either down or right at any point in time.

## Example
- **Input**: m = 3, n = 2
- **Output**: 3
- **Explanation**: From the top-left corner, there are a total of three paths to the bottom-right corner:
    1. Right -> Right -> Down
    2. Right -> Down -> Right
    3. Down -> Right -> Right

# Approach
This problem is a classic dynamic programming problem that can be solved by building a table `dp` where each entry `dp[i][j]` represents the number of unique paths to that position from the top-left corner.
1. **Initialize the DP Table**:
    - Create a 2D array `dp` with dimensions `m x n`.
    - Set `dp[0][0]` to `1` since there's exactly one way to be at the start.
2. **Fill the DP Table**:
    - Iterate through each cell in the table. For each cell at position `(row, col)`, determine the number of ways to reach it by summing the ways to reach the cell directly above it (`dp[row-1][col]` if `row > 0`) and the cell to the left of it (`dp[row][col-1]` if `col > 0`).
3. **Iterate Through the Grid**:
    - For each position, add the number of paths from the top (if applicable) and from the left (if applicable).
    - This approach ensures that each cell `dp[i][j]` contains the number of unique paths to reach it from the start.
4. **Return the Result**:
    - The bottom-right corner of the grid, `dp[m-1][n-1]`, will contain the total number of unique paths from the start position to the end.

## Complexity Analysis
- **Time Complexity**: O(m * n), where `m` is the number of rows and `n` is the number of columns. Each cell in the `m x n` grid is filled exactly once.
- **Space Complexity**: O(m * n) due to the space used by the `dp` table. This can be reduced to O(min(m, n)) by only keeping the current and previous rows or columns at any time.

# Code
```go
package main

import "fmt"

func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = 1 // Start position
    for row := 0; row < m; row++ {
        for col := 0; col < n; col++ {
            if row > 0 {
                dp[row][col] += dp[row-1][col]
            }
            if col > 0 {
                dp[row][col] += dp[row][col-1]
            }
        }
    }
    return dp[m-1][n-1] // End position
}

func main() {
    m := 3
    n := 2
    fmt.Println(uniquePaths(m, n)) // Output: 3
}
