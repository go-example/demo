package main

import (
	"fmt"
	"net"
)

func main()  {
	l, _ := net.Listen("tcp", "localhost:0")
	ip := l.Addr().String()
	fmt.Println(ip)
}
