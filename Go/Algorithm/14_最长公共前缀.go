package main

import (
	"fmt"
	"strings"
)

// 题目
//     编写一个函数来查找字符串数组中的最长公共前缀。
//     如果不存在公共前缀，返回空字符串 ""。
// 用例
//      输入: ["flower","flow","flight"]
//      输出: "fl"
//      输入: ["dog","racecar","car"]
//      输出: ""
//      解释: 输入不存在公共前缀。
func main() {

	a := []string{"aa", "a"}

	b := longestCommonPrefix(a)

	fmt.Println("x", b)

}
// 水平扫描法
//     遍历 S1-Sn 中所有字符，找到最长的前缀，如果不满足，则算法退出
// 思路
//     强类型啊~,好不适应啊！！！
//     前缀必须是每个单词都有的，所以取出第一个单词作为标准
//     遍历剩下的数组元素作为比较
// 耗时
//     执行用时 : 4 ms, 在Longest Common Prefix的Go提交中击败了91.13% 的用户
//     内存消耗 : 2.5 MB, 在Longest Common Prefix的Go提交中击败了32.53% 的用户
// 时间复杂度
//     时间复杂度：O(S)，S 是所有字符串中字符数量的总和。
func longestCommonPrefix(strs []string) string {

	var comStr string
	var falgStr string

	if len := len(strs); len <= 0 {
		return comStr
	}

	falgStr = string(strs[0])
	l := len(falgStr)

	for _, str := range strs {
		l = Min(l, len(str))
	}

	for i := 0; i < l;i++  {


		for _, str := range strs {

			s := string(str)

			if falgStr[i] != s[i] {
				return comStr
			}

		}
		comStr = comStr + string(falgStr[i])
	}

	return comStr
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//思路
//	纵向扫描法
//	纵向的比较每一个字符串，找到公共前缀
//耗时
//	执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//	内存消耗： 2.3 MB , 在所有 Go 提交中击败了 48.89% 的用户
func longestCommonPrefix2(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	res := strings.Builder{}

	for index, value := range []byte(strs[0]) {

		for i := 1; i < len(strs); i++ {

			if index > len(strs[i]) - 1 {
				goto skip
			}

			if strs[i][index] != value {
				goto skip
			}
		}

		res.Write([]byte{value})
	}

skip:

	return res.String()
}