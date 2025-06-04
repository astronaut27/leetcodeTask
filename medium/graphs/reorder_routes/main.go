package main

import "fmt"

func minReorder(n int, connections [][]int) int {
	// Создаем список смежности
	graph := make([][]int, n)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		graph[a] = append(graph[a], b)  // Оригинальное направление
		graph[b] = append(graph[b], -a) // Обратное направление
	}

	// Используем BFS для обхода графа
	visited := make([]bool, n)
	queue := []int{0}
	visited[0] = true
	changes := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[node] {
			if neighbor < 0 {
				neighbor = -neighbor
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			} else {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
					changes++
				}
			}
		}
	}

	return changes
}

func main() {
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	n := 6
	fmt.Println("Minimum reorder required:", minReorder(n, connections))
}
