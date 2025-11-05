package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 1, 1, 3, 3, 3, 1, 1, 1, 1, 0}))
}

func uniqueOccurrences(arr []int) bool {
    store:=make(map[int]int)
    seen := make(map[int]struct{})
    for _, val := range arr {
            store[val]++
    }
	for _, val := range store {
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
            continue
		}
        return false
	}
    return true
}