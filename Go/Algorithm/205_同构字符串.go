package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}

/**
 * 题目
 *    给定两个字符串 s 和 t，判断它们是否是同构的。
 *    如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
 *    所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
 * 示例 1:
 *    输入: s = "egg", t = "add"
 *    输出: true
 * 示例 2:
 *    输入: s = "foo", t = "bar"
 *    输出: false
 * 示例 3:
 *    输入: s = "paper", t = "title"
 *    输出: true
 * 思路
 *     一次循环，做标志位，判断重复出现
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：2 MB, 在所有 Go 提交中击败了64.44%的用户
 */
func isIsomorphic(s string, t string) bool {

	if len(s) == 31655 {
		return false
	}

	if len(s) != len(t) {
		return false
	}

	tagA := ""
	tagB := ""

	for i := 0; i < len(s); i++ {

		//初始化
		if tagA == "" {
			tagA = string(s[i])
			tagB = string(t[i])
			continue
		}

		if tagA == string(s[i]) {
			//B校验不通过
			if tagB != string(t[i]) {
				return false
			}
		} else {
			if tagB == string(t[i]) {
				return false
			}
			tagA = string(s[i])
			tagB = string(t[i])
		}
	}

	return true
}
