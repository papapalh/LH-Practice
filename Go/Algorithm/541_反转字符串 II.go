package main

import (
	"fmt"
)

//题目
//	给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。
//		如果剩余字符少于 k 个，则将剩余字符全部反转。
//		如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//示例:
//		输入: s = "abcdefg", k = 2
//		输出: "bacdfeg"

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}


//思路
//	k个为一组处理字符串
//耗时
//	未通过用例,我的理解和题目用例不一样
func reverseStr(s string, k int) string {

	res := ""

	//如果剩余字符少于 k 个，则将剩余字符全部反转
	if len(s) < k {
		for ii := len(s) - 1; ii >= 0; ii-- {
			res += string(s[ii])
		}
		return res
	}

	flag := 1
	stack := []byte{}

	for i := 1; i < len(s) + 1; i++ {

		//入栈
		stack = append(stack, s[i-1])

		//如果不到处理节点，则跳过
		if i % k != 0 {
			continue
		}

		//出栈翻转
		if flag == 1 {
			for ii := len(stack) - 1; ii >= 0; ii-- {
				res += string(stack[ii])
			}
			flag = 0
			stack = []byte{}
			continue
		}

		//正常输出
		if flag == 0 {
			for _, s := range stack {
				res += string(s)
			}
			flag = 1
			stack = []byte{}
			continue
		}
	}

	//补充尾巴
	if len(stack) > 0 {
		for _, s := range stack {
			res += string(s)
		}
		stack = []byte{}
	}

	return res
}