# Problem Statement
Given an array of coin denominations `coins` and an integer `amount`, compute the fewest number of coins needed to make up that amount. If it's not possible to make up that amount with the given coins, return `-1`.

## Example
- **Input**: coins = [1, 2, 5], amount = 11
- **Output**: 3
- **Explanation**: 11 = 5 + 5 + 1

# Approach
The problem is solved using a dynamic programming approach where:
1. **DP Array Initialization**:
    - Create an array `dp` of size `amount + 1` where `dp[i]` represents the minimum number of coins needed to make the amount `i`.
    - Initialize each element to `amount + 1` (a value larger than any possible number of coins required) to signify that the amount isn't achievable initially, except for `dp[0]` which is `0` (zero coins are needed to make amount `0`).
2. **Fill DP Array**:
    - For each sub-amount `i` from `1` to `amount`, iterate through each coin:
        - If the coin value is less than or equal to `i`, update `dp[i]` to be the minimum of its current value or `dp[i - coin] + 1`.
3. **Final Check**:
    - After filling the `dp` array, `dp[amount]` gives the minimum coins needed for the `amount`. If `dp[amount]` is still greater than `amount`, it means the amount isn't achievable with the given coins, hence return `-1`.

### Helper Function - `min()`
- **Purpose**: Finds the minimum of two numbers, aiding in updating the DP table with the least number of coins needed.

## Complexity Analysis
- **Time Complexity**: O(n * m), where `n` is the `amount` and `m` is the number of coin denominations. The nested loop structure results in this complexity because for each amount, the algorithm considers each coin denomination.
- **Space Complexity**: O(n), due to the space used by the DP array `dp` which has a length of `amount + 1`.

# Code
```go
package main

import "fmt"

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0
    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if coin <= i {
                dp[i] = min(dp[i], dp[i-coin] + 1)
            }
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    coins := []int{1, 2, 5}
    amount := 11
    fmt.Println(coinChange(coins, amount))  // Output: 3
}
