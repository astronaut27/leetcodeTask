package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	n := len(nums)
	prefixSum := make([]int, n+1)

	// Вычисляем префиксную сумму
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	maxFreq := 1
	left := 0

	for right := 1; right < n; right++ {
		// Общее количество операций, необходимых для приведения nums[left..right-1] к nums[right]
		operationsNeeded := nums[right]*(right-left) - (prefixSum[right] - prefixSum[left])

		// Пока количество операций превышает допустимое или изменение превышает k, сдвигаем левый указатель
		for operationsNeeded > k*numOperations || nums[right]-nums[left] > k || (right-left) > numOperations {
			left++
			operationsNeeded = nums[right]*(right-left) - (prefixSum[right] - prefixSum[left])
		}

		// Обновляем максимальную частоту
		currentFreq := right - left + 1
		if currentFreq > maxFreq {
			maxFreq = currentFreq
		}
	}

	// Максимальная частота не может превышать количество операций плюс количество неизмененных элементов
	if maxFreq > numOperations+1 {
		maxFreq = numOperations + 1
	}

	return maxFreq
}

func main() {
	nums1 := []int{1, 4, 5}
	k1 := 1
	numOperations1 := 2
	fmt.Println(maxFrequency(nums1, k1, numOperations1)) // Output: 2

	nums2 := []int{5, 11, 20, 20}
	k2 := 5
	numOperations2 := 1
	fmt.Println(maxFrequency(nums2, k2, numOperations2)) // Output: 2

	nums3 := []int{88, 53}
	k3 := 27
	numOperations3 := 2
	fmt.Println(maxFrequency(nums3, k3, numOperations3)) // Output: 2

	nums4 := []int{18, 57}
	k4 := 97
	numOperations4 := 2
	fmt.Println(maxFrequency(nums4, k4, numOperations4)) // Output: 2

	nums5 := []int{2, 70, 73}
	k5 := 39
	numOperations5 := 2
	fmt.Println(maxFrequency(nums5, k5, numOperations5)) // Output: 2
}
