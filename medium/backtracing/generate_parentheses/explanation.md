# Intuition
Given a matrix where each row is sorted in ascending from left to right and each column is sorted in ascending from top to bottom, we need an efficient method to determine if a target value exists in the matrix. Instead of using a brute-force search, we can leverage the sorted properties of the matrix to skip over large sections of data, akin to a binary search but adapted for a two-dimensional space.

# Approach
The approach is to use a method that reduces the search space with each comparison:
1. **Start at the Top-Right Corner**:
    - The chosen starting point allows us to move left if the current value is greater than the target (since all values to the left are smaller) or move down if the current value is less than the target (since all values below are larger).

2. **Iterative Search**:
    - While staying within the bounds of the matrix (`i` from 0 to `m-1` and `j` from `n-1` to 0), compare the matrix element at the current position with the target.
    - Move left (`j--`) if the current element is greater than the target to find smaller elements.
    - Move down (`i++`) if the current element is less than the target to find larger elements.
    - If you find the target, return `true`.

3. **Termination**:
    - If you exit the loop without having found the target, return `false` as the target does not exist in the matrix.

# Complexity
- **Time Complexity**: O(m + n), where `m` is the number of rows and `n` is the number of columns in the matrix. In the worst case, you might traverse an entire row and an entire column.
- **Space Complexity**: O(1), as no additional space is needed beyond the input matrix and pointers for iteration.


### Explanation
- **Start at the Top-Right**: This starting point is crucial because it gives two clear options: move left to find smaller elements or move down to find larger ones.
- **Matrix Properties**: The method takes advantage of the matrix's properties where each row is sorted in ascending order and each column is sorted from top to bottom in ascending order.
- **Optimized Search**: The search pattern ensures that at each step, the possible location of the target is reduced by one row or one column, thereby quickly narrowing down the potential area where the target can be located.

This method is highly efficient for searching in matrices with this specific sorting property, significantly reducing the need to look at every element.

# Code
```go
package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0]) // Get dimensions of the matrix
    i, j := 0, n-1 // Start at the top-right corner of the matrix
    for i < m && j >= 0 { // Ensure indices are within the bounds
        if matrix[i][j] == target {
            return true // Target found
        }
        if matrix[i][j] > target {
            j-- // Move left
        } else {
            i++ // Move down
        }
    }
    return false // Target not found
}

func main() {
    matrix := [][]int{
        {1, 4, 7, 11, 15},
        {2, 5, 8, 12, 19},
        {3, 6, 9, 16, 22},
        {10, 13, 14, 17, 24},
        {18, 21, 23, 26, 30},
    }
    target := 5
    fmt.Println("Target found:", searchMatrix(matrix, target))
}
```