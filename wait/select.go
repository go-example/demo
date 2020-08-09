package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server2"
}
func end(ch chan bool) {
	time.Sleep(3 * time.Second)
	ch <- true
}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	output3 := make(chan bool)
	go server1(output1)
	go end(output3)
	go server2(output2)
	//随机选择oo
	for {
		select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		case s3 := <-output3:
			if s3 == true {
				fmt.Println("退出啦")
				goto Loop
			}
		}
	}
Loop:
}
