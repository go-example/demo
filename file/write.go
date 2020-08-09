package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	content := []byte("hello golang")
	//将指定内容写入到文件中
	err := ioutil.WriteFile("./test.log", content, 0666)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
}