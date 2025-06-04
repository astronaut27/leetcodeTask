# Problem Statement
Given an integer array `nums`, find the length of the longest strictly increasing subsequence.

## Example
- **Input**: nums = [1, 2, 5, 3, 4]
- **Output**: 4
- **Explanation**: The longest increasing subsequence is [1, 2, 3, 4].

# Approach
The solution uses a dynamic programming approach where:
1. **Initialize an Array `dp`**: A DP array `dp` is initialized where `dp[i]` represents the length of the longest increasing subsequence ending with `nums[i]`.
2. **Set Initial Values**: Each element of `dp` is initialized to `1`, as the smallest LIS ending at each index is the element itself.
3. **Populate the DP Array**: Iterate over each element, comparing it with all previous elements:
    - If `nums[i]` is greater than a previous element `nums[j]`, then `nums[i]` can extend the subsequence ending at `j`.
    - Update `dp[i]` to be the maximum of its current value or `dp[j] + 1` (if it forms a longer increasing subsequence by adding `nums[i]`).
4. **Determine the Maximum Length**: After filling the `dp` array, the maximum value within `dp` indicates the length of the longest increasing subsequence in `nums`.

### Helper Function - `maxInSlice()`
- **Purpose**: Finds the maximum value in an array, used to extract the longest subsequence length from the `dp` array.

## Complexity Analysis
- **Time Complexity**: O(n^2), where `n` is the number of elements in `nums`. Each element is compared with all previous elements, resulting in a quadratic time complexity.
- **Space Complexity**: O(n) for storing the DP array.

# Code
```go
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
	fmt.Println(lengthOfLIS(nums)) // Output: 4
}
