package main

import (
	"fmt"
)

// Структура соседа с весом (ребро графа)
type neighbor struct {
	node   string
	weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Граф: переменная → список соседей с коэффициентами
	graph := make(map[string][]neighbor)

	// Построение графа из уравнений
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		val := values[i]

		// Добавляем два ребра: a→b (val) и b→a (1/val)
		graph[a] = append(graph[a], neighbor{b, val})
		graph[b] = append(graph[b], neighbor{a, 1.0 / val})
	}

	results := make([]float64, 0, len(queries))

	// Обрабатываем каждый запрос
	for _, q := range queries {
		start, end := q[0], q[1]

		// Проверка на отсутствие переменных в графе
		_, okStart := graph[start]
		_, okEnd := graph[end]
		if !okStart || !okEnd {
			results = append(results, -1.0)
			continue
		}

		// DFS-поиск пути с накоплением коэффициентов
		visited := make(map[string]bool)
		val, ok := dfs(graph, start, end, 1.0, visited)
		if !ok {
			results = append(results, -1.0)
		} else {
			results = append(results, val)
		}
	}

	return results
}

// DFS-поиск пути от curr до target с накоплением коэффициента acc
func dfs(graph map[string][]neighbor, curr, target string, acc float64, visited map[string]bool) (float64, bool) {
	if curr == target {
		return acc, true
	}

	visited[curr] = true

	for _, n := range graph[curr] {
		if !visited[n.node] {
			// Рекурсивный шаг: продолжаем путь
			if res, ok := dfs(graph, n.node, target, acc*n.weight, visited); ok {
				return res, true
			}
		}
	}

	return -1.0, false
}

func main() {
	// Пример 1
	eq := [][]string{{"a", "b"}, {"b", "c"}}
	val := []float64{2.0, 3.0}
	q := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(eq, val, q))
	// Ожидаемый результат: [6.0, 0.5, -1.0, 1.0, -1.0]
}
