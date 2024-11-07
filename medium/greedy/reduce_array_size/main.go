package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))

}

func minSetSize(arr []int) int {
	storeMap := make(map[int]int, len(arr))
	finalStore := make([]int, 0, len(arr))
	for _, v := range arr {
		if _, ok := storeMap[v]; ok {
			storeMap[v]++
			continue
		}
		storeMap[v] = 1
	}
	for _, v := range storeMap {
		finalStore = append(finalStore, v)
	}
	sort.Sort(ByDesc(finalStore))
	var result int
	limit := len(arr) / 2
	for _, count := range finalStore {
		limit -= count
		result++
		if limit <= 0 {
			return result
		}
	}
	return result
}

type ByDesc []int

func (a ByDesc) Len() int {
	return len(a)
}

func (a ByDesc) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a ByDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
