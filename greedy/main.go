package main

import (
	"fmt"
)

// maxRemoval calculates the maximum number of queries that can be excluded
func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)
	m := len(queries)

	// Helper function to check if nums can be converted to a zero array
	canConvertToZero := func(activeQueries []int) bool {
		effects := make([]int, n+1)

		// Apply the active queries
		for _, idx := range activeQueries {
			li, ri := queries[idx][0], queries[idx][1]
			effects[li]--
			effects[ri+1]++
		}

		// Calculate the cumulative effect and validate nums
		currEffect := 0
		for i := 0; i < n; i++ {
			currEffect += effects[i]
			if nums[i]+currEffect < 0 {
				return false
			}
		}
		return true
	}

	// Initial active queries (all queries are active)
	activeQueries := make([]int, m)
	for i := range activeQueries {
		activeQueries[i] = i
	}

	removedCount := 0

	// Try removing each query one by one
	for i := 0; i < m; i++ {
		// Temporarily remove the i-th query
		tempActiveQueries := make([]int, 0, len(activeQueries))
		for _, idx := range activeQueries {
			if idx != i {
				tempActiveQueries = append(tempActiveQueries, idx)
			}
		}

		// Check if nums can still be converted to a zero array
		if canConvertToZero(tempActiveQueries) {
			// If possible, permanently exclude the i-th query
			activeQueries = tempActiveQueries
			removedCount++
		}
	}

	// Final check if nums can still be converted to a zero array
	if !canConvertToZero(activeQueries) {
		return -1
	}

	return removedCount
}

func main() {
	// Example 1
	nums1 := []int{2, 0, 2}
	queries1 := [][]int{{0, 2}, {0, 2}, {1, 1}}
	fmt.Println(maxRemoval(nums1, queries1)) // Output: 1

	// Example 2
	nums2 := []int{1, 1, 1, 1}
	queries2 := [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}
	fmt.Println(maxRemoval(nums2, queries2)) // Output: 2

	// Example 3
	nums3 := []int{1, 2, 3, 4}
	queries3 := [][]int{{0, 3}}
	fmt.Println(maxRemoval(nums3, queries3)) // Output: -1
}
