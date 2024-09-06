package longest_subseq_limited_sum

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}

func answerQueries(nums []int, queries []int) []int {
	var result []int
	sort.Ints(nums)
	prefixSum := calcSum(nums)
	for i := 0; i < len(queries); i++ {
		result = append(result, searchInsert(prefixSum, queries[i]))
	}
	return result
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return left - 1
}

func calcSum(nums []int) []int {
	prefixSum := make([]int, len(nums)+1)
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		prefixSum[i+1] = tmp
	}
	return prefixSum
}
