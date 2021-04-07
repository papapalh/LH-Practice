package main

import (
	"fmt"
	"math"
	"strings"

	//"time"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//	输入: s = "abcabcbb"
//	输出: 3
//	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//	输入: s = "bbbbb"
//	输出: 1
//	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//	输入: s = "pwwkew"
//	输出: 3
//	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//示例 4:
//	输入: s = ""
//	输出: 0
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

//思路
//	滑动窗口
//	https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/zhi-zeng-da-bu-jian-xiao-de-hua-dong-chuang-kou-10/
//耗时
//	执行用时： 348 ms , 在所有 Go 提交中击败了 6.53% 的用户
//	内存消耗： 6.9 MB , 在所有 Go 提交中击败了 5.03% 的用户
func lengthOfLongestSubstring(s string) int {

	start, end := 0, 0

	for i := 0; i < len(s); i++ {

		index := strings.Index(s[start:i], string(s[i]))

		if index == -1 {

			if i + 1 > end {
				end = i + 1
			}
			
		} else {
			start += index + 1
			end += index + 1
		}
	}

	return end - start
}

func lengthOfLongestSubstring2(s string) int {

	var max float64 = 0

	dict := map[byte]struct{}{}

	for i := 0; i < len(s); i++ {

		dict[s[i]] = struct{}{}

		for j := i + 1; j < len(s); j++ {

			max = math.Max(max, float64(j - i))

			if _, ok := dict[s[j]]; ok {
				dict = map[byte]struct{}{}
				break
			}

			dict[s[j]] = struct{}{}
		}

		if len(dict) > 0 {
			max = math.Max(max, float64(len(s) - i))
		}
	}

	return int(max)
}
