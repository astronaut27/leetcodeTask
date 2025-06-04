package main

import (
	"fmt"
)

func openLock(deadends []string, target string) int {
	// Множество запретных состояний
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	// Если начальное или целевое состояние заблокировано — сразу -1
	if dead["0000"] {
		return -1
	}
	if target == "0000" {
		return 0
	}

	// Инициализация BFS
	visited := make(map[string]bool)
	queue := []string{"0000"}
	visited["0000"] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == target {
				return steps
			}

			// Генерируем все возможные следующие состояния
			for _, next := range getNextStates(curr) {
				if !visited[next] && !dead[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		steps++
	}

	return -1
}

// Генерация всех возможных комбинаций поворота на 1 шаг
func getNextStates(state string) []string {
	result := []string{}
	bytes := []byte(state)

	for i := 0; i < 4; i++ {
		original := bytes[i]

		// Поворот вперед
		bytes[i] = (original-'0'+1)%10 + '0'
		result = append(result, string(bytes))

		// Поворот назад
		bytes[i] = (original-'0'+9)%10 + '0'
		result = append(result, string(bytes))

		// Возвращаем оригинальное значение
		bytes[i] = original
	}

	return result
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))                         // 6
	fmt.Println(openLock([]string{"8888"}, "0009"))                                                         // 1
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")) // -1
}
