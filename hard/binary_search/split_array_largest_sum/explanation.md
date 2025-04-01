# Problem Statement
Given an array `nums` of non-negative integers and an integer `k`, the task is to split the array into `k` non-empty contiguous subarrays such that the largest sum among these subarrays is minimized.

## Example
- **Input**: nums = [7,2,5,10,8], k = 2
- **Output**: 18
- **Explanation**: The optimal way to split the array into 2 parts is [7,2,5] and [10,8]. The maximum sum of these subarrays is 18, which is the minimal possible.

# Approach
The solution involves using a binary search to find the minimum possible maximum subarray sum for `k` partitions:
1. **Define Search Range**: Set `low` to the maximum value in `nums` (smallest possible maximum subarray sum), and `high` to the sum of all elements in `nums` (largest possible maximum subarray sum).
2. **Binary Search**:
    - Compute the middle of the current range (`mid`).
    - Use a helper function `canSplit()` to check if it is possible to partition `nums` into at most `k` subarrays where the sum of each subarray does not exceed `mid`.
    - If true, it means we might be able to reduce `mid`, so adjust `high`.
    - If false, increase `low` to search for a larger possible `mid`.
3. The smallest `mid` for which `canSplit()` returns true is the answer.

### Helper Function - `canSplit()`
- **Purpose**: Checks if it is possible to partition the array into at most `k` subarrays such that no subarray sum exceeds `mid`.
- **Implementation**: Iterate through `nums`, aggregating sums and forming new subarrays whenever the current subarray sum exceeds `mid`.

## Complexity Analysis
- **Time Complexity**: O(n log S), where `n` is the number of elements in `nums`, and `S` is the sum of all elements minus the largest single element.
- **Space Complexity**: O(1), as the solution uses a constant amount of space.

# Code
```go
package main

import "fmt"

func splitArray(nums []int, k int) int {
    var low, high int
    for i := range nums {
        low = max(low, nums[i])
        high += nums[i]
    }
    for low <= high {
        mid := low + (high-low)/2
        if canSplit(nums, mid, k) {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return low
}

func canSplit(nums []int, mid, k int) bool {
    currentSum, subarrays := 0, 1
    for i := range nums {
        currentSum += nums[i]
        if currentSum > mid {
            subarrays++
            currentSum = nums[i]
            if subarrays > k {
                return false
            }
        }
    }
    return true
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func main() {
    nums := []int{7, 2, 5, 10, 8}
    k := 2
    fmt.Println(splitArray(nums, k))  // Output: 18
}
```