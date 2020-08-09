package main

import (
	"fmt"
	"time"
)


var name string = "start"

func stop()  {
	name = "stop"
}

func main()  {
	time.AfterFunc(2*time.Second, func(){
		stop()
		fmt.Println("2秒后停止")
	})
	fmt.Println(name)
	time.Sleep(5*time.Second)
	fmt.Println(name)
}
