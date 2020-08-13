package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	go w.Write([]byte("hello word"))

	data := make([]byte, 5)
	n, err := r.Read(data)
	fmt.Println("next read number", n) //next read number 5
	w.Close()	//此次关闭，就再也读不出了
	if err == io.EOF {
		fmt.Println("executing read return EOF")
		fmt.Println("executing read reads number", n)
	}
	n, _ = r.Read(data)
	fmt.Println(string(data))          //hello word
	fmt.Println("next read number", n) //next read number 0
}
