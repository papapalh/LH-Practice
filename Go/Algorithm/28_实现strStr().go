package main

import "fmt"
/**
 * 题目
 *     实现 strStr() 函数。
 *     给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
 * 示例
 *    输入: haystack = "hello", needle = "ll"
 *    输出: 2
 *    输入: haystack = "aaaaa", needle = "bba"
 *    输出: -1
 */
func main () {

	c := strStr("mississippi","issipi")

	fmt.Printf("x", c)
}
/**
 * 思路1
 *     暴力破解
 * 耗时
 *     执行用时 : 0 ms, 在Implement strStr()的Go提交中击败了100.00% 的用户
 *     内存消耗 : 2.3 MB, 在Implement strStr()的Go提交中击败了53.10% 的用户
 * 时间复杂度
 *     O（m * n）
 */
func strStr(haystack string, needle string) int {

	// 边界判断
	if haystack == ""  && needle == "" {
		return 0
	}
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {

			if len(needle) > len(haystack[i:]) {
				goto skip
			}

			for j := 0; j < len(needle); j++ {
				if needle[j] != haystack[i+j] {
					goto skip
				}
			}
			return i
		}
	skip:
		continue
	}

	return -1
}
/**
 * 思路2
 *     KMP 算法
 *     https://www.cnblogs.com/yjiyjige/p/3263858.html
 */

