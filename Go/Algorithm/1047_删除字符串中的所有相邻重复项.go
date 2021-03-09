package main

import (
	"fmt"
	"strings"
)

// 给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
// 在 S 上反复执行重复项删除操作，直到无法继续删除。
// 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。
// 示例：
// 		输入："abbaca"
// 		输出："ca"
// 解释：
// 		例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
func main() {

	a := []int{1, 2, 3}
	fmt.Println(a[:len(a)-1])

	fmt.Println(removeDuplicates("abbaca"))
}

// 思路
// 		栈，删除重复的栈底元素
// 耗时
// 		执行用时：32 ms, 在所有 Go 提交中击败了15.79%的用户
// 		内存消耗：7.6 MB, 在所有 Go 提交中击败了19.74%的用户
func removeDuplicates(S string) string {

	stack := []string{}

	for _, v := range S {

		//当栈为空OR导数两个被删除,则初始化
		if len(stack) == 0 {
			stack = append(stack, string(v))
			continue
		}

		//如果是重复字符,则最后一个出栈
		if string(v) == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, string(v))
	}

	return strings.Join(stack, "")
}
