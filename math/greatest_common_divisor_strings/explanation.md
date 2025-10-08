# Intuition
To find the greatest common divisor (GCD) of two strings, the idea is that if there is a common base string that can construct both strings through repetition, then concatenating them in both orders should yield the same result. If not, there's no such common divisor string.

# Approach
1. Check if `str1 + str2` is equal to `str2 + str1`. If not, return an empty string because there's no common base string that can form both.
2. If the condition is satisfied, compute the GCD of the lengths of the two strings.
3. Return the prefix of either string from index `0` to the length of the GCD. This substring will be the largest string that can divide both input strings.

# Complexity
- Time complexity:  
  $$O(n + m)$$  
  where \( n \) and \( m \) are the lengths of `str1` and `str2` respectively. This comes from string concatenation and comparison.

- Space complexity:  
  $$O(n + m)$$  
  due to the temporary strings created during concatenation and slicing.

# Code
```golang
func gcdOfStrings(str1 string, str2 string) string {
    if str1+str2 != str2+str1 {
        return ""
    }
    gcd := gcd(len(str1), len(str2))
    if len(str1) > len(str2) {
        return str1[0:gcd]
    }
    return str2[0:gcd]
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
