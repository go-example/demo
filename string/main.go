package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("字符串测试")

	fmt.Println("字符串转化")
	//获取程序运行的操作系统平台下 int 类型所占的位数，如：strconv.IntSize。
	//strconv.IntSize

	fmt.Println("将字符串转换为 int 型。")
	var trastr01 string = "100"
	traint01, err_tra := strconv.Atoi(trastr01)
	if err_tra != nil {
		fmt.Println(err_tra)
	} else {
		fmt.Println(traint01)
	}
	fmt.Println("将字符串转换为 float64 型")
	var trastr02 string = "100.55"
	trafloat01, err_float := strconv.ParseFloat(trastr02, 10)
	if err_float != nil {
		fmt.Println(err_float)
	} else {
		fmt.Println(trafloat01)
	}
	trastr03 := strconv.Itoa(99)
	fmt.Println("int 转字符安 " + trastr03)

	var str01 string = "hello,world"
	str02 := "你好,世界"
	fmt.Println(str01)
	fmt.Println(str02)

	//

	//字符串比较
	com01 := strings.Compare(str01, str02)
	if com01 == 0 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等 " + string(com01))
	}
	fmt.Println(com01)

	//查找 包含
	var isCon bool = strings.Contains(str01, "hello")
	fmt.Println(isCon) //true

	//查找位置
	var theIndex int = strings.Index(str01, ",")
	fmt.Println(theIndex)                     //5
	fmt.Println(strings.Index(str01, "haha")) //不存在返回-1

	lastIndex := strings.LastIndex(str01, "o")
	fmt.Println("在字符串中最后出现位置的索引 " + strconv.Itoa(lastIndex)) //7
	//-1 表示字符串 s 不包含字符串

	//统计给定子串sep的出现次数, sep为空时, 返回1 + 字符串的长度
	fmt.Println(strings.Count("cheeseeee", "ee")) // 3
	fmt.Println(strings.Count("five", ""))        // 5

	// 重复s字符串count次, 最后返回新生成的重复的字符串
	fmt.Println("hello " + strings.Repeat("world ", 10))

	fmt.Println("替换")
	// 在s字符串中, 把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	var str03 string = "/Users//Documents/GOPatch/src/MyGO/config/TestString/"
	str04 := strings.Replace(str03, "/", "**", -1)
	str05 := strings.Replace(str03, "/", "**", 4)

	fmt.Println(str04) //**Users****Documents**GOPatch**src**MyGO**config**TestString**
	fmt.Println(str05) //**Users****Documents**GOPatch/src/MyGO/config/TestString/

	fmt.Println("删除字符串的开头和尾部")
	fmt.Println("删除两头的/ = " + strings.Trim(str03, "/"))      //Users//Documents/GOPatch/src/MyGO/config/TestString
	fmt.Println("删除左边的/ =  " + strings.TrimLeft(str03, "/")) //Users//Documents/GOPatch/src/MyGO/config/TestString/
	//还有 TrimRight

	str06 := strings.TrimSpace(" hello hao hao hao ")
	fmt.Println("删除开头末尾的空格 =" + str06) //'hello hao hao hao'

	fmt.Println("大小写")
	str07 := "hello hao hao hao"
	fmt.Println(strings.Title(str07))                  //Hello Hao Hao Hao
	fmt.Println(strings.ToLower(" Hello Hao Hao Hao")) // hello hao hao hao
	fmt.Println(strings.ToUpper(str07))                //HELLO HAO HAO HAO

	//前缀 后缀
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasSuffix("Amigo", "go"))  // true

	fmt.Println("字符串分割")
	fieldsStr := "  hello   it's  a  nice day today    "
	//根据空白符分割,不限定中间间隔几个空白符
	fieldsSlece := strings.Fields(fieldsStr)
	fmt.Println(fieldsSlece) //[hello it's a nice day today]

	for i, v := range fieldsSlece {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}
	for i := 0; i < len(fieldsSlece); i++ {
		fmt.Println(fieldsSlece[i])
	}

	//根据特定字符分割
	slice01 := strings.Split("q,w,e,r,t,y,", ",")
	fmt.Println(slice01)      //[q w e r t y ]
	fmt.Println(cap(slice01)) //7  最后多个空""
	for i, v := range slice01 {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}

	//拼接
	//Join 用于将元素类型为 string 的 slice, 使用分割符号来拼接组成一个字符串：
	var str08 string = strings.Join(fieldsSlece, ",")
	fmt.Println("Join拼接结果=" + str08) //hello,it's,a,nice,day,today

	fmt.Println("------------对比字符串拼接效率----------------")
	var buffer bytes.Buffer

	start := time.Now()
	for i := 0; i < 100000; i++ {
		buffer.WriteString("test is here\n")
	}
	buffer.String() // 拼接结果
	end := time.Now()
	fmt.Println("Buffer time is ", end.Sub(start).Seconds())

	start = time.Now()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "test is here\n"
	}
	end = time.Now()
	fmt.Println("+= time is ", end.Sub(start).Seconds())

	start = time.Now()
	var sl []string
	for i := 0; i < 100000; i++ {
		sl = append(sl, "test is here\n")
	}
	strings.Join(sl, "")
	end = time.Now()
	fmt.Println("Join time is", end.Sub(start).Seconds())
	/*
		Buffer time is  0.00388283
		+= time is  11.730007558
		Join time is 0.016644653
	*/

}