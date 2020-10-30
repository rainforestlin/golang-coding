package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}

	newSlice := slice[1:3]
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	fmt.Println(slice)
	fmt.Println(newSlice)

	newSlice = append(newSlice, 40)

	fmt.Println(slice)
	fmt.Println(newSlice)
	// 当容量满足时，会修改共享的底层数组，当容量超过时，append会新建一个新的底层数组
	for i := 0; i < 10; i++ {
		newSlice = append(newSlice, i*i)
	}

	fmt.Println(slice)
	fmt.Println(newSlice)

	c := slice[1:2:7]
	fmt.Println(c)
}
