package main

import (
	"fmt"
)


// Структура для координат клетки
type Point struct {
	x, y int
}

// Основная функция для нахождения кратчайшего пути
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	// Если старт или конец заблокированы — пути нет
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}

	// Все 8 направлений (включая диагонали)
	directions := []Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
	}

	// Массив для отслеживания посещённых ячеек
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// BFS очередь и инициализация
	queue := []Point{{0, 0}}
	visited[0][0] = true
	pathLength := 1

	// Основной цикл BFS
	for len(queue) > 0 {
		nextQueue := []Point{}

		for _, cell := range queue {
			x, y := cell.x, cell.y

			// Если дошли до цели — возвращаем длину пути
			if x == n-1 && y == n-1 {
				return pathLength
			}

			// Обходим всех соседей
			for _, dir := range directions {
				nx, ny := x+dir.x, y+dir.y

				// 1. Проверка границ
				if nx < 0 || ny < 0 || nx >= n || ny >= n {
					continue
				}

				// 2. Проверка: проходимость и непосещённость
				if grid[nx][ny] != 0 || visited[nx][ny] {
					continue
				}

				// 3. Добавляем в очередь следующего уровня
				visited[nx][ny] = true
				nextQueue = append(nextQueue, Point{nx, ny})
			}
		}

		queue = nextQueue
		pathLength++ // Переход к следующему уровню BFS
	}

	// Если не дошли до конца — пути нет
	return -1
}

func main() {
	grid1 := [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid1)) // Output: 2

	grid2 := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid2)) // Output: 4

	grid3 := [][]int{
		{1, 0},
		{0, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid3)) // Output: -1
}