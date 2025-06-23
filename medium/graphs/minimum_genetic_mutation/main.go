package main

import (
	"fmt"
)

// minMutation находит минимальное количество мутаций от start до end
func minMutation(start string, end string, bank []string) int {
	// Преобразуем bank в map для быстрого доступа
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}
	// Если конечного гена нет в bank — пути не существует
	if !bankSet[end] {
		return -1
	}

	// BFS
	queue := []string{start}
	visited := map[string]bool{start: true}
	steps := 0
	genes := []byte{'A', 'C', 'G', 'T'}

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]

			if curr == end {
				return steps
			}

			currBytes := []byte(curr)
			for i := 0; i < len(currBytes); i++ {
				original := currBytes[i]
				for _, g := range genes {
					currBytes[i] = g
					next := string(currBytes)
					if bankSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
				currBytes[i] = original // откат буквы
			}
		}
		steps++
	}

	return -1 // не удалось найти путь
}

func main() {
	start := "AACCGGTT"
	end := "AAACGGTA"
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}

	result := minMutation(start, end, bank)
	fmt.Println("Minimum mutations:", result) // Ожидается: 2

	// Дополнительный пример
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"})) // 1
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA"})) // -1
}
