package main

import "fmt"

func main() {
	a := 3.102
	b := fmt.Sprintf("%.2f", a)
	fmt.Println(b)
	fmt.Println(1)

	var data = make(map[string]map[string]int)
	data["sid"] = make(map[string]int)
	data["sid"]["pid"] = 1

	fmt.Println(data["sid"]["pid"])

	fmt.Println(data["sid"]["pid1"])

}