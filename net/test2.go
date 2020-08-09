package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var srcHost, agencyHost string
	flag.StringVar(&srcHost, "srcHost", "", "srcHost")
	flag.StringVar(&agencyHost, "agencyHost", "", "agencyHost")
	flag.Parse()
	if srcHost == "" || agencyHost == "" {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", "参数为空")
		os.Exit(1)
	}
	//建立socket，监听端口
	netListen, err := net.Listen("tcp", srcHost)
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		//如果没有请求就一直等待
		if err != nil {
			continue
		}
		if conn != nil {
			Log(conn.RemoteAddr().String(), " tcp connect success")
			go handleConnection(conn, agencyHost) //go 可以实现异步并发请求
		}
	}
}

//处理连接
func handleConnection(conn net.Conn, agencyHost string) {
	time.Sleep(10 * time.Millisecond)
	buffer := ReceiveData(conn)
	if len(buffer) > 1 {
		arr := strings.Split(string(buffer), "\r\n")
		if len(arr) > 1 {
			arr[1] = "Host: " + agencyHost
			newstr := strings.Join(arr, "\r\n")
			SendAgencyHost([]byte(newstr), agencyHost, conn)
		}
	}
	conn.Close()
}
func SendAgencyHost(data []byte, host string, baseconn net.Conn) {
	conn, _ := net.Dial("tcp", host)
	conn.Write(data)
	time.Sleep(10 * time.Millisecond)
	bufferHead := ReceiveData(conn)
	time.Sleep(10 * time.Millisecond)
	bufferBody := ReceiveData(conn)
	var buf bytes.Buffer
	buf.Write(bufferHead)
	buf.Write(bufferBody)
	baseconn.Write(buf.Bytes())
	conn.Close()
}

//接收数据统一方法
func ReceiveData(conn net.Conn) []byte {
	var buf bytes.Buffer
	buffer := make([]byte, 8192)
	for {
		sizenew, err := conn.Read(buffer)
		buf.Write(buffer[:sizenew])
		if err == io.EOF || sizenew < 8192 {
			break
		}
	}
	return buf.Bytes()
}

//打印信息统一方法
func Log(v ...interface{}) {
	log.Println(v...)
}

//执行错误处理方法
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}