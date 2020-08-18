package main

import "fmt"

type A struct {
	Id int `json:"id"`
}

type B struct {}

type xx map[B]int

func main() {
	var a A
	bb := &a
	cc := &a
	dd := &A{}
	ee := &A{}

	m := make(xx)
	m[B{}] = 123

	fmt.Println(fmt.Sprintf("%+v,%p", &a, &a))
	fmt.Println(fmt.Sprintf("%+v,%p", bb, bb))
	fmt.Println(fmt.Sprintf("%+v,%p", cc, cc))
	fmt.Println(fmt.Sprintf("%+v,%p", dd, dd))
	fmt.Println(fmt.Sprintf("%+v,%p", ee, ee))
	fmt.Println(fmt.Sprintf("%v", &a))
	fmt.Println(fmt.Sprintf("%v", bb))
}
