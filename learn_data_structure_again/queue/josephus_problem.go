package main

import (
	"errors"
	"fmt"
)

type Man struct {
	no   int
	next *Man
}

type Queue struct {
	length int
	start  *Man
	end    *Man
}

func (q *Queue) Append(newMan *Man) {
	if q.length == 0 {
		q.start = newMan
		q.end = newMan
	} else {
		lastMan := q.end
		lastMan.next = newMan
		q.end = newMan
	}
	q.length++
}

func (q *Queue) Delete() *Man {
	if q.length == 0 {
		panic(errors.New("empty"))
	}
	currentStart := q.start
	q.start = currentStart.next
	q.length--
	return currentStart
}

func josephusRing(n, m, k int) {
	if k > n || m > n {
		fmt.Errorf("the total number is not enough")
	}
	queue := &Queue{}
	for i := 1; i <= n; i++ {
		queue.Append(&Man{
			no: i,
		})
	}
	//从第K个开始计数
	for i := 1; i < k; i++ {
		currElement := queue.Delete()
		queue.Append(currElement)
	}
	no := 1
	for queue.length > 0 {
		man := queue.Delete()
		if no < m {
			man.next = nil
			queue.Append(man)
			no++
		} else {
			no = 1
			fmt.Println(man.no)
		}
	}
}

func main() {
	josephusRing(10, 1, 1)
}
