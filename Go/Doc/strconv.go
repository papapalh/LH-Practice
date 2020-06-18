package main

import (
	"fmt"
	"strconv"
)

//strconv包实现了基本数据类型和其字符串表示的相互转换。
func main() {

	//布尔->字符串
	fmt.Println(strconv.FormatBool(true))

	/*
	 * 数字 -> 字符串 相关
	 */
	//数字(进制)->字符串
	fmt.Println(strconv.FormatInt(54, 10))

	//数字(无符号)(进制)->字符串
	fmt.Println(strconv.FormatUint(54, 10))

	//数字(十进制)->字符串
	fmt.Println(strconv.Itoa(54))

	/*
	 * 字符串 -> 数字
	 */
	//字符串->int : base进制(base2-36)bitSize类型(0->int,8->int8...)
	fmt.Println(strconv.ParseInt("54", 10, 0))

	//字符串->uint : base进制(base2-36)bitSize类型(0->int,8->int8...)
	fmt.Println(strconv.ParseUint("54", 10, 0))

	//字符串->float : bitSize类型(32->float32，64->float64)
	fmt.Println(strconv.ParseFloat("54", 32))

	//字符换->int : Atoi是ParseInt(s, 10, 0)的简写
	fmt.Println(strconv.Atoi("54"))
}
