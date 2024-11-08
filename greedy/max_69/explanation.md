# Problem Statement
Given a positive integer `num` consisting only of digits 6 and 9, return the maximum number you can get by changing at most one digit (6 becomes 9, or 9 becomes 6).

## Example
- **Input**: num = 9669
- **Output**: 9969
- **Explanation**: Change the first 6 to a 9.

# Approach
The solution scans the digits of `num` and modifies the first occurrence of the digit '6' to '9', which yields the maximum possible number:
1. **Convert Integer to String**: Convert the number `num` to its string representation for easy digit access.
2. **Iterate Over Digits**: Loop through each digit in the string representation of `num`.
    - If a '6' is found, calculate the value to add by computing the difference between '9' and '6' adjusted for the digit's position.
    - Add this value to `num` to get the maximum number.
3. **Return Modified Number**: If no '6' is found, return `num` as is. If a '6' is encountered, return the updated number.

### Helper Function - `pow()`
- **Purpose**: Computes the power of a number. Used to find the positional multiplier for the digit to be changed.
- **Implementation**: Multiplies the base `num` by itself `n` times to compute `num` raised to the power of `n`.

## Complexity Analysis
- **Time Complexity**: O(d), where `d` is the number of digits in `num`. This is due to the single pass needed to inspect each digit.
- **Space Complexity**: O(d), as the integer `num` is converted into a string for digit manipulation.

# Code
```go
package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
    str := strconv.Itoa(num)
    for i := 0; i < len(str); i++ {
        if str[i] == '6' {
            return num + 3 * pow(10, len(str
