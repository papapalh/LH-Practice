package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	/*
	 * Request 相关
	 * 	  构建请求包
	 */
	//构建Request结构体
	request, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/", strings.NewReader(""))

	//设置
	//设置cookie头(设置重复的cookie不会覆盖)
	request.AddCookie(&http.Cookie{})
	//设置basicAuth认证
	request.SetBasicAuth("user", "1234")
	//设置header头
	request.Header = http.Header{}
	//设置multipart包数据
	request.MultipartForm = &multipart.Form{}
	//设置Form表单数据，包括URL字段的query参数和POST或PUT的表单数据(本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代)
	request.Form = url.Values{}
	//设置PUT/POST表单数据(本字段只有在调用ParseMultipartForm后才有效,在客户端，会忽略请求中的本字段而使用Body替代)
	request.PostForm = url.Values{}
	//设置MultipartForm是解析好的多部件表单，包括上传的文件(本字段只有在调用ParseMultipartForm后才有效,在客户端，会忽略请求中的本字段而使用Body替代)
	request.PostForm = url.Values{}

	//功能
	//解析URL查询字符串,并将解析结果更新到r.Form字段
	request.ParseForm()
	//解析multipart/form-data解析
	request.ParseMultipartForm(1024)

	//查看
	//获取User-Agent头
	fmt.Println(request.UserAgent())
	//获取Referer头
	fmt.Println(request.Referer())
	//获取全部Cookies
	fmt.Println(request.Cookies())
	//获取指定Cookie
	fmt.Println(request.Cookie("1"))

	//检查http协议版本
	fmt.Println(request.ProtoAtLeast(1, 1))
	//解析URL中的查询字符串 -> 至r.Form
	request.ParseForm()
	//获取r.From的指定值
	request.FormValue("form")

	//使用完记得关闭！
	request.Body.Close()
}
