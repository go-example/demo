package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(5 * time.Second) // sleep one second
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")

	}
}

func CombineChannel(ch1, ch2 chan int) <-chan int {
	out := make(chan int, 3)
	go func() {
		defer close(out)
		for {
			select {
			case v1, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- v1
			case v2, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v2

			}

			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()
	return out
}

// 管理类
type PlayerManger struct {
	players sync.Map //据说这个map是线程安全的

	//players map[uint64]*Player
}

// 单例实现
var once sync.Once

var manger *PlayerManger = nil

func GetPlayerManger() *PlayerManger {

	once.Do(func() {
		manger = &PlayerManger{} //manger.players = make(map[uint64]*Player)
	})
	return manger

}
