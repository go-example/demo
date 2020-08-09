package main

import (
	"fmt"
	"github.com/go-zk/zk"
	"time"
)

func aaa(event zk.Event) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("<<<<<<<<<<<<<<<<<<<")
}

func main() {
	ZKOperateWatchTest()
}

func ZKOperateWatchTest() {
	fmt.Printf("ZKOperateWatchTest\n")

	option := zk.WithEventCallback(aaa)
	var hosts = []string{"192.168.40.51:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var path1 = "/zk_test_go1"
	var data1 = []byte("zk_test_go1_data1")
	exist, s, _, err := conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("path[%s] exist[%t]\n", path1, exist)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", ZkStateStringFormat(s))

	// try create
	var acls = zk.WorldACL(zk.PermAll)
	p, err_create := conn.Create(path1, data1, zk.FlagEphemeral, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
	fmt.Printf("created path[%s]\n", p)

	time.Sleep(time.Second * 2)

	exist, s, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("path[%s] exist[%t] after create\n", path1, exist)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", ZkStateStringFormat(s))

	// delete
	conn.Delete(path1, s.Version)

	exist, s, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("path[%s] exist[%t] after delete\n", path1, exist)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", ZkStateStringFormat(s))
}

func ZkStateString(s *zk.Stat) string {
	return fmt.Sprintf("Czxid:%d, Mzxid: %d, Ctime: %d, Mtime: %d, Version: %d, Cversion: %d, Aversion: %d, EphemeralOwner: %d, DataLength: %d, NumChildren: %d, Pzxid: %d",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}

func ZkStateStringFormat(s *zk.Stat) string {
	return fmt.Sprintf("Czxid:%d\nMzxid: %d\nCtime: %d\nMtime: %d\nVersion: %d\nCversion: %d\nAversion: %d\nEphemeralOwner: %d\nDataLength: %d\nNumChildren: %d\nPzxid: %d\n",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}