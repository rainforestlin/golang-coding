package main

//冒泡排序
//最差时间复杂度O(n^2)，最好时间复杂度O(n)，平均时间复杂度O(n^2)
//空间复杂度O(1)
//是一种稳定算法
func BubbleSort(arr []int) {
	swapped := true
	arrLen := len(arr)
	x := 0
	for swapped {
		x = x + 1
		swapped = false
		for i := 0; i < arrLen-1; i++ {
			//fmt.Printf("第%d轮-第%d次：%v", x, i, arr)
			//fmt.Println()
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
}

//选择排序
//最好时间复杂度O(n^2)，最差时间复杂度O(n^2)，平均时间复杂度O(n^2)
//空间复杂度O(1)
//是一种不稳定的排序算法
func SelectionSort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		minIndex := i
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
		//fmt.Printf("第%d次：%v", i, arr)
		//fmt.Println()
	}
}

//插入算法
//最好时间复杂度O(n)，最差时间复杂度O(n^2)，平均时间复杂度O(n^2)
//空间复杂度为O(1)
//插入算法是一种稳定的排序算法
func InsertionSort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		selected := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > selected {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				arr[j+1] = selected
				break
			}
		}
		//fmt.Printf("第%d次：%v", i, arr)
		//fmt.Println()
	}
}

//归并排序
//最好时间复杂度O(nlogn)，最差时间复杂度O(nlogn)，平均时间复杂度O(nlogn)
//空间复杂度为O(n)
//归并算法是一种稳定的排序算法
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	return merge(left, right)
}
func merge(a, b []int) []int {
	aLen, bLen := len(a), len(b)
	result := make([]int, aLen+bLen)
	index := 0
	i, j := 0, 0
	for i < aLen && j < bLen {
		if a[i] < b[j] {
			result[index] = a[i]
			i++
		} else {
			result[index] = b[j]
			j++
		}
		index++
	}
	for i != aLen {
		result[index] = a[i]
		index++
		i++
	}
	for j != bLen {
		result[index] = b[j]
		index++
		j++
	}
	return result
}

//快速排序
//最好时间复杂度O(nlogn)，最差时间复杂度O(n^2)，平均时间复杂度O(nlogn)
//空间复杂度为O(logn)~O(n)
//快速排序是一种不稳定的排序算法
func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}
func sort(arr []int, left, right int) {
	if right <= left {
		return
	}
	key := arr[(left+right)/2]
	i := left
	j := right
	for i <= j {
		for arr[i] < key {
			i++
		}
		for arr[j] > key {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
		if left < j {
			sort(arr, left, j)
		}
		if right > i {
			sort(arr, i, right)
		}
	}

}
