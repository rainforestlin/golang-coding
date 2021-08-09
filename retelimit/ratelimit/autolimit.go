package ratelimit

import (
	"sync"

	"golang.org/x/time/rate"
)

const (
	RTTDATASET = 100   // rtt数据集
	SAMPLESET  = 90    // 样本数据集
	MODULUS    = 0.8   // srtt经典算法更新系数
	UBOUND     = 800.0 // rto超时上限
	LBOUND     = 400.0 // rto超时下限

	RTOMODULUS = 1.5 // rto变化系数

)

type AutoLimit struct {
	rttset    []float64
	cwnd      rate.Limit
	sshthresh rate.Limit
	srtt      float64
	rto       float64
	mutex     sync.Mutex
	l         *rate.Limiter
}

func NewAutoLimit() *AutoLimit {
	return &AutoLimit{
		rttset:    make([]float64, 0, RTTDATASET),
		cwnd:      1,
		sshthresh: 30000,
		srtt:      0,
		rto:       LBOUND,
		l:         rate.NewLimiter(1, 1),
	}
}

// 收集rtt

func (al *AutoLimit) AddRtt(rtt float64) {
	al.mutex.Lock()
	defer al.mutex.Unlock()
	rto := al.rto
	cwnd := al.cwnd
	sshthresh := al.sshthresh
	if rtt < rto {
		w
	}
}

func (al *AutoLimit)smap  {

}