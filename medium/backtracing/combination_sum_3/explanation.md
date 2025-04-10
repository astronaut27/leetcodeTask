# Intuition
The problem asks to find all unique combinations where `k` distinct numbers sum up to `n`. Each number from 1 to 9 can be used at most once. This problem is suitable for a backtracking approach, which allows us to explore all possible combinations and backtrack when a path doesn't lead to a solution.

# Approach
1. **Recursive Backtracking**:
    - Use a recursive function to build each combination by adding one number at a time and moving to the next potential candidate.
    - Start from an index to ensure numbers are not reused and the combination remains unique.

2. **Base Case and Recursive Calls**:
    - **Base Case**: Check if the combination's length is `k` and its sum is `n`. If true, add the current combination to the result.
    - **Recursive Case**: Iterate through the numbers starting from the current index. For each number:
        - Add it to the current combination.
        - Recursively attempt to find valid combinations starting from the next number.
        - Backtrack by removing the last added number to try the next possibility.

3. **Efficient Sum Calculation**:
    - Maintain the current sum instead of recalculating it by passing the sum as an argument to the recursive function.

# Complexity
- **Time Complexity**: O(9! / (9-k)!), which approximates the number of ways to pick `k` distinct numbers from 1 to 9, given that the order of selection matters because we terminate early when the sum exceeds `n`.
- **Space Complexity**: O(k) due to the recursive call stack and the space used to store the current combination.

### Explanation
- **Backtracking Function**: The `backtrack` function tries each number from `index` to 9, ensuring each number is used at most once. It recursively builds up a combination of numbers and checks if it meets the requirements.
- **Validation**: If the current combination has the desired length `k` and sums to `n`, it's added to the result list.
- **Efficiency**: The function avoids unnecessary work by stopping early if the current sum exceeds `n` or if the combination length exceeds `k`.

This backtracking approach is efficient and straightforward, exploiting the problem constraints (use of numbers 1-9 and each number at most once) to simplify the search space.

# Code
```go
package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
    var ans [][]int
    backtrack([]int{}, 1, n, k, &ans)
    return ans
}

func backtrack(curr []int, index int, n int, k int, ans *[][]int) {
    if len(curr) == k && sum(curr) == n { // Check if the combination is valid
        *ans = append(*ans, append([]int{}, curr...)) // Copy the valid combination
        return
    }
    if len(curr) > k || sum(curr) > n { // Early termination
        return
    }
    for i := index; i < 10; i++ {
        curr = append(curr, i)
        backtrack(curr, i+1, n, k, ans) // Recurse with the next number
        curr = curr[:len(curr)-1] // Remove the last number (backtrack)
    }
}

func sum(curr []int) int { // Helper function to calculate the sum of elements in curr
    var ans int
    for _, num := range curr {
        ans += num
    }
    return ans
}

func main() {
    fmt.Println("Combinations sum to 9, pick 3:", combinationSum3(3, 9))
}
```