// 题目
//     判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//     你能不将整数转为字符串来解决这个问题吗？
// 用例
//     输入: 输入: 121
//     输出: true
//     输入: -121
//     解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//     输出: false
package main

import "fmt"

func main() {
	a := IsPalindrome(121)
	fmt.Println(a)
}
// 思路
//     回文数
//     和 7_整数反转 是一个道理
//     如果一个数是回文数，则 反转之后的数 = 原数字
// 耗时
// 	   执行用时 : 28 ms, 在Palindrome Number的Go提交中击败了95.42% 的用户
//     内存消耗 : 5.1 MB, 在Palindrome Number的Go提交中击败了65.52% 的用户
// 时间复杂度
//	   O（n）
func IsPalindrome(x int) bool {

	// 负数都不可能是回文数
	if x < 0 {
		return false
	}

	// 数字
	dmp    := x
	result := 0

	for ; dmp >= 1;  {
		result = (result * 10) + (dmp % 10)

		dmp = dmp / 10
	}

	if result == x {
		return true
	}


	return false
}