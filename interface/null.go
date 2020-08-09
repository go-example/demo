package main

import (
	"fmt"
)

func describe1(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	// 任何类型的变量传入都可以

	s := "Hello World"
	i := 55
	strt := &struct {
		name string
	}{
		name: "Naveen R",
	}
	describe1(s)
	describe1(i)
	describe1(strt)
}