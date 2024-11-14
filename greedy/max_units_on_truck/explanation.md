# Problem Statement
Given an array `boxTypes`, where `boxTypes[i] = [numberOfBoxes_i, numberOfUnitsPerBox_i]`, and an integer `truckSize` representing the maximum number of boxes that can be put on the truck, determine the maximum total number of units that can be loaded onto the truck.

## Example
- **Input**: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
- **Output**: 8
- **Explanation**: There are:
    - 1 box of the first type that contains 3 units.
    - 2 boxes of the second type that contains 2 units each.
    - 3 boxes of the third type that contains 1 unit each.
      You can take all the boxes of the first and second types, and one box of the third type. The total number of units will be 1*3 + 2*2 + 1*1 = 8.

# Approach
The solution leverages a greedy algorithm by prioritizing the box types with the highest units per box. The steps include:
1. **Sort `boxTypes`**: Sort the boxes by the number of units per box in descending order. This prioritizes the boxes with the highest units for selection.
2. **Iterate and Simulate Loading**: Traverse the sorted list and simulate loading the boxes onto the truck:
    - Deduct the box count from `truckSize` while adding to `unitCount` until the truck is full or all boxes are considered.
    - If the boxes of a type are more than the remaining `truckSize`, only a fraction of them are taken to fill the truck.

### Steps
1. Sort the `boxTypes` based on the units per box in descending order.
2. Initialize `unitCount` to zero to accumulate the total units loaded.
3. For each box type in `boxTypes`, determine how many can be loaded based on the remaining `truckSize`:
    - If the entire batch of boxes can be loaded, deduct their count from `truckSize` and add the units to `unitCount`.
    - If only part of the batch can be loaded due to space constraints, calculate accordingly and break out of the loop.
4. Return the accumulated `unitCount`.

## Complexity Analysis
- **Time Complexity**: O(n log n), where `n` is the number of box types due to the sorting step. Iterating through the sorted list is O(n).
- **Space Complexity**: O(1), as the algorithm uses constant space excluding the input.

# Code
```go
package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(i, j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    unitCount := 0
    for _, boxType := range boxTypes {
        if boxType[0] < truckSize {
            truckSize -= boxType[0]
            unitCount += boxType[0] * boxType[1]
        } else {
            unitCount += truckSize * boxType[1]
            break
        }
    }
    return unitCount
}

func main() {
    boxTypes := [][]int{{1,3}, {2,2}, {3,1}}
    truckSize := 4
    fmt.Println(maximumUnits(boxTypes, truckSize))  // Output: 8
}
