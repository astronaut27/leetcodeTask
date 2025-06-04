package main

import (
	"fmt"
)

func canReach(arr []int, start int) bool {
	n := len(arr)
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Условие успеха — достигли значения 0
		if arr[curr] == 0 {
			return true
		}

		// Переходы
		nextPos := []int{curr + arr[curr], curr - arr[curr]}
		for _, next := range nextPos {
			if next >= 0 && next < n && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5)) // true
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0)) // true
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))       // false
}
