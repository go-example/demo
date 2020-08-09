package main

import (
	"fmt"
	"reflect"
)

type T struct {
	
} 

func main() {
	var t *T
	var i interface{} = t
	fmt.Println(i)
	fmt.Println(t)
	fmt.Println(t == nil, i == t, i == nil,isNil(i),isNil(t))
}

func isNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}