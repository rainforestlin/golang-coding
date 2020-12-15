package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex
var wg sync.WaitGroup

func main() {
	m = new(sync.RWMutex)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
}

func read(i int) {

	println(i, "read start")

	m.RLock()
	println(i, "reading")
	time.Sleep(time.Second*1)
	defer m.RUnlock()

	println(i, "read over")
	defer wg.Done()
}
