package main

func main() {
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, 0, len(nums))
	dp = append(dp, nums[0])
	dp = append(dp, max(nums[0], nums[1]))
	for i := 2; i < len(nums); i++ {
		dp = append(dp, max(dp[i-1], dp[i-2]+nums[i]))
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
