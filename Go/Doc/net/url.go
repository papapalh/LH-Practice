package main

import (
	"fmt"
	"net/url"
)

func main() {
	uString := "http://bing.com/search?q=dotnet&p=xx"

	/*
	 * url encode/decode
	 */
	//QueryEscape函数对s进行转码使之可以安全的用在URL查询里。
	fmt.Println(url.QueryEscape(uString))

	//QueryUnescape函数用于将QueryEscape转码的字符串还原
	fmt.Println(url.QueryUnescape(uString))

	/*
	 * url 实体
	 */
	//实例化 URL 结构(Parse/ParseRequestURI都可以)
	urlParse, _ := url.Parse(uString)

	//解析URL各部分
	fmt.Println(fmt.Sprintf("Scheme: %s \n Host: %s \n Path: %s \n RawQuery: %s \n 是否为绝对URI: %t \n 参数解析: %+v \n 拼接完整请求URI: %s \n 完整URI: %s",
		urlParse.Scheme,
		urlParse.Host,
		urlParse.Path,
		urlParse.RawQuery,
		urlParse.IsAbs(),      //是否为绝对URI(平常用的就是绝对URI，../../这种是相对的)
		urlParse.Query(),      //参数解析
		urlParse.RequestURI(), //拼接请求URI
		urlParse.String(),     //构建完成URI
	))

	/*
	 * 构建请求参数
	 */
	//初始化 url.Value 结构体
	urlValue, _ := url.ParseQuery(uString)

	//设置参数(会覆盖同名参数)
	urlValue.Set("q", "Ava")

	//设置参数(叠加参数)
	urlValue.Add("q", "Jess")

	//删除参数
	urlValue.Del("p")

	//url_encode
	fmt.Println(urlValue.Encode())
}
