package main

import "fmt"

type Student struct {
	Name string
}
var b interface{} = Student{
	Name:     "aaa",
}
var c  = b.(Student)

func main()  {
	c.Name = "bbb"
	fmt.Println(b.(Student).Name)
}