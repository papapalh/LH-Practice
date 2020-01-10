package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("strconv包实现了基本数据类型和其字符串表示的相互转换。")
	fmt.Println("    FormatBool(b)              根据b的布尔值返回字符串                                       ", strconv.FormatBool(true))
	fmt.Println("    FormatInt(i,base)          返回i(int)的base进制字符串(base2-36)                          ", strconv.FormatInt(54, 10))
	fmt.Println("    FormatUint(i,base)         返回i(uint)的base进制字符串(base2-36)                         ", strconv.FormatInt(54, 10))
	fmt.Println("    Itoa(i)                    FormatInt(i,10)的简写                                         ", strconv.Itoa(54))
	fmt.Print("    ParseInt(i,base,bitSize)   返回i(int)的base进制(base2-36)bitSize类型(0->int,8->int8...)   ")
	fmt.Println(strconv.ParseInt("54", 10,0))
	fmt.Print("    ParseUint(i,base,bitSize)  返回i(uint)的base进制(base2-36)bitSize类型(0->int,8->int8...)  ")
	fmt.Println(strconv.ParseUint("54", 10,0))
	fmt.Print("    ParseFloat(i,bitSize)      返回i的bitSize类型(32->float32，64->float64)                   ")
	fmt.Println(strconv.ParseFloat("54", 32))
	fmt.Print("    Atoi(i)                    Atoi是ParseInt(s, 10, 0)的简写                                 ")
	fmt.Println(strconv.Atoi("54"))
}