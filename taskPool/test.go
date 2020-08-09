package main

import (
	"fmt"
	"sync"
	"time"
)



var sema = NewSemaphore(3)


func main()  {
	userCount := 10
	for i := 0; i < userCount; i++ {
		go Read(i)
	}

	sema.Wait()
}

func Read(i int) {
	defer sema.Done()
	sema.Add(1)

	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
}

//semq

type Semaphore struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewSemaphore(maxSize int) *Semaphore {
	return &Semaphore{
		c:  make(chan struct{}, maxSize),
		wg: new(sync.WaitGroup),
	}
}

func (s *Semaphore) Add(delta int) {
	s.wg.Add(delta)
	for i := 0; i < delta; i++ {
		s.c <- struct{}{}
	}
}

func (s *Semaphore) Done() {
	<-s.c
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}