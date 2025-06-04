package main

import (
	"fmt"
)

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func numIslandsIterative(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				stack := [][]int{{i, j}}
				grid[i][j] = '0'

				for len(stack) > 0 {
					node := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					x, y := node[0], node[1]

					for _, d := range directions {
						nx, ny := x+d[0], y+d[1]
						if nx >= 0 && ny >= 0 && nx < rows && ny < cols && grid[nx][ny] == '1' {
							grid[nx][ny] = '0'
							stack = append(stack, []int{nx, ny})
						}
					}
				}
			}
		}
	}
	return count
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid)) // Выведет 3
	grid1 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslandsIterative(grid1)) // Выведет 3
}
