package main

import (
	"container/list"
	"fmt"
)

func main()  {
	link := list.New()

	for i:=0; i<5 ; i++ {
		fmt.Println("插入元素: ", link.PushBack(i).Value)
	}
	fmt.Println("插入完毕，链表长度：", link.Len())  // 10

	for i:=0; i<5 ; i++ {
		fmt.Println("插入元素: ", link.PushBack(i).Value)
	}
	fmt.Println("插入完毕，链表长度：", link.Len())  // 20

	// 首尾互换
	link.MoveToBack(link.Front())
	link.MoveToFront(link.Back().Prev())
	fmt.Println(link.Front().Value, link.Back().Value)

	for v:=link.Front();v!=nil;v=v.Next(){
		fmt.Println(v.Value)
	}
}