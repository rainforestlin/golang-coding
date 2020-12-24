package main

import (
	"fmt"
	"sync"
)

// var m *sync.RWMutex
// var wg sync.WaitGroup

// func main() {
// 	m = new(sync.RWMutex)
// 	for i := 0; i < 4; i++ {
// 		wg.Add(1)
// 		go read(i)
// 	}
// 	wg.Wait()
// }

// func read(i int) {

// 	println(i, "read start")

// 	m.RLock()
// 	println(i, "reading")
// 	time.Sleep(time.Second*1)
// 	defer m.RUnlock()

// 	println(i, "read over")
// 	defer wg.Done()
// }

/**
其原因在于，即使所有的 goroutine 都创建完了，但 goroutine 不一定已经开始运行了。
也就是等到 goroutine 真正去执行输出时，变量 i （值拷贝）可能已经不是创建时的值了。
其整个程序扭转实质上分为了多个阶段，也就是各自运行的时间线并不同，可以其拆分为：
先创建：for-loop 循环创建 goroutine。
再调度：协程goroutine 开始调度执行。
才执行：开始执行 goroutine 内的输出。
同时 goroutine 的调度存在一定的随机性（建议了解一下 GMP 模型），那么其输出的结果就势必是无序且不稳定的。
*/
func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	num := 0
	for i := 0; i < 5; i++ {
		num ++ 
	//  go func(num int) {
	//   defer wg.Done()
	//   fmt.Printf("goroutine %d,print %d \n",num,i)
	//  }(num)
	// }	 
	go func() {
	  defer wg.Done()
	  fmt.Printf("goroutine %d,print %d \n",num,i)
	 }()
	}
	wg.Wait()
   }


/**
其本质还是创建 goroutine 与真正执行 fmt.Println 并不同步。因此很有可能在你执行 fmt.Println 时，循环 for-loop 已经运行完毕，因此变量 i 的值最终变成了 5。
那么相反，其也有可能没运行完，存在随机性。
*/
// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(5)
// 	num := 0
// 	for i := 0; i < 5; i++ {
// 	 num ++
// 	 go func(i int,num int) {
// 	  defer wg.Done()
// 	  fmt.Printf("goroutine %d,print %d \n",num,i)
	   
// 	 }(i,num)
// 	}
// 	wg.Wait()
//    }