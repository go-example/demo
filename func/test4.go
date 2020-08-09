package main

import (
	"fmt"
	"time"
)

func test() {
	a := []int{1, 2, 3}
	for _, v := range a {
		go func() {
			fmt.Printf("%d ", v)
		}()
	}
}

func main() {
	test()
	time.Sleep(time.Second)
}

// 3 3 3