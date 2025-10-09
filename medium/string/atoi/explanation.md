# Intuition
The task is to implement the behavior of the `atoi` function, converting a string to an integer while handling whitespace, optional sign, digits, and integer overflow/underflow. The approach is to simulate manual parsing, which is both straightforward and efficient.

# Approach
1. Trim leading and trailing whitespace from the string.
2. Check if the string is empty after trimming.
3. Determine the sign (`+` or `-`) if present.
4. Iterate through the string, converting valid digit characters into an integer.
5. On each step, check for potential overflow/underflow using `math.MaxInt32` and `math.MinInt32`.
6. Stop processing as soon as a non-digit character is encountered.
7. Return the final result multiplied by the sign.

# Complexity
- Time complexity:  
  $$O(n)$$  
  where \( n \) is the length of the input string, since we may iterate through each character once.

- Space complexity:  
  $$O(1)$$  
  as we use a constant amount of extra space regardless of input size.

# Code
```go
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	sign := 1
	result := 0
	i := 0

	// handle sign
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// Check overflow/underflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
