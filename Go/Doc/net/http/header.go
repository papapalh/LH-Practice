package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
	 * Header
	 *    本质是一个MAP，可以通过MAP进行配置
	 *    推荐使用本身的header函数构建
	 */
	h := http.Header{}

	//设置新的header(如果header已经存在值，则被覆盖)
	//	  也可以直接通过map设置(不推荐): h["xx"] = []string{"asdasd"}
	h.Set("l", "30")

	//添加新的header
	//   header不存在,则新建
	//   header存在,向后填充
	h.Add("ll", "30")

	//获取header信息
	fmt.Println(h.Get("l"))

	//删除header信息
	h.Del("ll")

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
}
