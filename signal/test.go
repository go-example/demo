package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT:
			log.Println("SIGSTOP")
			return
		case syscall.SIGHUP:
			log.Println("SIGHUP")
			return
		case syscall.SIGKILL:
			log.Println("SIGKILL")
			return
		default:
			log.Println("default")
			return
		}
	}
}
