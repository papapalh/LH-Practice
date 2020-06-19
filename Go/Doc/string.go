package main

import (
	"fmt"
	"strings"
)

//strings包实现了用于操作字符的简单函数。
//byte = uint8
//rune = int32
func main() {

	/*
	 * 字符串基本方法
	 */
	//字符串是否相同(utf8)(忽略大小写)
	fmt.Println(strings.EqualFold("x", "X"))

	//字符串全小写(不修改原字符串)
	fmt.Println(strings.ToLower("xxxX"))

	//字符串全大写(不修改原字符串)
	fmt.Println(strings.ToUpper("xxxX"))

	/*
	 * 字符串查找
	 */
	//判断前缀字符串prefix
	fmt.Println(strings.HasPrefix("prefix-", "prefix"))

	//判断后缀字符串prefix
	fmt.Println(strings.HasSuffix("-prefix", "prefix"))

	//判断子串substr
	fmt.Println(strings.Contains("-prefix-", "prefix"))

	//查找子串个数sep
	fmt.Println(strings.Count("xxxX", "x"))

	//子串sep第一次出现位置, 不存在返回-1
	fmt.Println(strings.Index("xsepx", "sep"))

	//子串sep最后一次出现位置, 不存在返回-1
	fmt.Println(strings.LastIndex("xsepx", "sep"))

	/*
	 * 字符串操作
	 */
	//串联s个count字符串
	fmt.Println(strings.Repeat("Repeat-", 3))

	//去除前后空格
	fmt.Println(strings.TrimSpace(" TrimSpace "))

	//去除前后cutset
	fmt.Println(strings.Trim("cutsetheiheiheicutset", "cutset"))

	//去除前cutset
	fmt.Println(strings.TrimLeft("cutsetheiheiheicutset", "cutset"))

	//去除后cutset
	fmt.Println(strings.TrimRight("cutsetheiheiheicutset", "cutset"))

	//切割字符串
	fmt.Println(strings.Split("x-x-x-x", "-"))

	//组合字符串
	fmt.Println(strings.Join([]string{"a", "b", "a", "b"}, "-"))

	/*
	 * 字符串替换
	 */
	//字符串替换 Replace(s,old,new,n) 将s中前n个old替换为new,n<0会替换所有old
	fmt.Println(strings.Replace("ssssssss", "s", "x", -1))

	/*
	 * 字符串Reader
	 */
	//构建Reader
	stringRander := strings.NewReader("123456789")

	//未被读取的部分
	fmt.Println(stringRander.Len())

	//读取n个字节数据
	n := make([]byte, 1)
	stringRander.Read(n)
	fmt.Println(string(n))

	//读取Byte字节数据
	fmt.Println(stringRander.ReadByte())

	//读取Rune字节数据(4字节)
	fmt.Println(stringRander.ReadRune())

	//实现io.Seeker接口
	// fmt.Println(stringRander.Seek())

	//实现io.ReadAt
	// fmt.Println(stringRander.ReadAt())

	//实现io.WriteTo接口
	// fmt.Println(stringRander.WriteTo())
}
