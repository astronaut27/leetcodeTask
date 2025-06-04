package main

import (
	"fmt"
)

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)

	// Построение графа: b[i] -> b[j] если j в радиусе действия i
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		x1, y1, r := bombs[i][0], bombs[i][1], bombs[i][2]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x2, y2 := bombs[j][0], bombs[j][1]
			dx, dy := x2-x1, y2-y1
			if dx*dx+dy*dy <= r*r {
				graph[i] = append(graph[i], j)
			}
		}
	}

	maxCount := 0
	for i := 0; i < n; i++ {
		visited := make(map[int]bool)
		count := dfs(i, graph, visited)
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

// 📦 Отдельная функция DFS
func dfs(node int, graph map[int][]int, visited map[int]bool) int {
	if visited[node] {
		return 0
	}
	visited[node] = true
	count := 1
	for _, nei := range graph[node] {
		count += dfs(nei, graph, visited)
	}
	return count
}

func main() {
	fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}})) // 2
	fmt.Println(maximumDetonation([][]int{{1, 1, 5}, {10, 10, 5}})) // 1
	fmt.Println(maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}})) // 5
}
