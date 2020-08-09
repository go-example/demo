package main

import (
	"fmt"
)

type Test interface {
	Tester()
}

type MyFloat float64
type MString string

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func (m MString) Tester ()  {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface 类型 %T ,  值： %v\n", t, t)
}

func main() {
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()

	s := MString("aa")
	describe(s)
	s.Tester()

	a := []byte(s)
	println(a)

}