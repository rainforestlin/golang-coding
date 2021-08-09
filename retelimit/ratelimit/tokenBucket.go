package ratelimit

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶
// 一个固定的桶，桶里存放着令牌（token）。一开始桶是空的，
// 系统按固定的时间（rate）往桶里添加令牌，直到桶里的令牌数满，
// 多余的请求会被丢弃。当请求来的时候，从桶里移除一个令牌，如果桶是空的则拒绝请求或者阻塞
type TokenBucket struct {
	rate         int64 // 投放速率 个/每秒
	capacity     int64 // 可装令牌数
	tokens       int64 // 令牌数
	lastTokenSec int64
	lock         sync.Mutex
}

func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now().Unix()
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate // 添加令牌
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now
	if l.tokens > 0 {
		l.tokens--
		return true
	} else {
		return false
	}
}

func (l *TokenBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.tokens = 0
	l.lastTokenSec = time.Now().Unix()
}
