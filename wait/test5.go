package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go f(i, &wg)
	}
	wg.Wait()
}

// 一定要通过指针传值，不然进程会进入死锁状态
func f(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}
