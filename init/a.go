package main

import "fmt"


func init() { //3.1


	fmt.Printf("init WhatIsThe in a.go `s init 3.1: %d\n", 2)

}

func init() { //3.2


	fmt.Printf("init WhatIsThe in a.go`s init 3.2: %d\n", 3)

}

func main()  {
	fmt.Println(1)
}