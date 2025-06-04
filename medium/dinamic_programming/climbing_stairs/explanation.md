# Problem Statement
Given a staircase with `n` steps, determine how many distinct ways there are to climb to the top of the staircase. You can climb either 1 step or 2 steps at a time.

## Example
- **Input**: n = 3
- **Output**: 3
- **Explanation**: There are three ways to climb to the top:
    1. 1 step + 1 step + 1 step
    2. 1 step + 2 steps
    3. 2 steps + 1 step

# Approach
This problem is a classic example of dynamic programming, often referred to as the "staircase problem". It is similar to the Fibonacci sequence, where the number of ways to reach the `n`th step is the sum of the ways to reach the `(n-1)`th step and the `(n-2)`th step:
1. **Base Cases**:
    - If `n == 1`, there is only 1 way to climb the staircase.
2. **Dynamic Programming Initialization**:
    - Use a DP array `dp` where `dp[i]` represents the number of ways to climb to the `i`th step.
    - Set `dp[0]` to 1 and `dp[1]` to 2, as there is 1 way to climb 1 step and 2 ways to climb 2 steps.
3. **Fill DP Array**:
    - For each step from 2 to `n`, calculate the number of ways to get there by summing the ways to get to the two previous steps:
      ```
      dp[i] = dp[i-1] + dp[i-2]
      ```
4. **Result**:
    - The answer for `n` steps is stored in `dp[n-1]`.

## Complexity Analysis
- **Time Complexity**: O(n), as we need to fill the DP array with `n` elements.
- **Space Complexity**: O(n), due to the space used by the DP array.

# Code
```go
package main

import "fmt"

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    dp := make([]int, n)
    dp[0], dp[1] = 1, 2
    for i := 2; i < n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n-1]
}

func main() {
    fmt.Println(climbStairs(3))  // Output: 3
}
