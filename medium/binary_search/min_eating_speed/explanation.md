# Problem Statement
Given an array of `piles` where each element represents the number of bananas in a pile, and an integer `h` representing the total hours available, determine the minimum integer speed `k` (bananas per hour) at which a person can eat all the bananas within `h` hours.

## Example
- **Input**: piles = [3, 6, 7, 11], h = 8
- **Output**: 4
- **Explanation**: With a speed of 4 bananas per hour, the person takes:
    - 1 hour to finish the pile of 3 bananas.
    - 2 hours for the pile of 6 bananas.
    - 2 hours for the pile of 7 bananas.
    - 3 hours for the pile of 11 bananas.
      Total = 8 hours.

# Approach
The solution uses a binary search approach to find the minimum eating speed:
1. **Initialize Bounds**: Set `left` to 1 (minimum possible speed) and `right` to the maximum number of bananas in any single pile (this is the maximum speed needed if only one hour is available).
2. **Binary Search**: Adjust the `left` and `right` bounds based on whether it's possible to eat all bananas at the current `mid` speed within `h` hours:
    - Compute the required hours to finish all piles at `mid` speed.
    - If the person can finish eating within `h` hours at `mid` speed, attempt a slower speed (`right = mid - 1`).
    - If not, try a faster speed (`left = mid + 1`).
3. **Check Function**: Calculate the total hours needed to eat all bananas at a given speed and compare it to `h`.

### Steps
1. Determine the maximum number of bananas in any pile to set the upper bound for speed.
2. Use binary search between `left = 1` and `right = max(pile)` to find the minimum viable speed.
3. In each iteration, calculate the total hours required at `mid` speed using a helper function that sums up the hours needed for each pile.
4. Continue adjusting the binary search bounds based on whether the total hours fit within `h`.

## Complexity Analysis
- **Time Complexity**: O(n log m), where `n` is the number of piles, and `m` is the maximum number of bananas in any pile. The binary search takes logarithmic time in terms of the range of possible speeds, and for each speed, we calculate the hours linearly in terms of piles.
- **Space Complexity**: O(1), as we use only a few extra variables for the binary search and calculations.

# Code
```go
import (
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, bananas := range piles {
		right = max(right, bananas)
	}
	for left <= right {
		mid := left + (right-left)/2
		if check(piles, mid, h) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(piles []int, speed, limit int) bool {
	var hours float64
	for _, bananas := range piles {
		hours += math.Ceil(float64(bananas) / float64(speed))
	}
	return int(hours) <= limit
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
