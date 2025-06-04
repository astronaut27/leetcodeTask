package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
}

func numRescueBoats(people []int, limit int) int {
	var ans, j int
	sort.Ints(people)
	for i := len(people) - 1; i >= j; i-- {
		if (people[i] + people[j]) <= limit {
			j++
		}
		ans++
	}
	return ans
}
