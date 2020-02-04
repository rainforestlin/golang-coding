package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	//终端信号。
	interrupt chan os.Signal
	//处理任务完成通道
	complete  chan error
	//时间间隔
	timeout   <-chan time.Time
	//函数切片
	tasks     []func(int)
}
//超时提醒
var ErrTimeout = errors.New("received timeout")
//终端提醒
var ErrInterrupt = errors.New("received interruption")
//返回一个新的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal,1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}
//新增任务进入Runner
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err:= <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout

	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
