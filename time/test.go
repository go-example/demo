package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	t := time.Now()   //2019-07-31 13:55:21.3410012 +0800 CST m=+0.006015601
	fmt.Println(t.Format("20060102150405"))

	//当前时间戳
	t1 := time.Now().Unix()  //1564552562
	fmt.Println(t1)
	//时间戳转化为具体时间
	fmt.Println(time.Unix(t1, 0).String())

	//基本格式化的时间表示
	fmt.Println(time.Now().String()) //2019-07-31 13:56:35.7766729 +0800 CST m=+0.005042501

	fmt.Println(time.Now().Format("2006-01-02"))  //2019-07-31
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))  //2019-07-31 13:57:52
	fmt.Println(time.Now().Format("01-02"))  //2019-07-31 13:57:52

	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n",time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n",time.Now().UnixNano() / 1e9)

	t1 = 1575623284882621200
	println(time.Unix(t1, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(t1, 0).Weekday())

	TimeF("2019-12-06",0)
}

func TimeF(format string,time int)  {
	fmt.Println()
	//var bt = "2006-01-02 15:04:05"
	format = "Y-m-d"
	a := "Y-m-d H:i:s"
	println(a)
	b := strings.ReplaceAll(a,"Y","2016")
	fmt.Println(b)
}