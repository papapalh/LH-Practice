package main

import "fmt"

//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
// 示例 1：
// 输入: s1 = "abc", s2 = "bca"
// 输出: true
// 示例 2：
// 输入: s1 = "abc", s2 = "bad"
// 输出: false
// 说明：
// 0 <= len(s1) <= 100
// 0 <= len(s2) <= 100
func main() {
	fmt.Println(CheckPermutation("abc", "cba"))
}

//思路
//    双哈希表去重
//    其实我挺想做做递归的问题，一会可以试下
//耗时
//    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//    内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := map[string]int{}

	for _, v := range s1 {
		m1[string(v)] += 1
	}

	for _, v := range s2 {

		x, ok := m1[string(v)]
		if !ok {
			return false
		}

		if x <= 0 {
			return false
		}

		m1[string(v)] -= 1
	}

	return true
}
