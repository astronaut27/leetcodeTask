# Problem Statement
Given an integer array `nums` representing the amount of money in each house, determine the maximum amount of money that can be robbed without robbing two adjacent houses. The houses are arranged in a straight line.

## Example
- **Input**: nums = [2,3,2]
- **Output**: 4
- **Explanation**: Robbing house 1 (amount = 2) and house 3 (amount = 2) yields the maximum profit of 4.

## Approach
The solution uses dynamic programming to calculate the maximum amount that can be robbed up to each house without triggering an alarm:
1. **Base Cases**:
    - If there are no houses (`len(nums) == 0`), return 0.
    - If there is only one house (`len(nums) == 1`), return the amount in that house.
2. **Dynamic Programming (DP) Array**:
    - Initialize a DP array `dp` where `dp[i]` represents the maximum amount that can be robbed up to house `i`.
    - Set `dp[0]` to `nums[0]` (robbing the first house).
    - Set `dp[1]` to `max(nums[0], nums[1])` (choosing between the first or second house).
3. **Fill DP Array**:
    - For each house `i` from 2 to `len(nums)-1`, determine whether to rob the current house and add it to `dp[i-2]` or skip it and take the maximum up to `dp[i-1]`. This choice is represented by the formula:
      ```
      dp[i] = max(dp[i-1], dp[i-2] + nums[i])
      ```
4. **Result**:
    - The maximum money that can be robbed is the value at `dp[len(nums) - 1]`.

### Helper Function - `max()`
- **Purpose**: Returns the maximum of two integers. Used for selecting the optimal amount to rob between two choices.

## Complexity Analysis
- **Time Complexity**: O(n), where `n` is the number of houses. The algorithm iterates through `nums` once to fill the DP array.
- **Space Complexity**: O(n), for storing the DP array that tracks the maximum money that can be robbed up to each house.

# Code
```go
package main

import "fmt"

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
        dp = append(dp, max(dp[i-1], dp[i-2] + nums[i]))
    }
    return dp[len(nums)-1]
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func main() {
    nums := []int{2, 3, 2}
    fmt.Println(rob(nums))  // Output: 4
}
