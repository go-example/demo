package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
}