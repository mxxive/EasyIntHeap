package EasyIntHeap

import (
	"container/heap"
)

type intHeap []int
type IntHeap intHeap

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewIntHeap() *IntHeap {
	var h intHeap
	heap.Init(&h)
	h2 := IntHeap(h)
	return &h2
}
func (h *IntHeap) Push(x int) {
	h.Push(x)
}
func (h *IntHeap) Pop() int {
	return h.Pop()
}
