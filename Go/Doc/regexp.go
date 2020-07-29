package main

import (
	"regexp"
	"fmt"
)

func main (){
	str := "123456end654321"

	/*
	 * Match 检查是否匹配
	 */

	//string是否匹配
	//	校验-六位连续的数字 -> 输出true
	fmt.Println(regexp.MatchString("\\d{6}", str))

	//Byte是否匹配
	//	校验是否为汉字
	fmt.Println(regexp.Match("[\u4e00-\u9fa5]", []byte("经度")))

	//判断是否含有字符（大小写）
	fmt.Println(regexp.Match("[a-zA-Z]", []byte("av132")))

	//判断是否含有以数字开头，不是为true
	fmt.Println(regexp.Match(`[^\d]`, []byte("as132")))

	//判断是否含有为IP地址
	fmt.Println(regexp.MatchString( "[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]+", "10.32.12.01"))

	/*
	 * 获取匹配字符串
	 */

	//返回第一个匹配的字符串
	reg := regexp.MustCompile("\\d{6}")

	//Byte方式
	data := reg.Find([]byte(str))
	fmt.Println(string(data))  //880218

	//String方式
	data2 := reg.FindString(str)
	fmt.Println(data2)
}