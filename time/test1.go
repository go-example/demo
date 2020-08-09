package main

import (
	"fmt"
	"time"
)

func updateDateTime()  {
	go func() {
		select {
		case <-time.After(time.Second*2):
			updateDateTime()
		}
	}()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	updateDateTime()
	time.Sleep(5*time.Second)
}