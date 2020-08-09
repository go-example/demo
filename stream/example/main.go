package main

import (
	"fmt"
	"github.com/go-example/demo/stream"
)

//起床
func GetUP(arg interface{}) (interface{}, error) {
	t, _ := arg.(string)
	fmt.Println("铃铃.......", t, "###到时间啦，再不起又要迟到了！")
	return "醒着的状态", nil
}

//蹲坑
func GetPit(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###每早必做的功课，蹲坑！")
	return "舒服啦", nil
}

//洗脸
func GetFace(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###洗脸很重要！")
	return "脸已经洗干净了，可以去见人了", nil
}

//刷牙
func GetTooth(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###刷牙也很重要！")
	return "牙也刷干净了，可以放心的大笑", nil
}

//吃早饭
func GetEat(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###吃饭是必须的(需求变更了，原来的流程里没有，这次加上)")
	return "吃饱饱了", nil
}

//换衣服
func GetCloth(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###还要增加一个换衣服的流程！")
	return "找到心仪的衣服了", nil
}

//出门
func GetOut(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###一切就绪，可以出门啦！")
	return "", nil

}
func main() {
	stream.NewStream().
		Next(GetUP).
		Next(GetPit).
		Next(GetTooth).
		Next(GetFace).
		Next(GetEat). //需求变更了后加上的
		Next(GetCloth).
		Next(GetOut).
		Go("2018年1月28日8点10分")
}
