package main

import (
	"fmt"

	"time"
)

func main() {

	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	for {
		a:=<-t.C
		fmt.Println("timer running..."+a.String())
		// 需要重置Reset 使 t 重新开始计时
		t.Reset(time.Second * 2)
	}
}