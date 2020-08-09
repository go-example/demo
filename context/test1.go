package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// ...
		// 你的逻辑
		// ...
		fmt.Println("1")
		select {
		case <-ctx.Done():
		}
	}()
	fmt.Println("aa")
	cancel()
	fmt.Println("2")
}