# Problem Statement
Given an array of integers `prices` where `prices[i]` represents the price of a given stock on the `i`th day, and an integer `fee` representing a transaction fee for each sale, calculate the maximum profit you can achieve. You can complete as many transactions as you like, but you need to pay the transaction fee for each sale.

## Example
- **Input**: prices = [1, 3, 2, 8, 4, 9], fee = 2
- **Output**: 8
- **Explanation**: The optimal trading strategy is to buy on day 0, sell on day 3, buy on day 4, and sell on day 5. The total profit is ((8-1)-2) + ((9-4)-2) = 8.

# Approach
The solution involves a dynamic programming approach to track the state of owning or not owning a stock on each day:
1. **State Definitions**:
    - `dp0[i]`: Maximum profit on day `i` if we do not hold a stock.
    - `dp1[i]`: Maximum profit on day `i` if we hold a stock.
2. **Base Cases**:
    - `dp0[0]` = 0 (starting without any stocks).
    - `dp1[0]` = -prices[0] (buying the stock on the first day, hence a negative profit initially).
3. **State Transitions**:
    - `dp0[i]` can transition from `dp0[i-1]` (not holding the stock the previous day) or from `dp1[i-1] + prices[i] - fee` (selling the stock owned the previous day).
    - `dp1[i]` can transition from `dp1[i-1]` (holding the stock from the previous day) or from `dp0[i-1] - prices[i]` (buying a new stock today).
4. **Final Result**:
    - The value of `dp0[n-1]` will represent the maximum profit achievable by the end of the last day without holding any stock.

### Helper Function - `max()`
- **Purpose**: Determines the maximum of two integers, facilitating the choice between two competing strategies (e.g., selling vs. not selling).

## Complexity Analysis
- **Time Complexity**: O(n), where `n` is the number of days. Each day requires a constant amount of computation.
- **Space Complexity**: O(n) due to the space used by the `dp0` and `dp1` arrays. This can be optimized to O(1) by using only two variables to store the states of the previous day.

# Code
```go
package main

import "fmt"

func maxProfit(prices []int, fee int) int {
    n := len(prices)
    if n == 0 {
		return 0
	}
    dp0 := make([]int, n) // Maximum profit not holding stock
    dp1 := make([]int, n) // Maximum profit holding stock
    
    dp0[0] = 0
    dp1[0] = -prices[0]
    for i := 1; i < n; i++ {
        dp0[i] = max(dp0[i-1], dp1[i-1] + prices[i] - fee)
        dp1[i] = max(dp1[i-1], dp0[i-1] - prices[i])
    }
    return dp0[n-1]
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func main() {
    prices := []int{1, 3, 2, 8, 4, 9}
    fee := 2
    fmt.Println(maxProfit(prices, fee)) // Output: 8
}
