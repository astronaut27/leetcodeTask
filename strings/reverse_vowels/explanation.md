# Intuition
The task is to reverse only the vowels in a given string. All other characters must remain in their original positions. The first idea that comes to mind is to use two pointers — one starting from the beginning and the other from the end — and swap vowels when both pointers point to them.

# Approach
We use the two-pointer technique:
- Initialize two indices `l` and `r` at the start and end of the string.
- Move `l` forward until it points to a vowel.
- Move `r` backward until it points to a vowel.
- When both `l` and `r` are at vowels, swap them.
- Repeat this process while `l < r`.

To check for vowels efficiently, we use a map `vowels` with all lowercase and uppercase vowels as keys.

We convert the string to a mutable `[]byte` slice so we can perform swaps directly, then convert it back to a string at the end.

# Complexity
- Time complexity:  
  $$O(n)$$ — where \( n \) is the length of the string. Each character is visited at most once by either pointer.

- Space complexity:  
  $$O(n)$$ — to store the result as a byte slice (in Go, strings are immutable).

# Code
```golang []
var vowels = map[byte]struct{}{
    'a':{}, 'e':{}, 'i':{}, 'o':{}, 'u':{},
    'A':{}, 'E':{}, 'I':{}, 'O':{}, 'U':{},
}

func reverseVowels(s string) string {
    var l, r int
    r = len(s) - 1
    result := []byte(s)
    
    for l < r {
        if _, ok := vowels[result[l]]; !ok {
            l++
            continue
        }
        if _, ok := vowels[result[r]]; !ok {
            r--
            continue
        }
        result[l], result[r] = result[r], result[l]
        l++
        r--
    }
    
    return string(result)
}