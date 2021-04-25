package main

import "net/http"

type myHandler struct{

}

func (this myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func main() {

	/*
	 * http 路由器 - NewServeMux
	 * NewServeMux
	 * 匹配模式:
	 *    1.长模式优先短模式 - 因此如果模式"/images/"和"/images/thumbnails/"都注册了处理器，后一个处理器会用于路径以"/images/thumbnails/"开始的请求
	 */
	c := http.NewServeMux()

	//注册路由(两种方法)
	c.Handle("/456", myHandler{})
	c.HandleFunc("/123", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	//赋给默认路由
	http.DefaultServeMux = c

	/*
	 * Server 相关
	 */
	server := http.Server{
		Addr: ":8080",               // 监听的TCP地址，如果为空字符串会使用":http"
		//Handler:     Handler,      // 调用的处理器，如为nil会调用http.DefaultServeMux
		//ReadTimeout  time.Duration // 请求的读取操作在超时前的最大持续时间
		//WriteTimeout time.Duration // 回复的写入操作在超时前的最大持续时间
	}

	//是否启用keep-alive
	server.SetKeepAlivesEnabled(true)

	//监听地址
	server.ListenAndServe()
}
