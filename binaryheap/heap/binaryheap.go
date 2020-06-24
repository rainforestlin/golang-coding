package binaryHeap

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Element []int
}

//定义构造方法
func NewMinHeap() *MinHeap {
	return &MinHeap{Element: []int{math.MinInt64}}
}

func (H *MinHeap) Insert(value int) {
	H.Element = append(H.Element, value)
	i := len(H.Element) - 1
	// 上浮
	for ; H.Element[i/2] > value; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}
	H.Element[i] = value
}

func (H *MinHeap) Delete() (int, error) {
	if len(H.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}
	// 0为空，便于计算父子节点
	minElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i < len(H.Element)/2; i = child {
		child := i * 2
		// 找到子节点的最小值
		if child < len(H.Element)-1 && H.Element[child+1] < H.Element[child] {
			child++
		}
		// 将最后一个节点进行下沉操作
		if lastElement > H.Element[child] {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return minElement, nil
}

func (H *MinHeap) Min() (int, error) {
	if len(H.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}
	return H.Element[1], nil
}

func (H *MinHeap) String() string {
	return fmt.Sprintf("%v", H.Element[1:])
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MaxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
