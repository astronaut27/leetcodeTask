package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

}

func topKFrequent(nums []int, k int) []int {
	// add to map
	store := make(map[int]int, len(nums))
	for _, v := range nums {
		store[v]++
	}
	h := &StructHeap{}
	heap.Init(h)
	for key, v := range store {
		heap.Push(h, Pair{
			val:      key,
			frequent: v,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(Pair).val)
	}
	return ans
}

type StructHeap []Pair

type Pair struct {
	val      int
	frequent int
}

func (h StructHeap) Len() int { return len(h) }

func (h StructHeap) Less(i, j int) bool { return h[i].frequent < h[j].frequent }

func (h StructHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StructHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *StructHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
