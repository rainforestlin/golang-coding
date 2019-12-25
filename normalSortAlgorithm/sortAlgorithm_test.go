package main

import "testing"

var sorted = []int{0, 1, 2, 3, 4, 5, 6, 9, 12, 54, 70}


func TestMain(m *testing.M) {
	m.Run()
}

func TestSortWorkFlow(t *testing.T)  {
	t.Run("BubbleSort",testBubbleSort)
	t.Run("InsertSort",testInsertionSort)
	t.Run("MergeSort",testMergeSort)
	t.Run("QuickSort",testQuickSort)
	t.Run("SelectionSort",testSelectionSort)
}

func testBubbleSort(t *testing.T) {
	var unsorted = []int{2,1,6,12,3,4,9,0,5,70,54}
	BubbleSort(unsorted)
	for index,value :=range unsorted{
		if value!=sorted[index]{
			t.Errorf("BubbleSort has a wrong test")
		}
	}
}

func testInsertionSort(t *testing.T) {
	var unsorted = []int{2,1,6,12,3,4,9,0,5,70,54}
	InsertionSort(unsorted)
	for index,value :=range unsorted{
		if value!=sorted[index]{
			t.Errorf("InsertionSort has a wrong test")
		}
	}
}

func testMergeSort(t *testing.T) {
	var unsorted = []int{2,1,6,12,3,4,9,0,5,70,54}
	mergeSort := MergeSort(unsorted)
	for index,value :=range mergeSort{
		if value!=sorted[index]{
			t.Errorf("MergeSort has a wrong test")
		}
	}
}

func testQuickSort(t *testing.T) {
	var unsorted = []int{2,1,6,12,3,4,9,0,5,70,54}
	QuickSort(unsorted)
	for index,value :=range unsorted{
		if value!=sorted[index]{
			t.Errorf("QuickSort has a wrong test")
		}
	}
}

func testSelectionSort(t *testing.T) {
	var unsorted = []int{2,1,6,12,3,4,9,0,5,70,54}
	SelectionSort(unsorted)
	for index,value :=range unsorted{
		if value!=sorted[index]{
			t.Errorf("SelectionSort has a wrong test")
		}
	}
}