### 代码
```
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

	http.ListenAndServe("0.0.0.0:8888", nil)
}
```
在浏览器使用 http://localhost:8888/debug/pprof/ 即可访问



### 查看profile ：
1.如果安装过graphviz直接提交过这步骤，否则可以到 http://www.graphviz.org/download/ 下载，并把bin加入到环境变量
windows 下载地址：https://graphviz.gitlab.io/_pages/Download/Download_windows.html

2.在命令行输入 
```
go tool pprof http://localhost:8888/debug/pprof/profile?seconds=30
```
此后的30秒进入收集profile信息的状态。30秒后进入pprof的交互模式，然后输入
```
web
```
然后浏览器自动弹开到网页展示svg图

3.查看已经保存的profile文件
```
go tool pprof profile C:\Users\user\pprof\pprof.samples.cpu.004.pb.gz
```
然后也是进入pprof的交互模式，然后输入web

4.还可以查看heap和goroutine
```
go tool pprof http://localhost:8888/debug/pprof/heap
go tool pprof http://127.0.0.1:8888/debug/pprof/goroutine
```

5.开启本地图型界面
```
go tool pprof -http=:8889 http://localhost:8888/debug/pprof/profile
```
会自动打开http://localhost:8889/ui/