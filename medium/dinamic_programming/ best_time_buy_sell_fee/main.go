package main

import (
	"fmt"
)

func maxProfitWithFee(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// Создаем массивы для dp
	dp0 := make([]int, n) // dp0[i]: максимальная прибыль, если не держим акцию на день i
	dp1 := make([]int, n) // dp1[i]: максимальная прибыль, если держим акцию на день i

	// Инициализация первого дня
	dp0[0] = 0          // В первый день без акций прибыль равна 0
	dp1[0] = -prices[0] // Если купили акцию в первый день, прибыль отрицательна

	// Заполняем массивы
	for i := 1; i < n; i++ {
		dp0[i] = max(dp0[i-1], dp1[i-1]+prices[i]-fee) // Либо продолжаем не держать, либо продаем акцию
		dp1[i] = max(dp1[i-1], dp0[i-1]-prices[i])     // Либо продолжаем держать, либо покупаем
	}

	// Максимальная прибыль в состоянии "не держим акцию" на последний день
	return dp0[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfitWithFee(prices, fee)) // Ожидаемый результат: 8
}
