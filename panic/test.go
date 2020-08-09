package main

import "fmt"

func main() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()
	f()              //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
	fmt.Println("a")
	if true {
		panic("异常信息")
	}
	fmt.Println("b") //这里开始下面代码不会再执行
}
