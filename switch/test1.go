package main

import "fmt"

func main() {
	val := 4
	switch val {
	case 3,4:
		fmt.Println("case 3")
		fallthrough
	case 6:
		fmt.Println("case 4")
	case 5:
		fmt.Println("case 5")
	}
}
//case 3
//case 4