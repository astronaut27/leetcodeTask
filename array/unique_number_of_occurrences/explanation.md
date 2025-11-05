# Intuition
To determine if each element in the array appears a unique number of times, we need a way to count occurrences and check for duplicates in those counts. A map is a natural choice for counting, and a set (represented as another map in Go) helps track uniqueness.

# Approach
1. **Sort the input array** so that identical elements are grouped together.
2. Traverse the array and count the frequency of each unique number using a `store` map.
3. After recording the frequencies, loop through the map values and check if each frequency is unique by using a `seen` map.
4. If any frequency is repeated, return `false`. Otherwise, return `true`.

# Complexity
- Time complexity:  
  $$O(n \log n)$$ — due to sorting the array. Counting and checking uniqueness are both linear time operations.

- Space complexity:  
  $$O(n)$$ — in the worst case, each element is unique, requiring space for counting and seen maps.

# Code
```golang []
func uniqueOccurrences(arr []int) bool {
    sort.Ints(arr)
    store := make(map[int]int)
    seen := make(map[int]bool)
    last, count := arr[0], 1

    for i := 1; i < len(arr); i++ {
        if last == arr[i] {
            count++
            continue
        }
        store[last] = count
        last, count = arr[i], 1
    }
    store[last] = count

    for _, val := range store {
        if _, ok := seen[val]; !ok {
            seen[val] = true
            continue
        }
        return false
    }

    return true
}
