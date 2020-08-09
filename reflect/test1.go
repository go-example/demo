package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(1)
	s := t.Name()
	fmt.Println(s)
}