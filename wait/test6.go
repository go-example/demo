package main

import (
	"fmt"
	"sync"
)

// 等待所有 goroutine 执行完毕
// 使用传址方式为 WaitGroup 变量传参
// 使用 channel 关闭 goroutine

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{})

	workerCount := 5
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, ch, done, &wg) // wg 传指针，doIt() 内部会改变 wg 的值
	}

	for i := 0; i < workerCount; i++ { // 向 ch 中发送数据，关闭 goroutine
		ch <- i
	}

	close(done)
	wg.Wait()
	close(ch)
	fmt.Println("all done!")
}

func doIt(workerID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	for {
		select {
		case m := <-ch:
			fmt.Printf("[%v] m => %v\n", workerID, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerID)
			return
		}
	}
}
