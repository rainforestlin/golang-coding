package main

import "fmt"

func mergeSort(nums []int) {
	merge(nums, 0, len(nums)-1)
}

func merge(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	var temp []int
	merge(nums, start, mid)
	merge(nums, mid+1, end)
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		temp = append(temp, nums[i])

	}
	for ; j <= end; j++ {
		temp = append(temp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = temp[i-start]
	}
}

func main() {
	s := []int{3, 2, 1, 6, 5, 4, 9, 8, 7}
	mergeSort(s)
	fmt.Println(s)
}
