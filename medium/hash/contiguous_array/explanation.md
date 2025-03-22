# Problem Statement
Given a binary array `nums`, return the maximum length of a contiguous subarray with an equal number of 0 and 1.

## Example
- **Input**: nums = [0, 1]
- **Output**: 2
- **Explanation**: The entire array [0, 1] has equal numbers of 0s and 1s, so the length is 2.

# Approach
The solution involves using a hashmap to store the first occurrence of a cumulative balance of 1s and 0s, where:
1. Replace every `0` in the array with `-1`.
2. Calculate a running `count` which increases by 1 for every `1` encountered and decreases by 1 for every `0` (converted `-1`).
3. The goal is to find the longest subarray which has a cumulative sum (`count`) of zero.
    - This is equivalent to finding two identical counts at different indices, indicating that the subarray between these indices balanced out to zero.

### Steps
1. **Initialize Variables**:
    - `count` to keep track of the balance between 1s and 0s.
    - `maxlen` to record the maximum length of a balanced subarray.
    - `store` hashmap to map each unique `count` value to the earliest index it appears at.
    - Set `store[0] = -1` to handle the case where the balanced subarray starts from index 0.
2. **Iterate Over `nums`**:
    - Update `count` based on the value of `nums[i]`.
    - If the current `count` has been seen before (`store[count]` exists), it means a balanced subarray is found:
        - Calculate the length of the current balanced subarray and update `maxlen` if this is the longest found so far.
    - If the `count` has not been seen before, store its index.
3. **Return `maxlen`**: The maximum length of a contiguous balanced subarray found.

### Helper Function - `max()`
- **Purpose**: Returns the larger of two integers, used to keep track of the longest balanced subarray.

## Complexity Analysis
- **Time Complexity**: O(n), where `n` is the length of `nums`. The algorithm iterates through the array a single time, performing constant time operations for each element.
- **Space Complexity**: O(n), due to the hashmap used to store up to `n` different `count` values and their first occurrences.

# Code
```go
package main

import "fmt"

func findMaxLength(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    count, maxlen := 0, 0
    store := make(map[int]int)
    store[0] = -1
    for i, v := range nums {
        if v == 1 {
            count++
        } else {
            count--
        }
        if v, ok := store[count]; ok {
            maxlen = max(maxlen, i - v)
        } else {
            store[count] = i
        }
    }
    return maxlen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{0, 1, 0}
    fmt.Println(findMaxLength(nums))  // Output: 2
}
