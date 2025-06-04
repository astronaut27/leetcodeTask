package main

import (
	"fmt"
	"strconv"
)

func numsSameConsecDiff(n int, k int) []int {
	var ans []int
	backtrack([]int{}, n, k, &ans)
	return ans
}

// Основная функция backtrack
func backtrack(curr []int, size, k int, ans *[]int) {
	// Базовый случай: когда длина текущей комбинации равна size
	if len(curr) == size {
		// Преобразуем комбинацию в число и добавляем в результат
		num, _ := strconv.Atoi(join(curr))
		*ans = append(*ans, num)
		return
	}

	// Перебор возможных цифр для добавления
	// Если это первая цифра, она не может быть 0
	start := 0
	if len(curr) == 0 {
		start = 1
	}

	for i := start; i < 10; i++ {
		// Если первая цифра, то её разница с соседними должна быть k
		if len(curr) == 0 || abs(curr[len(curr)-1]-i) == k {
			// Добавляем цифру в текущую комбинацию
			curr = append(curr, i)
			// Рекурсивно вызываем backtrack для следующей позиции
			backtrack(curr, size, k, ans)
			// Откатываем изменения
			curr = curr[:len(curr)-1]
		}
	}
}

// Преобразуем срез целых чисел в строку
func join(nums []int) string {
	var result string
	for _, num := range nums {
		result += strconv.Itoa(num)
	}
	return result
}

// Абсолютное значение
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
}
