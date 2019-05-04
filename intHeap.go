package intheap

import (
	"container/heap"
)

type container []int
type intHeap struct {
	heap container
}

func (h container) Len() int           { return len(h) }
func (h container) Less(i, j int) bool { return h[i] < h[j] }
func (h container) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *container) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *container) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewIntHeap() *intHeap {
	var c container
	heap.Init(&c)
	h := intHeap{c}
	return &h
}
func (h *intHeap) Push(x int) {
	heap.Push(&h.heap, x)
}
func (h *intHeap) Pop() int {
	return heap.Pop(&h.heap).(int)
}
