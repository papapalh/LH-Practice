package main

import (
	"net/http"
	"strings"
)

func main() {

	/*
	 * Http
	 */

	//默认客户端发发起GET请求
	http.Get("http://php.com/index.php")

	//默认客户端发发起HEAD请求
	http.Head("http://php.com/index.php")

	//默认客户端发发起post请求
	http.Post("http://php.com/index.php", "application/x-www-form-urlencoded", strings.NewReader(""))
}
