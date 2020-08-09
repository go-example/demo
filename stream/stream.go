package stream

import (
	"errors"
)

/**
流式工作原理：
各个任务都过指针链表的方式组成一个任务链，这个任务链从第一个开始执行，直到最后一个
每一个任务节点执行完毕会将结果带入到下一级任务节点中。
每一个任务是一个Stream节点，每个任务节点都包含首节点和下一个任务节点的指针,
除了首节点，每个节都会设置一个回调函数的指针，用本节点的任务执行,
最后一个节点的nextStream为空,表示任务链结束。
**/

//定回调函数指针的类型
type CB func(interface{}) (interface{}, error)

//任务节点结构定义
type Stream struct {
	//任务链表首节点,其他非首节点此指针永远指向首节点
	firstStream *Stream
	//任务链表下一个节点，为空表示任务结束
	nextStream *Stream
	//当前任务对应的执行处理函数，首节点没有可执行任务，处理函数指针为空
	cb CB
}

/**
创建新的流
**/
func NewStream() *Stream {
	//生成新的节点
	stream := &Stream{}
	//设置第一个首节点，为自己
	//其他节点会调用run方法将从firs指针开始执行，直到next为空
	stream.firstStream = stream
	//fmt.Println("new first", stream)
	return stream
}

/**
流结束
arg为流初始参数，初始参数放在End方法中是考虑到初始参数不需在任务链中传递
**/
func (this *Stream) Go(arg interface{}) (interface{}, error) {
	//设置为任务链结束
	this.nextStream = nil
	//fmt.Println("first=", this.firstStream, "second=", this.firstStream.nextStream)
	//检查是否有任务节点存在，存在则调用run方法
	//run方法是首先执行本任务回调函数指针，然后查找下一个任务节点，并调用run方法
	if this.firstStream.nextStream != nil {
		return this.firstStream.nextStream.run(arg)
	} else {
		//流式任务终止
		return nil, errors.New("Not found execute node.")
	}
}
func (this *Stream) run(arg interface{}) (interface{}, error) {
	//fmt.Println("run,args=", args)
	//执行本节点函数指针
	result, err := this.cb(arg)
	//然后调用下一个节点的Run方法
	if this.nextStream != nil && err == nil {
		return this.nextStream.run(result)
	} else {
		//任务链终端，流式任务执行完毕
		return result, err
	}
}
func (this *Stream) Next(cb CB) *Stream {
	//创建新的Stream，将新的任务节点Stream连接在后面
	this.nextStream = &Stream{}
	//设置流式任务链的首节点
	this.nextStream.firstStream = this.firstStream
	//设置本任务的回调函数指针
	this.nextStream.cb = cb
	//fmt.Println("next=", this.nextStream)
	return this.nextStream
}
