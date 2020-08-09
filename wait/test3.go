package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() { //子协程   //没来的及执行主进程结束
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 2; i++ { //默认先执行主进程主进程执行完毕
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("执行")
	}
}
