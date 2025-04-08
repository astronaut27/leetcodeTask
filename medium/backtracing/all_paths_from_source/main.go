package main

import "fmt"

// Основная функция для поиска всех путей
func allPathsSourceTarget(graph [][]int) [][]int {
	return backtrack(0, []int{0}, graph) // Начинаем с узла 0 и пути [0]
}

// Самодостаточная функция Backtracking
func backtrack(node int, curPath []int, graph [][]int) [][]int {
	// Если достигли конечной вершины (n-1), возвращаем текущий путь как результат
	if node == len(graph)-1 {
		return [][]int{append([]int{}, curPath...)} // Копируем путь
	}

	var paths [][]int
	// Проходим по всем соседям текущего узла
	for _, child := range graph[node] {
		// Рекурсивно вызываем backtrack для соседа
		newPaths := backtrack(child, append(curPath, child), graph)
		paths = append(paths, newPaths...) // Добавляем новые пути в общий результат
	}
	return paths
}

func main() {
	graph1 := [][]int{{1, 2}, {3}, {3}, {}}
	graph2 := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}

	fmt.Println(allPathsSourceTarget(graph1)) // [[0,1,3],[0,2,3]]
	fmt.Println(allPathsSourceTarget(graph2)) // [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
}
