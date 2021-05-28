package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	x := a
	for i, v := range x {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		fmt.Println(v)
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
