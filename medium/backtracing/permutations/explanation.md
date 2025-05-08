# Intuition
The problem of generating all permutations of a given array of integers can be approached using backtracking. The idea is to create every possible permutation by choosing each element as a starting point, then recursively doing the same for the remaining elements until all positions in the permutation have been filled.

# Approach
The solution involves a recursive function `backtrack` that constructs permutations step by step:
1. **Base Case**: If the `nums` list is empty, it means we've constructed a full permutation. We then copy this permutation to the result list.
2. **Recursive Step**: For each element in `nums`, make it the next element in the current permutation, and recursively call the function on the remaining elements (excluding the chosen one).
    - **Removing the Element**: When choosing an element for the current position of the permutation, remove it from the available list of numbers by slicing before and after the chosen index and combining those slices.
    - **Combining Current with Chosen Element**: Append the chosen element to the current permutation and recurse with the updated current permutation and the reduced list of remaining numbers.
3. **Collect Results**: Every time a full permutation is constructed (when `nums` is empty), append a copy of it to the result list.

# Complexity
- **Time complexity**: O(n!), where n is the length of the input list `nums`. This complexity arises because there are n! (factorial of n) permutations of n items.
- **Space complexity**: O(n), this space is used to store the recursion stack. Although we create multiple permutations, the stack height will only go up to n, where n is the number of elements in the input array. However, the total space for the output (all permutations stored) is not included in this analysis.

# Code
```go
package main

import "fmt"

func permute(nums []int) [][]int {
    var result [][]int
    backtrack([]int{}, nums, &result)
    return result
}

func backtrack(curr []int, nums []int, result *[][]int) {
    if len(nums) == 0 { // Base case: no more elements to add
        *result = append(*result, append([]int{}, curr...)) // Make a copy of curr and add to result
        return
    }
    
    for i := 0; i < len(nums); i++ {
        // Form a new list excluding the current element
        nextNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
        backtrack(append(curr, nums[i]), nextNums, result) // Recur with the new current and remaining elements
    }
}

func main() {
    nums := []int{1, 2, 3}
    result := permute(nums)
    fmt.Println("All permutations:", result)
}
