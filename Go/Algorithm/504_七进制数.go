package main

import (
	"fmt"
	"strconv"
)

//题目
//	给定一个整数，将其转化为7进制，并以字符串形式输出。
//示例
//	输入: 100
//	输出: "202"
//示例 2:
//	输入: -7
//	输出: "-10"
func main() {
	fmt.Println(convertToBase7(0))
}


//思路
//	取余法获取进制
//耗时
//	执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//	内存消耗： 2 MB , 在所有 Go 提交中击败了 73.03% 的用户
func convertToBase7(num int) string {

	if num == 0 {
		return "0"
	}

	res := ""

	n := num

	if num < 0 {
		n *= -1
	}

	for n > 0 {

		res = strconv.Itoa(n%7) + res

		n = n/7
	}

	if num < 0 {
		res = "-" + res
	}

	return res
}