package main

import "fmt"

func maximizeSweetness(sweetness []int, k int) int {
	low, high := 1, 0
	for _, sweet := range sweetness {
		high += sweet
	}
	for low <= high {
		mid := low + (high-low)/2
		if canDivide(sweetness, mid, k) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

func canDivide(sweetness []int, mid, k int) bool {
	currentSum, pieces := 0, 0
	for _, sweet := range sweetness {
		currentSum += sweet
		if currentSum >= mid {
			pieces += 1
			currentSum = 0
		}
	}
	return pieces >= k+1
}

func main() {
	sweets := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 5
	fmt.Println(maximizeSweetness(sweets, k)) // Output should be 6
}
