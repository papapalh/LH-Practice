package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
	 * Response 相关
	 * 表示一个http请求的回复
	 */
	response := http.Response{}

	//获取cookie信息
	response.Cookies()
	fmt.Println(response)

	//读取完response.body记得关闭!
}
