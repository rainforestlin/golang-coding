package main

import (
	"log"
	"sync"
	"time"

	"github.com/julianlee107/learning/retelimit/ratelimit"
)

type RateLimiter interface {
	Allow() bool
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	var rateLimiter RateLimiter

	for i := 0; i < 10; i++ {
		once.Do(func() {
			//rateLimiter = CreateRequest(&ratelimit.TokenBucket{}, int64(3), int64(1))
			rateLimiter = CreateRequest(&ratelimit.LeakyBucket{}, float64(1), float64(3))
		})
		wg.Add(1)
		log.Println("创建请求:", i)
		go func(i int) {
			if rateLimiter.Allow() {
				log.Println("响应请求:", i)
			}
			wg.Done()
		}(i)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Wait()
}

func CreateRequest(i RateLimiter, args ...interface{}) RateLimiter {
	switch i.(type) {
	case *ratelimit.TokenBucket:
		var tokenBucket ratelimit.TokenBucket
		if len(args) == 2 {
			tokenBucket.Set(args[0].(int64), args[1].(int64))
		}
		return &tokenBucket
	case *ratelimit.LeakyBucket:
		var leakyBucket ratelimit.LeakyBucket
		if len(args) == 2 {
			leakyBucket.Set(args[0].(float64), args[1].(float64))
		}
		return &leakyBucket
	}
	return nil
}
