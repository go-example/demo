package main

import (
	"fmt"
	"log"
	"net"
)

//开启服务程序
func startServer() {

	var i int = 0
	//服务端监听ip和端口号
	listener, err := net.Listen("tcp", "localhost:8888")
	checkError(err)
	defer listener.Close()

	fmt.Println("监听成功")
	for {
		i++
		fmt.Println("before accept")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		//开辟一个协程处理客户端请求
		go doServerStuff(conn, i)
	}
}

//处理客户端信息
func doServerStuff(conn net.Conn, num int) {

	fmt.Println("accept num:", num)

	clientInfo := make([]byte, 2048)
	//读取客户端发来消息
	_, err := conn.Read(clientInfo)
	if err != nil {
		fmt.Println(" connection error: ", err)
		return
	}

	fmt.Println(string(clientInfo))
	sInfo := "I am server, hello"
	clientInfo = []byte(sInfo)
	//向客户端发送消息
	_, err = conn.Write(clientInfo)
	if err != nil {
		fmt.Println(" connection error: ", err)
		return
	}
}

//检查错误
func checkError(err error) {
	if err != nil {
		log.Fatal("an error!", err.Error())
	}
}

//主函数
func main() {
	//开启服务程序
	startServer()
}