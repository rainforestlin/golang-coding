package main

import "fmt"

// 一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级... 它也可以跳上 n 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
func jumpFloorII(number int) int {
	// write code here
	arr1, arr2 := 1, 2
	if number == 1 {
		return arr1
	}
	for i := 0; i < number-2; i++ {
		arr2 = arr1 + arr2 + 1
		arr1 = arr2 - 1
	}
	return arr2
}

func jumpFloorIIAdvanced(num int) int {
	if num < 2 {
		return 1
	}
	return 2<<num - 1
}

func main() {
	fmt.Println(jumpFloorII(10))
}
