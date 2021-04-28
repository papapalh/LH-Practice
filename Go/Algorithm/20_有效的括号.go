package main

import "fmt"
// 题目
//     给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//     有效字符串需满足：
//     左括号必须用相同类型的右括号闭合。
//     左括号必须以正确的顺序闭合。
//     注意空字符串可被认为是有效字符串。
// 用例
//     输入: "()"
//     输出: true
//     输入: "()[]{}"
//     输出: true
//     输入: "(]"
//     输出: false
//     输入: "{[]}"
//     输出: true
func main () {

	b := isValid("()[]{}")
	fmt.Println(b)
}
// 思路
//     栈
//     如果发现是 关闭 符号，则 出栈 获取。
//     如果不同，则退出算法
//     一定注意栈的边界
// 详解
//     1：输入 '[{}]}'
//     2：因为 '[{' 不是关闭符，所以入栈，当前栈结构 ['[', '{']
//     3：当 '}' 进入，发现是关闭符，所以出栈获取字典 '{',相同，所以是符合规则
// 耗时
//     执行用时 : 0 ms, 在Valid Parentheses的Go提交中击败了100.00% 的用户
//     内存消耗 : 2.1 MB, 在Valid Parentheses的Go提交中击败了46.65% 的用户
// 时间复杂度
//     O(n)
func isValid(s string) bool {

	if s == "" || string(s[0]) == ")" ||  string(s[0]) == "]" ||  string(s[0]) == "}" {
		return false
	}

	m := map[string]string{
		"(":")",
		"[":"]",
		"{":"}",
	}

	track := []string{}

	for _, v := range s {

		v_s := string(v)

		if v_s == "(" || v_s == "[" || v_s == "{" {
			track = append(track, m[v_s])
			continue
		}

		if len(track) == 0 {
			return false
		}
		if v_s != track[len(track)-1] {
			return false
		}
		track = track[:len(track)-1]
	}

	return len(track) == 0
}