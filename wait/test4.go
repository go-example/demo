package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	i := 0
	for {
		if i == 5 {
			break
		}
		wg.Add(3)
		go func(n int) {
			fmt.Println("n:", n)
			t := time.Duration(n) * time.Second
			time.Sleep(t)

			wg.Done()
		}(1)

		go func(n int) {
			fmt.Println("n:", n)
			t := time.Duration(n) * time.Second
			time.Sleep(t)
			time.Sleep(t)

			wg.Done()
		}(2)

		go func(n int) {
			fmt.Println("n:", n)
			t := time.Duration(n) * time.Second
			time.Sleep(t)

			wg.Done()
		}(3)

		wg.Wait()
		i++
	}

	fmt.Println("main exit...")
}
