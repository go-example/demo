package main

import (
	"fmt"
	"time"
)

type Demo struct {
	input        chan string
	output       chan string
	maxGoroutine chan int
}

func NewDemo() *Demo {
	d := new(Demo)
	d.input = make(chan string, 24)
	d.output = make(chan string, 24)
	d.maxGoroutine = make(chan int, 20)
	return d
}

func (d *Demo) Goroutine() {
	var i = 1000
	for {
		d.input <- time.Now().Format("2006-01-02 15:04:05")
		time.Sleep(time.Second * 1)
		if i < 0 {
			break
		}
		i--
	}
	close(d.input)
}

func (d *Demo) Handle() {
	for t := range d.input {
		fmt.Println("datatime is :", t)
		d.output <- t
	}
}

func main() {
	demo := NewDemo()
	go demo.Goroutine()
	demo.Handle()
}
