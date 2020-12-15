package main

import (
	"fmt"
	"sync"
	"time"
)

var max int64
var mutex *sync.Mutex

func CollatzOrbits(n int64, ch chan bool) {
	if n == 1 {
		ch <- true
		return
	}
	if n >= max {
		mutex.Lock()
		max = n
		mutex.Unlock()
	}

	if n%2 == 0 {
		CollatzOrbits(n/2, ch)
	} else {
		CollatzOrbits(n*3+1, ch)
	}
}

func main() {
	ch := make(chan bool)
	i := int64(1)

	mutex = new(sync.Mutex)
cont:
	for {

		go func() {
			max = 1
			CollatzOrbits(i, ch)
		}()
		select {
		case <-ch:
			fmt.Printf("num:%d max:%d \n", i, max)
			i++
			goto cont
		case <-time.After(3 * time.Second):
			fmt.Printf("%d has a time out", i)
			goto forEnd
		default:
			fmt.Printf("%d has an error\n", i)

		}
	}
forEnd:
	fmt.Println("xxxx")

}
