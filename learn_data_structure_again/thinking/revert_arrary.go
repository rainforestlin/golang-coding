package thinking

//【题目】给定一个经过任意位数的旋转后的排序数组，判断某个数是否在里面。
//例如，对于一个给定数组 {4, 5, 6, 7, 0, 1, 2}，它是将一个有序数组的前三位旋转地放在了数组末尾。
//假设输入的 target 等于 0，则输出答案是 4，即 0 所在的位置下标是 4。如果输入 3，则返回 -1。

// step 1. 先做复杂度分析
//这个问题就是判断某个数字是否在数组中，因此，复杂度极限就是全部遍历地去查找，也就是 O(n) 的复杂度。
// step 2. 定位问题
//判断某个数是否在数组里面，这就是一个查找问题。
// step 3. 数据操作分析
//原数组是经过某些处理的排序数组，也就是说原数组是有序的。有序和查找，你就会很快地想到，这个问题极有可能用二分查找的方式去解决，时间复杂度是 O(logn)，相比上面 O(n) 的基线也是有显著的提高。

func getIndex(arr []int, target, begin, end int) int {
	if begin == end {
		if target == arr[begin] {
			return begin
		} else {
			return -1
		}
	}

	middle := (begin + end) / 2
	if target == arr[middle] {
		return middle
	}

	if arr[begin] <= arr[middle-1] {
		if arr[begin] <= target && arr[middle-1] >= target {
			return getIndex(arr, target, begin, middle-1)
		} else {
			return getIndex(arr, target, middle+1, end)
		}
	} else {
		if arr[middle+1] <= target && target <= arr[end] {
			return getIndex(arr, target, middle+1, end)
		} else {
			return getIndex(arr, target, begin, middle-1)
		}
	}
}
