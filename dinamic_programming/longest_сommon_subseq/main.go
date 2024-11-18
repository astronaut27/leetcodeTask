package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	// Создаем двумерный массив dp размером (n+1) x (m+1), инициализированный нулями
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Заполняем dp снизу вверх
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1] // символы совпадают
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1]) // выбираем максимальный из двух вариантов
			}
		}
	}

	return dp[0][0]
}

// Вспомогательная функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace")) // Ожидаемый результат: 3
}
