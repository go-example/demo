package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Urls struct {
	Key  string
	Url  string
	Data string
}

func requestUrl(u *Urls, c chan *Urls) {
	u.Data, _ = Get(u.Url)
	c <- u
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	i := r.URL.Query().Get("i")
	time.Sleep(1 * time.Millisecond)
	fmt.Fprintln(w, "hello"+i)
}

func main() {
	start_time := Millisecond()
	go func() {
		http.HandleFunc("/", IndexHandler)
		http.ListenAndServe(":8080", nil)
	}()

	c := make(chan *Urls)

	for i := 0; i < 5; i++ {
		url := &Urls{}
		url.Key = strconv.Itoa(i)
		url.Url = "http://localhost:8080?i=" + strconv.Itoa(i)
		go requestUrl(url, c)
	}

	for u := range c {
		fmt.Println(u)
	}
	close(c)
	fmt.Println(Millisecond() - start_time)
}

//微秒(当前)
func Millisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

//通用http请求
func Request(url string, method string, data string, header map[string]string) (re string, response *http.Response) {
	if method == "" {
		method = "POST"
	}
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		log.Printf("url:%s,error_info:%s\n", url, err.Error())
		return "", nil
	}
	defer req.Body.Close()
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	if req.Header.Get("contentType") == "" {
		req.Header.Add("contentType", "application/json")
	}

	clt := http.Client{}
	resp, _ := clt.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("url:%s,error_info:%s\n", url, err.Error())
		return "", nil
	}

	if resp.StatusCode != 200 {
		log.Printf("url:%s,response:%s\n", url, string(body))
	}
	return string(body), resp
}

//get
func Get(url string) (re string, response *http.Response) {
	return Request(url, "GET", "", map[string]string{"Content-Type": "application/json"})
}
