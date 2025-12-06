# Intuition
The goal is to move all the zeroes to the end of the list while maintaining the relative order of the non-zero elements.

My first thought was to use two pointers:
- One pointer (`i`) keeps track of the position to place the next non-zero value.
- The other pointer (`j`) iterates through the array.

# Approach
We iterate over the list with a for-loop using pointer `j`.  
If we find a non-zero element at index `j` and the element at `i` is zero, we **swap** the elements at positions `i` and `j`.  
Then we increment `i` to move the pointer forward for the next potential zero.

This approach ensures:
- Non-zero elements are shifted to the front.
- Zeroes are pushed to the end.
- The operation is done in-place (no extra array).

# Complexity
- Time complexity:  
  $$O(n)$$ — We traverse the list only once.

- Space complexity:  
  $$O(1)$$ — In-place operation, no additional memory used.

# Code
```python3 []
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i = 0
        for j in range(len(nums)):
            if (i == j) or (nums[i] == 0 and nums[j] == 0):
                continue
            if nums[i] == 0 and nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]
            i += 1
```