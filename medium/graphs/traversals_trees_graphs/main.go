package main

import "fmt"

type Point struct {
	x, y int
}

var directions = []Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func nearestExit(maze [][]byte, entrance []int) int {
	n, m := len(maze), len(maze[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	visited[entrance[0]][entrance[1]] = true
	return bfs(maze, entrance[0], entrance[1], visited)
}

func bfs(maze [][]byte, startX, startY int, visited [][]bool) int {
	n, m := len(maze), len(maze[0])
	queue := []Point{{startX, startY}}
	steps := 0

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]

			x, y := curr.x, curr.y
			if maze[x][y] == '.' &&
				(x == 0 || x == n-1 || y == 0 || y == m-1) &&
				!(x == startX && y == startY) {
				return steps
			}

			for _, d := range directions {
				nx, ny := x+d.x, y+d.y
				if nx >= 0 && ny >= 0 && nx < n && ny < m &&
					maze[nx][ny] == '.' && !visited[nx][ny] {
					visited[nx][ny] = true
					queue = append(queue, Point{nx, ny})
				}
			}
		}
		steps++
	}
	return -1
}

func main() {
	fmt.Println(nearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}))
	fmt.Println(nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}))
}
