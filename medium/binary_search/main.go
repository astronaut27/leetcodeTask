package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("final", minEatingSpeed([]int{3, 6, 7, 11}, 8))
}

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, bananas := range piles {
		right = max(right, bananas)
	}
	for left <= right {
		mid := left + (right-left)/2
		if check(piles, mid, h) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(piles []int, speed, limit int) bool {
	var hours float64
	for _, bananas := range piles {
		hours += math.Ceil(float64(bananas) / float64(speed))
	}
	return hours <= float64(limit)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
