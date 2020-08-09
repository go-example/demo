package main

import (
	"fmt"
	"github.com/go-example/demo/timingwheel"
	"time"
)

func main() {

	//这里我们创建了一个timingwheel，精度是1s，最大的超时等待时间为3600s
	w := timingwheel.NewTimingWheel(1*time.Second, 3600)

	for {
		select {
		//等待10s
		case <-w.After(3 * time.Second):
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
