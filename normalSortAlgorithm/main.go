package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := &IntHeap{100, 23, 45, 76, 2, 52, 6674, 2342, 34, 56, 23}
	heap.Init(h)
	fmt.Println(h.Sort())
}
