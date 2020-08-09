package main

import "fmt"

type Skills interface {
	Running()
	GetName() string
}

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name   string
	Salary int
}

func (p Student) GetName() string { //实现Getname方法
	fmt.Println(p.Name)
	return p.Name
}

func (p Student) Running() { // 实现 Running方法
	fmt.Printf("%s running", p.Name)
}

func (p Teacher) GetName() string { //实现Getname方法
	fmt.Println(p.Name)
	return p.Name
}

func (p Teacher) Running() { // 实现 Running方法
	fmt.Printf("\n%s running", p.Name)
}
func main() {
	var skill Skills
	var stu1 Student
	var t1 Teacher
	t1.Name = "wang"
	stu1.Name = "wd"
	stu1.Age = 22
	skill = stu1
	skill.Running()
	skill = t1
	t1.Running()
}

//wd running
//wang running
