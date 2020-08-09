package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func say(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str)
	}
}

func sayStat(str string, ch chan int64) {
	for i := 0; i < 5000; i++ {
		runtime.Gosched()
		fmt.Println(str)
		ch <- int64(i)
	}
	close(ch)
}

func sayStat_2_Worker(str string, ch chan int64) {
	sum := 0
	for i := 0; i < 5000; i++ {
		runtime.Gosched()
		fmt.Println(str)
		sum += i
	}
	ch <- int64(sum)
	//    close(ch)
}

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, i := range nums {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for c := range in {
			select {
			case out <- c * c:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, ins ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(ins))
	out := make(chan int)
	// ERROR: http://studygolang.com/articles/7994
	// REF:   "for"声明中的迭代变量和闭包
	//    for _, in := range ins {
	//        go func() {
	//            for c := range in {
	//                out <- c
	//            }
	//            wg.Done()
	//        }()
	//    }
	// Solution1: New func Outline
	//    ff := func(in <-chan int) {
	//        for c := range in {
	//            out <- c
	//        }
	//        wg.Done()
	//    }
	//    for _, in := range ins {
	//        go ff(in)
	//    }
	// Solution2: Inline func with parameter
	//    for _, in := range ins {
	//        go func(in <-chan int) {
	//            for c := range in {
	//                out <- c
	//            }
	//            wg.Done()
	//        }(in)
	//    }
	// Solution3: Inline func with parameter copy bak
	for _, in := range ins {
		in_copy := in
		go func() {
			defer wg.Done()
			for c := range in_copy {
				select {
				case out <- c:
				case <-done:
					return
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func genNew(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func main() {
	// DEFAULT VALUE: NUMBER OF CPU CORE
	fmt.Println(runtime.GOMAXPROCS(-1))
	runtime.Gosched()
	fmt.Println(runtime.GOMAXPROCS(-1))
	fmt.Println(runtime.NumCPU())

	//    go say("hello")
	//    say("world")

	ch := make(chan int64)
	go sayStat("hello", ch)
	//    go sayStat("hello", ch)
	//    sayStat("world", ch)
	var stat int64 = 0
	for c := range ch {
		fmt.Println(c)
		stat += c
	}
	fmt.Println(stat) // 12497500

	//    // DEAD LOCK !
	//    cc := make(chan int)
	//    // NO GOROUTINE RECEIVE THE UNBUFFERED CHANNEL DATA !
	//    cc <- 888
	//    fmt.Println(<-cc)

	stat = 0
	cc := make(chan int64)
	worker_num := 2
	for i := 0; i < worker_num; i++ {
		go sayStat_2_Worker("TEST-"+strconv.Itoa(i), cc)
	}
	for i := 0; i < worker_num; i++ {
		stat += <-cc
	}
	close(cc)
	fmt.Println(stat) // 12497500 * 2 = 24995000

	//    out := square(gen(1, 2, 3, 4, 5))
	//    for c := range out {
	//        fmt.Println(c)
	//    }

	done := make(chan struct{})
	//    defer close(done)

	out_new := gen(done, 1, 2, 3, 4, 5)
	c1 := square(done, out_new)
	c2 := square(done, out_new)
	//    for r1 := range c1 {
	//        fmt.Println(r1)
	//    }
	//    for r2 := range c2 {
	//        fmt.Println(r2)
	//    }
	//    for r := range merge(c1, c2) {
	//        fmt.Println(r)
	//    }
	mg := merge(done, c1, c2)
	fmt.Println(<-mg)
	fmt.Println(<-mg)
	fmt.Println(<-mg)
	close(done)
	//    fmt.Println(<-mg)
	//    fmt.Println(<-mg)
	//    fmt.Println(<-mg)
	//    fmt.Println(<-mg)
	for {
		if msg, closed := <-mg; !closed {
			fmt.Println("<-mg has closed!")
			return
		} else {
			fmt.Println(msg)
		}
	}

	//    gen_new := genNew(1, 2, 3, 4, 5)
	//    //    close(gen_new)
	//    for gn := range gen_new {
	//        fmt.Println(gn)
	//    }
}