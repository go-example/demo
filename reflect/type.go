package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "咖啡色的羊驼"

	// TypeOf会返回目标数据的类型，比如int/float/struct/指针等
	reflectType := reflect.TypeOf(name)

	// valueOf返回目标数据的的值，比如上文的"咖啡色的羊驼"
	reflectValue := reflect.ValueOf(name)

	fmt.Println("type: ", reflectType)
	fmt.Println("value: ", reflectValue)
}
