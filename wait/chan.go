package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	//只有写数据后才能继续执行
	done <- false

	time.Sleep(time.Second)
	fmt.Println("hello go routine awake and going to write to done")

	done <- true
}
func main() {
	done := make(chan bool, 3)
	go hello(done)
	<-done
	//fmt.Println("Main received data")
}
