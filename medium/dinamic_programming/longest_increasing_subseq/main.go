package main

import "fmt"

func lengthOfLIS(nums []int) int {
	numsLen := len(nums)

	dp := make([]int, numsLen)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < numsLen; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return maxInSlice(dp)
}

func maxInSlice(nums []int) int {
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	nums := []int{1, 2, 5, 3, 4}
	fmt.Println(lengthOfLIS(nums)) // Вывод: 4
}
