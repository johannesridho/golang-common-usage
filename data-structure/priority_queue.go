// priority queue / heap
// https://pkg.go.dev/container/heap

package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(a, b int) bool {
	return h[a] <= h[b]
}
func (h MinHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}
func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *MinHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func usingHeap() {
	fmt.Println("--------------------------------")
	fmt.Println("using heap")

	h := MinHeap{}
	heap.Push(&h, 1)
	heap.Push(&h, 2)
	heap.Push(&h, 3)

	for h.Len() != 0 {
		fmt.Println("pop", heap.Pop(&h))
	}
}
