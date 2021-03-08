package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner在给定的超时时间内执行一组任务，
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt通道报告从操作系统
	//发送信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout报告处理任务已经超时
	timout <-chan time.Time

	// tasks持有一组以索引顺序依次执行的
	//函数
	tasks []func(int)
}

//超时
var ErrTimeout = errors.New("timeout-超时了~~")

//中断
var ErrInterrupt = errors.New("interrupt-任务中断~~")

func main() {
}

//获取一个 新的runner 准备使用
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timout:    time.After(d),
		tasks:     nil,
	}
}

// Add将一个任务附加到Runner上。这个任务是一个
// 接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//执行所有任务 并监视通道
func (r *Runner) Start() error {

	//接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		Runner{}
		r.complete <- r.run()
	}()

	select {
	//任务完成时候发出的信号
	case err := <-r.complete:
		return err
	//任务处理超时发出的信号
	case <-r.timout:
		return ErrTimeout
	}
}

//执行每一个注册的任务
func (r *Runner) Run() error {
	for id, task := range r.tasks {
		//检测是否有中断信号
		if r.goIntrrupt() {
			return ErrInterrupt
		}
		//执行已经注册的任务
		task(id)
	}
	return nil
}

//检测是否收到了中断信号
func (r *Runner) goIntrrupt() bool {
	select {
	//当中断事件被触发 发出信号
	case <-r.interrupt:
		//停止接收后续任何信号
		signal.Stop(r.interrupt)
		return true

	//继续正常运行
	default:
		return false
	}
}
