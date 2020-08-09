// author songaimin@outlook.com 20170623
// golang 使用unixsocket交换数据
// 单元测试示例：
// package utils

// import (
//  "fmt"
//  u "lpaiche.com/utils"
//  // "os"
//  "testing"
//  "time"
// )

// func TestServer(t *testing.T) {
//  //声明unixsocket
//  us := u.NewUnixSocket("/tmp/us.socket")
//  //设置服务端接收处理
//  us.SetContextHandler(func(context string) string {
//      fmt.Println(context)
//      now := time.Now().String() + "sssssss"
//      return now
//  })
//  //开始服务
//  go us.StartServer()
//  time.Sleep(time.Second * 30)
// }

// func TestClient(t *testing.T) {
//  //声明unixsocket
//  us := u.NewUnixSocket("/tmp/us.socket")
//  //发送数据unixsocket并返回服务端处理结果
//  r := us.ClientSendContext("eeeeeee")
//  fmt.Println("===============" + r)
// }

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	us := NewUnixSocket("./us.sock", 10000)
	us.SetContextHandler(func(s string) string {
		fmt.Println(s)
		now := time.Now().String() + "sssssss"
		return now
	})
	go us.createServer()
	time.Sleep(2*time.Millisecond)
	echo := us.ClientSendContext("hello")
	fmt.Println(echo)
}

type UnixSocket struct {
	filename string
	bufsize  int
	handler  func(string) string
}

func NewUnixSocket(filename string, size ...int) *UnixSocket {
	size1 := 10480
	if size != nil {
		size1 = size[0]
	}
	us := UnixSocket{filename: filename, bufsize: size1}
	return &us
}

func (this *UnixSocket) createServer() {
	os.Remove(this.filename)
	addr, err := net.ResolveUnixAddr("unix", this.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	listener, err := net.ListenUnix("unix", addr)
	defer listener.Close()
	if err != nil {
		panic("Cannot listen to unix domain socket: " + err.Error())
	}
	fmt.Println("Listening on", listener.Addr())
	for {
		c, err := listener.Accept()
		if err != nil {
			panic("Accept: " + err.Error())
		}
		go this.HandleServerConn(c)
	}

}

//接收连接并处理
func (this *UnixSocket) HandleServerConn(c net.Conn) {
	defer c.Close()
	buf := make([]byte, this.bufsize)
	nr, err := c.Read(buf)
	if err != nil {
		panic("Read: " + err.Error())
	}
	// 这里，你需要 parse buf 里的数据来决定返回什么给客户端
	// 假设 respnoseData 是你想返回的文件内容
	result := this.HandleServerContext(string(buf[0:nr]))
	_, err = c.Write([]byte(result))
	if err != nil {
		panic("Writes failed.")
	}
}

func (this *UnixSocket) SetContextHandler(f func(string) string) {
	this.handler = f
}

//接收内容并返回结果
func (this *UnixSocket) HandleServerContext(context string) string {
	if this.handler != nil {
		return this.handler(context)
	}
	now := time.Now().String()
	return now
}

func (this *UnixSocket) StartServer() {
	this.createServer()
}

//客户端
func (this *UnixSocket) ClientSendContext(context string) string {
	addr, err := net.ResolveUnixAddr("unix", this.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	//拔号
	c, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("DialUnix failed.")
	}
	//写出
	_, err = c.Write([]byte(context))
	if err != nil {
		panic("Writes failed.")
	}
	//读结果
	buf := make([]byte, this.bufsize)
	nr, err := c.Read(buf)
	if err != nil {
		panic("Read: " + err.Error())
	}
	return string(buf[0:nr])
}
