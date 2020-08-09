package main

import(
	"fmt"
	"net/http"
	"sync"
)


// HandlerFunc defines the handler used by  middleware as return value.
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

//定义的上下文
type Context struct {
	Request   *http.Request
	Writer    http.ResponseWriter
	handlers HandlersChain
	index    int8

}

//模拟的调用堆栈
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		//按顺序执行HandlersChain内的函数
		//如果函数内无c.Next()方法调用则函数顺序执行完
		//如果函数内有c.Next()方法调用则代码执行到c.Next()方法处压栈，等待后面的函数执行完在回来执行c.Next()后的命令
		c.handlers[c.index](c)
		c.index++
	}
}
func (c *Context) reset() {
	c.handlers = nil
	c.index = -1
}

//中间件组
type RouterGroup struct {
	//存储定义的中间件
	Handlers HandlersChain
	engine   *Engine
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

func (group *RouterGroup) AddRoute(absolutePath string, handlers ...HandlerFunc) {
	handlers = group.combineHandlers(handlers)
	//建立路由和相关中间件组的绑定
	group.engine.addRoute(absolutePath, handlers)
}

//将定义的公用中间件和路由相关的中间件合并
func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

type Engine struct{

	tree map[string]HandlersChain	// tree为了简化做成了map路由路径完全匹配
	RouterGroup
	pool             sync.Pool		// 正常情况存在大量的上下文切换，所以使用一个临时对象存储
}

func NewEngine() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
		},
		tree: make(map[string]HandlersChain),
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (engine *Engine) allocateContext() *Context {
	return &Context{}
}

//url请求时，默认执行入口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.Writer = w
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	rPath := c.Request.URL.Path

	handlers := engine.getValue(rPath)
	if handlers != nil {
		c.handlers = handlers
		//按顺序执行中间件
		c.Next()
		return
	}

}

//获取路由下的相关HandlersChain
func (engine *Engine)getValue(path string)(handlers HandlersChain){
	handlers,ok := engine.tree[path]
	if !ok {
		return nil
	}
	return
}

func (engine *Engine) addRoute(path string, handlers HandlersChain) {
	engine.tree[path]=handlers
}

func (engine *Engine) Use(middleware ...HandlerFunc)  {
	engine.RouterGroup.Use(middleware...)
}

func main(){
	engine := NewEngine()
	engine.Use(func(c * Context){
		fmt.Println("begin middle1")
		fmt.Println("end middle1")
	})
	engine.Use(func(c * Context){
		fmt.Println("begin middle2")
		c.Next()
		fmt.Println("end middle2")
	})
	engine.Use(func(c * Context){
		fmt.Println("begin middle3")
		c.Next()
		fmt.Println("end middle3")
	})
	engine.AddRoute("/path1",func( c *Context){
		fmt.Println("path1")
		c.Writer.Write([]byte("path1"))

	})
	engine.AddRoute("/path2",func( c *Context){
		fmt.Println("path2")
		c.Writer.Write([]byte("path2"))

	})
	http.ListenAndServe(":8080", engine)
}