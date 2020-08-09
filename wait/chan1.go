package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go func() {
		a := 0
		for {
			fmt.Println(a)
			a++
			time.Sleep(time.Second * 10)
		}
	}()
	<-c
}
