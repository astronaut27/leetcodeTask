# Intuition
The goal is to reverse the order of words in a given string. A word is defined as a sequence of non-space characters. There may be multiple spaces between words, and we must ensure the result contains only single spaces between words and no leading/trailing spaces.

# Approach
We present two approaches to solve this problem:

### 1. **Using Go standard libraries**:
- Use `strings.Split` to split the input string into words.
- Iterate from the end of the resulting slice to build the reversed list of words.
- Use `strings.Join` to join the reversed words with a single space.
- Skip any empty strings during reversal to handle multiple spaces properly.

### 2. **Manual parsing without using strings library**:
- Iterate through the input string character by character.
- Build each word manually by collecting characters until a space is found.
- Skip any extra spaces.
- Reverse the collected words and build the final string using manual string concatenation.

# Complexity
- Time complexity:  
  $$O(n)$$ — where \( n \) is the length of the input string. Each character is processed at most twice.

- Space complexity:  
  $$O(n)$$ — to store the list of words and the final result string.

# Code (Using `strings` package)
```golang []
import "strings"

func reverseWords(s string) string {
    bankWord := strings.Split(s, " ")
    var reversed []string
    for i := len(bankWord)-1; i >= 0; i-- {
        if bankWord[i] == "" {
            continue
        }
        reversed = append(reversed, bankWord[i])
    }
    return strings.Join(reversed, " ")
}
```
# Code (Manual parsing without strings package)

func reverseWords(s string) string {
var words []string
word := ""

    // Manual split
    for i := 0; i < len(s); i++ {
        ch := s[i]
        if ch != ' ' {
            word += string(ch)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }

    // Add the last word if exists
    if word != "" {
        words = append(words, word)
    }

    // Manual join in reverse order
    result := ""
    for i := len(words) - 1; i >= 0; i-- {
        result += words[i]
        if i != 0 {
            result += " "
        }
    }

    return result
}
