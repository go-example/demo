package main

import (
	"fmt"
	"io"
	"net"
)
func main() {

	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"//请求首行
	msg += "Host: www.baidu.com\r\n"//请求头部
	msg += "Connection: close\r\n"//请求头部
	msg += "\r\n\r\n"//再下面是请求数据,这里没有携带数据

	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}