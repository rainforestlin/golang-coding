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
	ch := make(chan bool,1)
	i := int64(1)
	ch <- false
	mutex = new(sync.Mutex)
cont:
	for {
		go func() {
			max = 1
			CollatzOrbits(i, ch)
		}()
		select {
		case ok:= <-ch:
			if ok{
				fmt.Printf("num:%d max:%d \n", i, max)
				i++
			goto cont
			}else{
				fmt.Printf("%d has an error\n", i)
			fmt.Println(<-ch)
			}
			
		case <-time.After(3 * time.Second):
			fmt.Printf("%d has a time out", i)
			goto forEnd
		}
forEnd:
	fmt.Println("xxxx")

}
}