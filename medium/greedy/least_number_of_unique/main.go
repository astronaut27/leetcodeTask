package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	counts := make(map[int]int)

	// Подсчитываем количество вхождений каждого числа
	for _, num := range arr {
		counts[num]++
	}

	// Создаем срез для хранения количества уникальных чисел
	var uniqueCounts []int
	for _, count := range counts {
		uniqueCounts = append(uniqueCounts, count)
	}

	// Сортируем по возрастанию
	sort.Ints(uniqueCounts)

	// Удаляем элементы, чтобы минимизировать количество уникальных чисел
	remainingUnique := len(uniqueCounts)
	for _, count := range uniqueCounts {
		if k >= count {
			k -= count
			remainingUnique-- // Удаляем одно уникальное число
		} else {
			break
		}
	}

	// Количество оставшихся уникальных чисел
	return remainingUnique
}

func main() {
	arr := []int{4, 3, 1, 1, 1, 3, 2}
	k := 3
	result := findLeastNumOfUniqueInts(arr, k)
	fmt.Println(result) // Вывод: 2
}
