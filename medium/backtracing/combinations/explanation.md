# Intuition
The problem involves generating all possible combinations of `k` numbers out of a range from `1` to `n`. This classic combinatorial problem can be efficiently tackled using a backtracking approach. The idea is to construct each combination step-by-step, choosing an element, including it in the current combination, and then moving forward to include the next elements until the combination reaches the desired size (`k`).

# Approach
1. **Recursive Backtracking**:
    - Start from `1` and continue to `n`, at each step deciding whether to include the current number in the ongoing combination (`curr`).
    - Use a helper function `backtrack` to manage the state of each combination as it's being constructed.

2. **Backtracking Logic**:
    - If the current combination has `k` elements, add it to the final results (`ans`).
    - Otherwise, iterate from the current index up to `n`. For each number, add it to the current combination and recurse to fill the next elements. After returning from recursion, backtrack by removing the last added element to try the next possibility.

3. **Base Case and Recursive Case**:
    - **Base Case**: When the combination's length equals `k`, add it to the result list.
    - **Recursive Case**: Recurse with the next element until the entire range is explored.

# Complexity
- **Time complexity**: O(C(n, k)), where `C(n, k)` is the binomial coefficient representing the number of ways to choose `k` elements from `n` without regard to order. This is the number of combinations generated.
- **Space complexity**: O(C(n, k) * k), for storing each combination. Additionally, O(k) space is used for the recursion stack.


### Explanation
- **Backtracking Function**: The `backtrack` function constructs combinations by recursively adding elements to `curr`. When `curr` reaches the desired size (`k`), it is added to the list of answers.
- **Iteration and Recursion**: The function iterates over the range `[index, n]`, adding each number to the current combination, then recursing to add subsequent numbers. After each recursive call, it backtracks by removing the last number added, allowing it to try the next number in the sequence.
- **Optimizations**: The recursive calls efficiently navigate the solution space, avoiding unnecessary work by only considering numbers that maintain the increasing order, ensuring all combinations are unique.

This structure outlines the approach in a structured way, detailing the intuition, approach, complexity analysis, and providing a complete code example.

# Code
```go
package main

import "fmt"

func combine(n int, k int) [][]int {
    var ans [][]int
    backtrack([]int{}, 1, n, k, &ans)
    return ans
}

func backtrack(curr []int, index int, n int, k int, ans *[][]int) {
    if len(curr) == k { // Base case: length k is reached
        *ans = append(*ans, append([]int{}, curr...))
        return
    }
    for i := index; i <= n; i++ {
        curr = append(curr, i)
        backtrack(curr, i+1, n, k, ans)
        curr = curr[:len(curr)-1]  // Backtrack
    }
}

func main() {
    n := 4
    k := 2
    result := combine(n, k)
    fmt.Println("All combinations:", result)
}
```