package main

import (
    "reflect"
    "fmt"
)

type Student1 struct {
    Id   int
    Name string
}

type People struct {
    Student1 // 匿名字段
}

func main() {
    p := People{Student1{Id: 1, Name: "咖啡色的羊驼"}}

    t := reflect.TypeOf(p)
    // 这里需要加一个#号，可以把struct的详情都给打印出来
    // 会发现有Anonymous:true，说明是匿名字段
    fmt.Printf("%#v\n", t.Field(0))

    // 取出这个学生的名字的详情打印出来
    fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

    // 获取匿名字段的值的详情
    v := reflect.ValueOf(p)
    fmt.Printf("%#v\n", v.Field(0))
}