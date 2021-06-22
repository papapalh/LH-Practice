package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

//思路
//	本质还是使用回溯+剪枝进行递归找出所有符合条件的
//  这个人讲的很好:https://leetcode-cn.com/problems/generate-parentheses/solution/sui-ran-bu-shi-zui-xiu-de-dan-zhi-shao-n-0yt3/
//耗时
//	执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//	内存消耗： 2.8 MB , 在所有 Go 提交中击败了 47.07% 的用户
func generateParenthesis(n int) []string {

	res := []string{}

	if n == 0 {
		return res
	}

	dns22(&res, n, "", 0, 0)

	return res
}

//left和right用于记录左右括号数量-用于剪枝
func dns22(res *[]string, n int, node string, left int, right int) {

	//这个判断去除了不合法的括号，左括号一定是多余右括号的
	if right > left {
		return
	}

	if left > n || right > n {
		return
	}

	if len(node) == n * 2 {
		*res = append(*res, node)
		return
	}

	dns22(res, n, node + "(", left+1, right)
	dns22(res, n, node + ")", left, right+1)
}