package main

import (
	"fmt"
	"runtime"
	"time"
)

// 为什么会输出9 0-8
//go在把goroutine放入到队列中时，go会把每个processor所管理的最后一个goroutine放入到next的位置
//go的设计认为：如果一个P的goroutine队列在顺序执行的时候，因为go sched会有很多抢占或者调度，那么从被执行的概率上来分析的话，放入一个next位置可使得每个goroutine的执行效率相当
func main() {
	runtime.GOMAXPROCS(1)
	LoopOne()
	runtime.Gosched()
	time.Sleep(time.Second)
}

func LoopOne() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
}

func LoopTwo() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
