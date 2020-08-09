package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//加锁，注意锁要以指针的形式传进来，不然只是拷贝
func total1(num *int, mu *sync.Mutex, ch chan bool) {
	mu.Lock();
	for i := 0; i < 1000; i++ {
		*num += i;
	}
	ch <- true;
	mu.Unlock();
}

//不加锁
func total2(num *int, ch chan bool) {
	for i := 0; i < 1000; i++ {
		*num += i;
	}
	ch <- true;
}

//Lock、Unlock与RLock、RUnlock不能嵌套使用
func total3(num *int, rwmu *sync.RWMutex, ch chan bool) {
	for i := 0; i < 1000; i++ {
		rwmu.Lock();
		*num += i;
		rwmu.Unlock();
		if (i == 500) {
			//读锁定
			rwmu.RLock();
			fmt.Print(*num, " ");
			rwmu.RUnlock();
		}
	}
	ch <- true;
}

func printNum(num int, cond *sync.Cond) {
	cond.L.Lock();
	if num < 5 {
		//num小于5时，进入等待状态
		cond.Wait();
	}
	//大于5的正常输出
	fmt.Println(num);
	cond.L.Unlock();
}

func main() {
	//Once.Do()保证多次调用只执行一次
	once := sync.Once{};
	ch := make(chan bool, 3);
	for i := 0; i < 3; i++ {
		go func(n int) {
			once.Do(func() {
				//只会执行一次，因为闭包引用了变量n，最后的值为2
				fmt.Println(n)
			});
			//给chan发送true，表示执行完成
			ch <- true;
		}(i);
	}
	for i := 0; i < 3; i++ {
		//读取三次chan,如果上面三次没执行完会一直阻塞
		<-ch;
	}

	//互斥锁，保证某一时刻只能有一个访问对象
	mutex := sync.Mutex{};
	ch2 := make(chan bool, 20);
	//使用多核，不然下面的结果会一样
	runtime.GOMAXPROCS(runtime.NumCPU());
	num1 := 0;
	num2 := 0;
	for i := 0; i < 10; i++ {
		go total1(&num1, &mutex, ch2);
	}
	for i := 0; i < 10; i++ {
		go total2(&num2, ch2);
	}
	for i := 0; i < 20; i++ {
		<-ch2;
	}
	//会发现num1与num2计算出的结果不一样
	//而num1的结果才是正确的，因为total2没有加锁，导致多个goroutine操作num时发生数据混乱
	fmt.Println(num1, num2);

	//读写锁，多了读锁定，和读解锁，让多个goroutine同时读取对象
	rwmutex := sync.RWMutex{};
	ch3 := make(chan bool, 10);
	num3 := 0;
	for i := 0; i < 10; i++ {
		go total3(&num3, &rwmutex, ch3);
	}
	for i := 0; i < 10; i++ {
		<-ch3;
	}
	fmt.Println(num3);

	//组等待，等待一组goroutine的结束
	wg := sync.WaitGroup{};
	//增加计数器
	wg.Add(10);
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Print(n, " ");
			//这里表示该goroutine执行完成
			wg.Done();
		}(i);
	}
	//等待所有线程执行完成
	wg.Wait();
	fmt.Println("");

	//条件等待
	mutex2 := sync.Mutex{};
	//使用锁创建一个条件等待
	cond := sync.NewCond(&mutex2);
	for i := 0; i < 10; i++ {
		go printNum(i, cond);
	}

	time.Sleep(time.Second * 1);
	//等待一秒后，我们先唤醒一个等待，输出一个数字
	cond.L.Lock()
	cond.Signal();
	cond.L.Unlock();
	time.Sleep(time.Second * 1);
	//再次待待一秒后，唤醒所有，输出余下四个数字
	cond.L.Lock()
	cond.Broadcast();
	cond.L.Unlock();
	time.Sleep(time.Second * 1);
}
