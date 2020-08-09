package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println("https://github.com/rushuinet/webdemo")
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}