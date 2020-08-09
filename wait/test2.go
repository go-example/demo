package main

import (
	"fmt"
	"runtime"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	defer fmt.Println("结束") //defer会被调用
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
		runtime.Goexit() //立即结束该协程
	}
}
func main() {
	//numbers()和alphabets()并发执行
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main Over")
}
