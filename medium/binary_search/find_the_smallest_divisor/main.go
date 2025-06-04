package main

import "math"

func main() {

}

var limit int

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, 0
	for _, num := range nums {
		right = max(num, right)
	}
	for left <= right {
		mid := left + (right-left)/2
		if check(nums, mid, threshold) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(nums []int, divisor, limit int) bool {
	result := 0.0
	for _, num := range nums {
		result += math.Ceil(float64(num) / float64(divisor))
	}
	return int(result) <= limit
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
