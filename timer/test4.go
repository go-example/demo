package main

import (
	"fmt"
	"time"
)



func main()  {
	tchan := time.After(time.Second*3)
	fmt.Printf("tchan type=%T\n",tchan)
	fmt.Println("mark 1")
	fmt.Println("tchan=",<-tchan)
	fmt.Println(<-time.After(2*time.Second))
	var t  time.Time
	for  {
		select {
		case t=<-time.After(time.Second*2):
			fmt.Println("当前时间",t.String())
		}
	}
}