package main

import "fmt"

func main()  {
	arr := []int{2,1,6,12,3,4,9,0,5,70,54}
	//BubbleSort(arr)
	//SelectionSort(arr)
	//InsertionSort(arr)
	MergeSort(arr)
	//QuickSort(arr)
	fmt.Print(arr)
}