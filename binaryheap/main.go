package main

import (
	binaryHeap "binaryheap/heap"
	"container/heap"
	"fmt"
)

func main() {
	maxHeap := make(binaryHeap.MaxHeap, 0)
	heap.Init(&maxHeap)
	unsorted := []int{1, 2, 3, 4, 5, 6,3}
	for _, i := range unsorted {
		heap.Push(&maxHeap, i)
	}
	fmt.Println(heap.Pop(&maxHeap))
}
