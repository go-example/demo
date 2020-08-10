package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		start := time.Now().UnixNano() / 1e6
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		useTime := time.Now().UnixNano()/1e6 - start

		str := "请求用时" + fmt.Sprintf("%d",useTime) + "毫秒"
		log.Println(str)
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(simpleMw)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ArticlesHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("articles")
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("home")
	time.Sleep(time.Second)
}

func ProductsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("products")
}
