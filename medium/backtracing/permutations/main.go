package main

import "fmt"

// Основная функция для генерации перестановок
func permute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	var curr []int

	// Вызов вспомогательной функции backtrack
	backtrack(nums, used, curr, &result)
	return result
}

// Вспомогательная функция Backtracking
func backtrack(nums []int, used []bool, curr []int, result *[][]int) {
	// Базовый случай: перестановка завершена
	if len(curr) == len(nums) {
		*result = append(*result, append([]int{}, curr...)) // Сохраняем копию текущей комбинации
		return
	}

	// Перебираем все элементы массива
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue // Пропускаем уже использованные элементы
		}

		// Добавляем элемент в текущую перестановку
		curr = append(curr, nums[i])
		used[i] = true

		// Рекурсивно вызываем Backtracking для следующего шага
		backtrack(nums, used, curr, result)

		// Откатываем изменения
		curr = curr[:len(curr)-1]
		used[i] = false
	}
}

func permute1(nums []int) [][]int {
	var ans [][]int
	backtrack1([]int{}, &ans, nums)
	return ans
}

func backtrack1(curr []int, ans *[][]int, nums []int) {
	// Базовый случай: если длина curr равна длине nums
	if len(curr) == len(nums) {
		*ans = append(*ans, append([]int{}, curr...)) // Создаём копию текущего пути
		return
	}

	// Перебираем все числа в nums
	for _, num := range nums {
		if contains(curr, num) {
			continue // Пропускаем, если число уже в текущем пути
		}
		// Добавляем число в текущий путь
		curr = append(curr, num)
		backtrack1(curr, ans, nums)
		// Удаляем последнее число для отката
		curr = curr[:len(curr)-1]
	}
}

// Вспомогательная функция для проверки, содержится ли число в срезе
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
