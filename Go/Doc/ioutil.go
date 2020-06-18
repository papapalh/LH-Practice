package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

//io工具函数
func main() {
	reader := bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8})

	//读取全部字节
	fmt.Println(ioutil.ReadAll(reader))

	//读取文件全部内容[底层为os+readAll]
	fmt.Println(ioutil.ReadFile("/Users/lihong/Desktop/on-line/Practice/Go/Doc/io.go"))

	//写文件内容
	//	文件不存在将按给出的权限创建文件
	//	文件存在则清空文件并重新写入
	fmt.Println(ioutil.WriteFile("/Users/lihong/Desktop/on-line/Practice/Go/Doc/ioutil-WriteFile.go", []byte("hello word"), 0777))

	//读取目录列表
	fmt.Println(ioutil.ReadDir("/Users/lihong/Desktop/on-line/Practice/Go/Doc/"))

	//创建临时文件夹[以prefix为前缀]
	fmt.Println(ioutil.TempDir("/Users/lihong/Desktop/on-line/Practice/Go/Doc/", "prefix"))
}
