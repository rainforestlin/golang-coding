package main

import (
	"fmt"

	"github.com/ahmetb/go-linq"
)

type score struct {
	Score int
	Index int
}

func main() {
	//score := [7]int{}
	//count := 42
	//remain := count % 6
	//for i := 1; i < len(score); i++ {
	//	if remain > 0 {
	//		score[i] = i * (count/6 + 1)
	//		remain--
	//		continue
	//	}
	//	score[i] = score[i-1] + count/6
	//	fmt.Println(score[i])
	//
	//}
	//fmt.Println(score)
	x := []score{
		{1, 0}, {2, 0}, {3, 0}, {2, 0}, {5, 0}, {7, 0}, {10, 0}, {7, 0}, {6, 0}, {5, 0}, {3, 0}, {2, 0}, {1, 0},
	}
	fmt.Println(len(x))
	var scoreList []score
	linq.From(x).OrderByDescending(func(i interface{}) interface{} {
		return i.(score).Score
	}).ToSlice(&scoreList)
	scoreMap := make(map[int]interface{})
	index := 1
	for i, val := range scoreList {
		if _, ok := scoreMap[val.Score]; !ok {
			scoreMap[val.Score] = index
			scoreList[i].Index = index
		} else {
			scoreList[i].Index = scoreMap[val.Score].(int)
		}
		index++
	}
	fmt.Println(scoreList)
}

//func paginate(x []int, skip int, size int) []int {
//	limit := func() int {
//		if skip+size > len(x) {
//			return len(x)
//		} else {
//			return skip + size
//		}
//
//	}
//
//	start := func() int {
//		if skip > len(x) {
//			return len(x)
//		} else {
//			return skip
//		}
//	}
//	return x[start():limit()]
//}
