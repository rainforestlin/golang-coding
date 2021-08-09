package ratelimit

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	rate       float64 // 固定每秒出水速率
	capacity   float64 // 桶容量
	water      float64 // 桶中当前水
	lastLeakMs int64   // 上次漏水时间
	lock       sync.Mutex
}

func (l *LeakyBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().UnixNano() / 1e6
	eclipse := float64(now-l.lastLeakMs) * l.rate / 1000
	l.water = l.water - eclipse
	l.lastLeakMs = now
	if (l.water + 1) < l.capacity {
		l.water++
		return true
	} else {
		return false
	}
}

func (l *LeakyBucket) Set(r, c float64) {
	l.rate = r
	l.capacity = c
	l.water = 0
	l.lastLeakMs = time.Now().UnixNano() / 1e6
}
