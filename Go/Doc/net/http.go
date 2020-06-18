package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
	 * Head 相关
	 */
	//规范化 head 例如 :"accept-encoding"规范化为"Accept-Encoding"。
	fmt.Println(http.CanonicalHeaderKey("accept-encoding"))

	//确定数据的Content-Type。函数总是返回一个合法的MIME类型；如果它不能确定数据的类型，将返回"application/octet-stream"。
	fmt.Println(http.DetectContentType([]byte("{\"a\":1}")))

	//获取HTTP版本号
	fmt.Println(http.ParseHTTPVersion("HTTP/1.1"))

	//解析head.Date
	fmt.Println(http.ParseTime("Tue, 16 Jun 2020 06:32:13 GMT"))

	//解析状态码
	fmt.Println(http.StatusText(200))

	//Http Head 结构
	header := http.Header{}
	//设置 Head
	header.Set("1", "1")
	//读取head,返回匹配的第一个值
	fmt.Println(header.Get("1"))
	//追加head，如果重复则追加
	header.Add("1", "1")
	//删除head
	header.Del("1")
	fmt.Println(header)

	/*
	 * Cookie 相关
	 */
	cookie := http.Cookie{}
	//返回该cookie的序列化结果
	cookie.String()
	// http.SetCookie()
	//设置cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})

	/*
	 * Request 相关
	 */
	//构建Request结构体
	request, _ := http.NewRequest("GET", "http://127.0.0.1:8000/", nil)
	//检查http协议版本
	fmt.Println(request.ProtoAtLeast(1, 1))
	//获取User-Agent头
	fmt.Println(request.UserAgent())
	//获取Referer头
	fmt.Println(request.Referer())
	//获取Cookies
	fmt.Println(request.Cookies())
	//获取指定Cookie
	fmt.Println(request.Cookie("1"))
	//解析URL中的查询字符串 -> 至r.Form
	request.ParseForm()
	//获取r.From的指定值
	request.FormValue("form")

	/*
	 * Response 相关
	 */
	response := http.Response{}
	fmt.Println(response)

	/*
	 * Client 相关
	 */
	client := http.Client{}
	//发送请求，一般应用GET/POST代替
	fmt.Println(client.Do(request))
	//发送HEAD请求
	h, _ := client.Head("http://www.baidu.com")
	fmt.Println(h.Header)
	//发送GET请求
	fmt.Println(client.Get("http://www.baidu.com"))
	//发送POST请求
	//	bodyType为POST数据的类型
	//  body为POST数据，作为请求的主体。如果参数body实现了io.Closer接口，它会在发送请求后被关闭。调用者有责任在读取完返回值resp的主体后关闭它。
	fmt.Println(client.Post("http://www.baidu.com", "json", nil))
	//发送POST请求（application/x-www-form-urlencoded）
	// fmt.Println(client.PostForm())

	/*
	 * Server 相关
	 */
	server := http.Server{}
	//设置keep-alive
	server.SetKeepAlivesEnabled(true)
	//监听地址
	server.ListenAndServe()
	//绑定对应函数处理
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
		w.Body()
	})

	/*
	 * 其他 相关
	 */
	//返回指定错误码
	// http.Error()
}
