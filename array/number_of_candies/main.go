package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2,3,5,1,3},3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxi, result := 0, make([]bool, len(candies))

    // Find the maximum number of candies
    for _, v := range candies {
        maxi = max(maxi, v)
    }

    // Check if each child can reach or exceed the maximum
    for i, v := range candies {
        if maxi <= v + extraCandies {
            result[i] = true
        } else {
            result[i] = false
        }
    }
    return result
}
