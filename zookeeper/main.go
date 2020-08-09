package main

import (
	"fmt"
	"github.com/go-zk/zk"
	"net/http"
	"time"
)

var hosts = []string{"127.0.0.1:2181"}

var path1 = "/whatzk"

var flags int32 = zk.FlagEphemeral
var data1 = []byte("hello,this is a zk go test demo!!!")
var acls = zk.WorldACL(zk.PermAll)

func main() {
	option := zk.WithEventCallback(callback)

	conn, _, err := zk.Connect(hosts, time.Minute*5, option)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}


	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	create(conn, path1, data1)

	time.Sleep(time.Second * 2)

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//delete(conn, path1)

	http.ListenAndServe("8080",nil)
}

func callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func create(conn *zk.Conn, path string, data []byte) {
	_, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}

}

func delete(conn *zk.Conn, path string) error{
	if err := conn.Delete(path,1);err !=nil{
		return err
	}
	return nil
}