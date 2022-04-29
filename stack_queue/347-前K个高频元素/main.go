package main

import (
	"container/heap"
)

type TopHeap [][2]int

func (t TopHeap) Len() int {
	return len(t)
}

func (t TopHeap) Less(i, j int) bool {
	return t[i][1] < t[j][1]
}

func (t TopHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *TopHeap) Push(x interface{}) {
	*t = append(*t, x.([2]int))
}

func (t *TopHeap) Pop() interface{} {
	old := *t
	v := old[len(old)-1]
	*t = old[:len(old)-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	record := map[int]int{}
	for _, v := range nums {
		record[v]++
	}
	h := &TopHeap{}
	heap.Init(h)

	for key, value := range record {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).([2]int)[0])
	}
	return result
}

func main() {

}
