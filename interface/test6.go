package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	var b interface{}
	fmt.Println(a)
	//a = make(map[string]string)
	str := "{\"aa\":\"a\",\"b\":\"bb\"}"
	json.Unmarshal([]byte(str),&a)
	fmt.Println(a)
	str1 := "1"
	json.Unmarshal([]byte(str1),&b)
	fmt.Println(b)

	ok,err := a.(map[string]interface{})
	fmt.Println(ok,err)
	fmt.Println(reflect.TypeOf(a))

	ok1,err1 := b.(string)
	fmt.Println(ok1,err1)
	fmt.Println(reflect.TypeOf(b))
}
