package main

import (
	"fmt"
)

// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
// 示例 1：
//     输入: s = "leetcode"
//     输出: false
// 示例 2：
//     输入: s = "abc"
//     输出: true
// 限制：
//     0 <= len(s) <= 100
// 如果你不使用额外的数据结构，会很加分。
func main() {
	fmt.Println(isUnique("abc"))
}

// 思路(位处理)(大神思路)
//    不用额外的数据结构记录，使用一个数字bit记录a-z出现的位置，出现则在对应位置标1
// 耗时
//    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//    内存消耗：1.9 MB, 在所有 Go 提交中击败了95.45%的用户
func isUnique(astr string) bool {
	mark := 0

	//abc.....的字节编码是 96.97.98......
	//映射到对应的bit上
	for i := 0; i < len(astr); i++ {

		bit := int(astr[i]) - 97

		if mark&(1<<bit) != 0 {
			return false
		}

		mark |= 1 << bit
	}

	return true
}
