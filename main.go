package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("123")
		ch <- 1
	}()
	<-ch
	fmt.Println("3455")
}
