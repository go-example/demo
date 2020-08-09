package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 手动取消
	a := 1
	b := 2
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		cancel() // 在调用处主动取消
	}()
	res := Add(ctx, 1, 2)
	fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
}