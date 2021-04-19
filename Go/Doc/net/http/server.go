package main

import "net/http"

type myHandler struct{

}

func (this myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func main() {

	/*
	 * Server 相关
	 */
	server := http.Server{
		Addr: ":8080",
		Handler: myHandler{},
	}

	//是否启用keep-alive
	server.SetKeepAlivesEnabled(true)

	//监听地址
	server.ListenAndServe()


	/*
	 * 其他 相关
	 */
	//返回指定错误码
	// http.Error()
}
