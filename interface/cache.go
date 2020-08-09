package main

import (
	"fmt"
	"sync"
)


type version struct {
	id         int64
	effectTime int64 //生效时间
	mu         sync.RWMutex
	data       map[string]interface{}
}

func newVersion() *version {
	return &version{
		id:         0,
		effectTime: 0,
		data:       make(map[string]interface{}),
	}
}

//设置数据
func (v *version) Set(key string, val interface{}) {
	v.mu.Lock()
	v.data[key] = val
	v.mu.Unlock()
}

//获取数据
func (v *version) Get(key string) interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.data[key]
}

//长度
func (v *version) Len() int {
	return len(v.data)
}

//Key是否存在
func (v *version) Exists(key string) bool {
	_, ok := v.data[key]
	return ok
}



func main() {
	d := newVersion()
	d.Set("test","aaa")
	fmt.Println(d)
}
