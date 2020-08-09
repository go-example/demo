package main

import "fmt"

func main() {
	//创建一个容量为3的缓冲信道
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))   //capacity is 3
	fmt.Println("length is", len(ch))     //length is 2
	fmt.Println("read value", <-ch)       //read value naveen
	fmt.Println("new length is", len(ch)) //new length is 1
}
