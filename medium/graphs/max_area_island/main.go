package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    maxArea := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                area := 0
                dfs(grid, i, j, &area)
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }
    return maxArea
}

func dfs(grid [][]int, i, j int, area *int) {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
        return
    }
    grid[i][j] = 0
    *area++
    dfs(grid, i+1, j, area)
    dfs(grid, i-1, j, area)
    dfs(grid, i, j+1, area)
    dfs(grid, i, j-1, area)
}

func main() {
    grid := [][]int{
        {1, 1, 0, 0, 0},
        {1, 1, 0, 0, 0},
        {0, 0, 0, 1, 1},
        {0, 0, 0, 1, 1},
    }
    fmt.Println("Maximum area of island is:", maxAreaOfIsland(grid))
}