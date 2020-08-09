package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type aa struct {
	Key                 string
	CheckTime           string `json:"check_time"`            //检查版本时间间隔（秒）
	VersionGenerateTime int    `json:"version_generate_time"` //版本生成间隔(完成一次版本生效后)
	KeepVersionNum      int    `json:"keep_version_num"`      //保留的版本数
}

func main() {
	a := &aa{Key: "aaa"}
	set(a)

	var b *string
	set1(b)

	var c *int
	set1(c)

	set2(nil)

	setF(func(s string) interface{} {
		f := &aa{}
		_ = json.Unmarshal([]byte(s), &f)
		return f
	})
	//fmt.Println(a)
}

func set(a interface{}) {
	str := "{\"check_time\":\"5\",\"version_generate_time\":600,\"keep_version_num\":6,\"key\":\"aaaa\"}"

	err := json.Unmarshal([]byte(str), &a)

	fmt.Println(a, reflect.TypeOf(a), err)
}

func set1(a interface{}) {
	str := "1"
	fmt.Println(reflect.TypeOf(a))
	err := json.Unmarshal([]byte(str), &a)

	fmt.Println(a, reflect.TypeOf(a), err)
}

func set2(a interface{}) {
	var err error
	str := "1"
	if a != nil {
		err = json.Unmarshal([]byte(str), &a)
	}
	fmt.Println(a, reflect.TypeOf(a), err)
}

func setF(f func(s string) interface{}) {
	var data []interface{}
	for i := 0; i < 5; i++ {
		str := "{\"check_time\":\"" + time.Now().Format("2006-01-02 15:04:05") + "\",\"version_generate_time\":600,\"keep_version_num\":6,\"key\":\"aaaa\"}"
		data = append(data, f(str))
		time.Sleep(1*time.Second)
	}
	fmt.Println(data, reflect.TypeOf(data))
	for _, item := range data {
		fmt.Println(item)
	}
}
