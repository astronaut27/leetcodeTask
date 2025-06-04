# Problem Statement
Given a `m x n` grid filled with either 0 or 1, where `0` indicates an open cell and `1` indicates an obstacle, determine the number of unique paths from the top-left corner to the bottom-right corner. You can only move either down or right at any point in time. If the start or end point is an obstacle, return 0.

## Example
- **Input**: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
- **Output**: 2
- **Explanation**: There are two ways to reach the bottom-right corner:
    1. Right -> Right -> Down -> Down
    2. Down -> Down -> Right -> Right

# Approach
The problem is tackled using a dynamic programming approach, similar to the "unique paths" problem but with an additional condition to handle obstacles:
1. **Initialize the DP Table**:
    - Create a 2D array `dp` with the same dimensions as `obstacleGrid` to store the number of ways to reach each cell.
    - Set `dp[0][0]` to 1 if the starting cell is not an obstacle; otherwise, set it to 0.
2. **Fill the DP Table**:
    - Iterate over each cell in the grid:
        - If a cell contains an obstacle (`obstacleGrid[row][col] == 1`), set `dp[row][col]` to 0.
        - Otherwise, accumulate paths from the top `(row-1, col)` and from the left `(row, col-1)`, considering boundary conditions.
3. **Boundary Conditions**:
    - For cells in the first row or first column, ensure that they do not sum paths from outside the grid boundaries and that they do not traverse through an obstacle.
4. **Result Extraction**:
    - The value at `dp[m-1][n-1]` will represent the total number of unique paths from the top-left to the bottom-right of the grid, considering obstacles.

## Complexity Analysis
- **Time Complexity**: O(m * n), where `m` is the number of rows and `n` is the number of columns. Each cell is visited exactly once.
- **Space Complexity**: O(m * n) for the `dp` table that holds the paths counts for each cell.

# Code
```go
package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 {
        return 0
    }

    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    
    dp[0][0] = 1
    for row := 0; row < m; row++ {
        for col := 0; col < n; col++ {
            if obstacleGrid[row][col] == 1 {
                dp[row][col] = 0
                continue
            }
            if row > 0 {
                dp[row][col] += dp[row-1][col]
            }
            if col > 0 {
                dp[row][col] += dp[row][col-1]
            }
        }
    }
    return dp[m-1][n-1]
}

func main() {
    obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
    fmt.Println(uniquePathsWithObstacles(obstacleGrid)) // Output: 2
}
