package main

import (
	"fmt"
	"time"
)

func CollatzOrbits(n int64, ch chan bool) {
	if n == 1 {
		ch <- true
		return
	}
	if n%2 == 0 {
		CollatzOrbits(n/2, ch)
	} else {
		CollatzOrbits(n*3+1, ch)
	}
}

func main() {
	ch := make(chan bool)
	i:= int64(1)
cont :for  {

		go func() {
			CollatzOrbits(i, ch)
		}()
		select {
		case  <-ch:
			fmt.Println(i)
			i++
			goto cont
		case <-time.After(3 * time.Second):
			fmt.Printf("%d has a time out", i)
			goto forEnd
		default:
			fmt.Printf("%d has an error\n", i)

		}
	}
forEnd :
	fmt.Println("xxxx")

}
