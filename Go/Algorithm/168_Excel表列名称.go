package main

import "fmt"

// 思路
//    建立MAP，根据尾数循环创建字符串，和题解的进制的取值是一个意思
// 耗时
//    执行用时 :0 ms, 在所有 golang 提交中击败了100.00%的用户
//    内存消耗 :1.9 MB, 在所有 golang 提交中击败了93.75%的用户
func convertToTitle(n int) string {

	dict := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z"}

	outPut := ""

	for n > 0 {
		n--
		outPut = dict[n%26] + outPut
		n /= 26
	}
	return outPut
}

func main() {
	i := 28
	fmt.Print(convertToTitle(i))

	s := "hahaha哈哈哈"
	for _, v := range s {
		fmt.Print(v)
		fmt.Print("     ")
	}
}
