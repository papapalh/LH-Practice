package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
 	 * Cookie 相关
 	*/
	cookie := http.Cookie{
		Name:"cookie",
		Value:"123",
	}
	fmt.Printf("%+v", cookie)
	//返回该cookie的序列化结果
	//cookie.String()
	// http.SetCookie()
	//设置cookie
	//http.SetCookie(w, &http.Cookie{
	//	Name:  "my-cookie",
	//	Value: "some value",
	//})
}
