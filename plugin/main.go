package main

import (
	"fmt"
	"plugin"
)

func main() {

	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	m, err := p.Lookup("GetProductBasePrice")
	if err != nil {
		panic(err)
	}
	res := m.(func(int) int)(30)
	fmt.Println(res)
}