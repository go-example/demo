package main

import "fmt"

type ITest interface {
	Add() string
	Bb() string
}

type A struct {

}

func (a *A)Add() string  {
	return "A"
}

type B struct {

}

func (b *B)Add() string {
	return "bb"
}

func (b *B)Bb() string {
	return "bba"
}


var _ ITest = (*A)(nil)
var _ ITest = (*B)(nil)

var _ ITest = &A{}
var _ ITest = &B{}

func main()  {
	a := &A{}
	as := a.Add()
	fmt.Println(as)
}
