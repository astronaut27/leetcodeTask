package main

import "fmt"

func main() {
	fmt.Println(fibonacci(6))
}

// first approach--"top-down" dynamic programming.
var fibonacciMap = make(map[int]int)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if v, ok := fibonacciMap[n]; ok {
		return v
	}

	fibonacciMap[n] = fibonacci(n-1) + fibonacci(n-2)
	return fibonacciMap[n]
}

// "Bottom-Up"
func fibonacciBottomUp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}
