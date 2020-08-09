package main

import (
	"fmt"
	"time"
)

func main() {
	pool := New(2)

	for i := 1; i < 5; i++ {
		pool.NewTask(func() {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Now())
		})
		fmt.Println("aa")
	}

	// 保证所有的协程都执行完毕
	time.Sleep(5 * time.Second)
}

type Pool struct {
	work  chan func()   // 任务
	count chan bool // 数量
}

func New(size int) *Pool {
	return &Pool{
		work:  make(chan func()),
		count: make(chan bool, size),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.count <- true:
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() { <-p.count }()
	for {
		task()
		task = <-p.work
	}
}
