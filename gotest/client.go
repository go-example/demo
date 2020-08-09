
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

//连接服务器
func connectServer() {
	//接通
	conn, err := net.Dial("tcp", "localhost:8888")
	checkError(err)
	fmt.Println("连接成功！\n")

	cInfo := "Hello.I am client     "
	buffer := make([]byte, 2048)
	buffer = []byte(cInfo)

	conn.Write(buffer)
	//向服务器发送消息
	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println(" connection error: ", err)
		return
	}
	//接收服务器发送消息
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println(" connection error: ", err)
		return
	}

	fmt.Println(string(buffer))
}

//检查错误
func checkError(err error) {
	if err != nil {
		log.Fatal("an error!", err.Error())
	}
}

//主函数
func main() {

	var goCount *int
	goCount = flag.Int("goCount", 5, "goroutine number")

	//解析输入的参数
	flag.Parse()
	fmt.Println("go count = ", *goCount)

	//get current time
	tInsert := time.Now()
	fmt.Println("tStart time: ", tInsert)

	for i := 0; i < *goCount; i++ {

		fmt.Println("goroutine number: ", i)
		//连接servser
		connectServer()
	}
	//获取时间差
	elapsed := time.Since(tInsert)
	fmt.Println("Insert elapsed: ", elapsed)
	fmt.Println("/n")

}