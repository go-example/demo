package main

import (
	"fmt"
	"time"
)

func main() {
	input := make(chan interface{})
	//producer - produce the messages
	go func() {
		for i := 0; i < 5; i++ {
			input <- i
		}
		input <- "hello, world"
	}()

	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {
		//consumer - consume the messages
		case msg := <-input:
			fmt.Println(msg)

		case <-t1.C:
			println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			println("10s timer")
			t2.Reset(time.Second * 10)
		}
	}
}

