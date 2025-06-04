package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	backtrack([]int{}, 0, nums, &ans)
	return ans
}

func backtrack(curr []int, index int, nums []int, ans *[][]int) {
	*ans = append(*ans, append([]int{}, curr...)) // Добавляем копию текущего подмножества

	for i := index; i < len(nums); i++ {
		curr = append(curr, nums[i])    // Добавляем элемент
		backtrack(curr, i+1, nums, ans) // Рекурсивно вызываем backtrack
		curr = curr[:len(curr)-1]       // Удаляем элемент (откат)
	}
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result)
}
