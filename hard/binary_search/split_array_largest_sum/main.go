package main

import "fmt"

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))

}

func splitArray(nums []int, k int) int {
	var low, high int
	for i := range nums {
		low = max(low, nums[i])
		high += nums[i]
	}
	for low <= high {
		mid := low + (high-low)/2
		if canSplit(nums, mid, k) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func canSplit(nums []int, mid, k int) bool {
	currentSum, subarrays := 0, 1
	for i := range nums {
		currentSum += nums[i]
		if currentSum > mid {
			subarrays++
			currentSum = nums[i]
			if subarrays > k {
				return false
			}
		}
	}
	return true
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
