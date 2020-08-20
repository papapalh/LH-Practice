package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

var m map[string]bool = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
	"A": true,
	"E": true,
	"I": true,
	"O": true,
	"U": true,
}

/**
 * 方法1
 *     第一次构建字符串，找到元音出现位置，记录位置
 *     第二次循环交换处理数组
 * 耗时
 *    执行用时：24 ms, 在所有 Go 提交中击败了5.26%的用户
 *    内存消耗：9.1 MB, 在所有 Go 提交中击败了6.67%的用户
 */
func reverseVowels1(s string) string {

	type r struct {
		no int
		v  string
	}

	res := []r{}

	sArray := []string{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok {
			res = append(res, r{i, string(s[i])})
		}

		sArray = append(sArray, string(s[i]))
	}

	min := 0
	max := len(res) - 1

	for min <= max {
		sArray[res[min].no] = res[max].v
		sArray[res[max].no] = res[min].v

		min++
		max--
	}

	return strings.Join(sArray, "")
}

/**
 * 方法2
 *     高低指针，一次处理
 * 耗时
 *    执行用时：4 ms, 在所有 Go 提交中击败了83.83%的用户
 *    内存消耗：4.1 MB, 在所有 Go 提交中击败了63.33%的用户
 */
func reverseVowels(s string) string {

	//高低指针
	min := 0
	max := len(s) - 1

	//处理byte
	v := []byte(s)

	for true {
		for min < max {
			if _, ok := m[string(v[min])]; ok {
				break
			}
			min++
		}
		for min < max {
			if _, ok := m[string(v[max])]; ok {
				break
			}
			max--
		}
		if min >= max {
			break
		}
		v[min], v[max] = v[max], v[min]
		min++
		max--
	}
	return string(v)
}
