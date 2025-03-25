# Problem Statement
Given an array `cost` where `cost[i]` is the cost of the `i`th step on a staircase, determine the minimum cost to reach the top of the staircase. You can either start from the step with index 0 or index 1, and you can climb one or two steps at a time.

## Example
- **Input**: cost = [10, 15, 20]
- **Output**: 15
- **Explanation**: The cheapest way to the top is to start on cost[1], pay 15, and then jump to the top.

## Approach
The solution involves dynamic programming to determine the minimum cost to reach each step. We build upon the smaller subproblems to solve for the larger problems:
1. **DP Array Initialization**:
    - Create a DP array `dp` of size `n+1` where `n` is the number of steps in `cost`. This accounts for the top of the staircase being a step beyond the last element in `cost`.
    - Initialize `dp[0]` and `dp[1]` to `0` because it costs nothing to start at these steps.
2. **Fill DP Array**:
    - For each step `i` from 2 to `n` (inclusive), calculate the minimum cost to reach that step:
        - The cost to reach step `i` from step `i-1` plus the cost at step `i-1`.
        - The cost to reach step `i` from step `i-2` plus the cost at step `i-2`.
    - Use the `min` function to decide the lesser of the two options.
3. **Result**:
    - The value at `dp[n]` gives the minimum cost required to reach the top of the staircase.

### Helper Function - `min()`
- **Purpose**: Returns the smaller of two integers, helping to select the minimum cost between two potential paths.

## Complexity Analysis
- **Time Complexity**: O(n), where `n` is the number of steps in the staircase. Each step calculation involves constant time operations.
- **Space Complexity**: O(n), for the DP array that holds the minimum cost up to each step.

# Code
```go
package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp := make([]int, n+1)
    
    dp[0], dp[1] = 0, 0
    
    for i := 2; i <= n; i++ {
        dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
    }
    
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    cost := []int{10, 15, 20}
    fmt.Println(minCostClimbingStairs(cost)) // Output: 15
}
