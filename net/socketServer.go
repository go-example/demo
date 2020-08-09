package main

import (
	"log"
	"net"
)

func chkError1(err error) {
	if err != nil {
		log.Fatal(err);
	}
}

func main() {
	//创建一个TCP服务端
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080");
	chkError1(err);
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
	chkError1(err2);
	//死循环的处理客户端请求
	for {
		//等待客户的连接
		//注意这里是无法并发处理多个请求的
		conn, err3 := tcplisten.Accept();
		//如果有错误直接跳过
		if err3 != nil {
			continue;
		}

		//向客户端发送数据，并关闭连接
		_, _ = conn.Write([]byte("hello,client \r\n"))
		_ = conn.Close()
	}
}