# Intuition
The problem requires us to find the highest altitude a biker reaches given an array of altitude gains between checkpoints. Since the starting altitude is 0, we can simulate the altitude after each gain and track the maximum.

# Approach
We initialize the current altitude (`tmp`) and the maximum altitude (`result`) to 0. Then, for each gain value in the input array:
- Add the gain to the current altitude.
- Update the result if the current altitude exceeds it.

This approach mimics a running total, and by tracking the maximum along the way, we can determine the highest point reached.

# Complexity
- Time complexity:  
  $$O(n)$$ — where \( n \) is the length of the `gain` array. We iterate through the array once.

- Space complexity:  
  $$O(1)$$ — constant space is used for tracking altitude and the maximum.

# Code
```golang []
func largestAltitude(gain []int) int {
    var tmp, result int
    for i := 0; i < len(gain); i++ {
        tmp += gain[i]
        result = max(result, tmp)
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
