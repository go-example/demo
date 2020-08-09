package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建一个 list
	l := list.New()
	//把4元素放在最后
	e4 := l.PushBack(4)
	//把1元素放在最前
	e1 := l.PushFront(1)
	//在e4元素前面插入3
	l.InsertBefore(3, e4)
	//在e1后面插入2
	l.InsertAfter(2, e1)
	// 遍历所有元素并打印其内容
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}