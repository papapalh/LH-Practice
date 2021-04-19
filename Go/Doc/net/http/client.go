package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main() {

	request,_ := http.NewRequest(http.MethodGet, "http://php.com/index.php?a=1", nil)

	//HTTP客户端
	client := http.Client{}

	//Do方法根据构建的Request发起请求
	client.Do(request)

	//发送head请求
	client.Head("http://php.com/index.php")

	//发送get请求
	client.Get("http://php.com/index.php")

	//发送POST请求
	client.Post("http://php.com/index.php", "application/x-www-form-urlencoded", strings.NewReader(""))

	//发送POST请求(content-type=application/x-www-form-urlencoded)
	client.PostForm("http://php.com/index.php", url.Values{})



}
