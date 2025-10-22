# Intuition
The problem asks whether we can plant `n` flowers in a flowerbed such that no two flowers are adjacent. The initial idea is to iterate through the flowerbed and greedily plant flowers wherever it is allowed — that is, where both neighboring positions (if they exist) are empty.

# Approach
We iterate through each position in the `flowerbed` array and check three conditions before planting a flower:
1. The current plot is empty (`flowerbed[i] == 0`).
2. The left neighbor is either empty or does not exist (`i == 0 || flowerbed[i-1] == 0`).
3. The right neighbor is either empty or does not exist (`i == len(flowerbed)-1 || flowerbed[i+1] == 0`).

If all three conditions are satisfied, we place a flower by setting `flowerbed[i] = 1` and increment a counter. Finally, we check if we have planted at least `n` flowers.

This greedy strategy ensures we plant flowers in valid positions and stop early if possible.

# Complexity
- Time complexity:  
  $$O(n)$$, where \( n \) is the length of the flowerbed. We scan the array once.

- Space complexity:  
  $$O(1)$$, as we use only constant extra space.

# Code
```golang []
func canPlaceFlowers(flowerbed []int, n int) bool {
    var count int 
    for i := range flowerbed {
        if flowerbed[i] != 0 {
            continue
        }
        if i != 0 && flowerbed[i-1] != 0 {
            continue
        }
        if i != len(flowerbed)-1 && flowerbed[i+1] != 0 {
            continue
        }
        flowerbed[i] = 1
        count++
    }
    return count >= n
}
