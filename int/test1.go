package main

import "fmt"

func main() {
	var a int8 = 127
	fmt.Println(a + 1)

	a = 1
	a += 010
	fmt.Println(a)
}
