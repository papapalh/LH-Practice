package main

import "fmt"

func main() {
	fmt.Println(checkRecord("LALL"))
}

/**
 * 题目
 *    给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符：
 *    'A' : Absent，缺勤
 *    'L' : Late，迟到
 *    'P' : Present，到场
 *    如果一个学生的出勤记录中不超过一个'A'(缺勤)并且不超过两个连续的'L'(迟到),那么这个学生会被奖赏。
 * 示例
 *    输入: "PPALLP"
 *    输出: True
 * 示例 2:
 *    输入: "PPALLL"
 *    输出: False
 * 思路
 *    就循环
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：2 MB, 在所有 Go 提交中击败了74.07%的用户
 */
func checkRecord(s string) bool {
	a := 0
	l := 0

	for _, v := range s {
		if string(v) == "A" {
			a++
			l = 0
			if a > 1 {
				return false
			}
		} else if string(v) == "L" {
			l++
			if l > 2 {
				fmt.Println(l)
				return false
			}
		} else {
			l = 0
		}
	}

	return true
}
