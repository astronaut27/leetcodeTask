# Intuition
The problem asks us to determine for each child whether they can have the greatest number of candies if given some extra candies. The intuitive idea is to first find the maximum number of candies any child currently has, and then for each child, check if adding the extra candies makes them reach or exceed this maximum.

# Approach
1. **Find the maximum number of candies** among all children using a simple loop.
2. Create a result slice of booleans with the same length as the `candies` array.
3. For each child, check whether their candies plus the `extraCandies` is greater than or equal to the maximum number found in step 1.
    - If yes, mark the result as `true`.
    - Otherwise, mark it as `false`.
4. Return the result array.

# Complexity
- Time complexity:  
  $$O(n)$$ — One pass to find the maximum and another pass to check each child.

- Space complexity:  
  $$O(n)$$ — For the output boolean slice of the same size as the input.

# Code
```golang []
func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxi, result := 0, make([]bool, len(candies))

    // Find the maximum number of candies
    for _, v := range candies {
        maxi = max(maxi, v)
    }

    // Check if each child can reach or exceed the maximum
    for i, v := range candies {
        if maxi <= v + extraCandies {
            result[i] = true
        } else {
            result[i] = false
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
