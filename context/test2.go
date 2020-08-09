package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	go task(ctx)
	time.Sleep(time.Second * 10)
}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	go func() {
		// 模拟4秒耗时任务
		time.Sleep(time.Second * 4)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}