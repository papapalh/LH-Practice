package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 题目
 *    判断一个二叉树是否是平衡的
 *    https://leetcode-cn.com/problems/balanced-binary-tree/
 * 方法
 *    1.暴力求解(从顶至底)
 *        获取一颗树全部的路径深度判断，在判断是否平衡。
 *    2.从底至顶
 *        从底部开始递归，遇到不满足的就退出
 *        https://leetcode-cn.com/problems/balanced-binary-tree/solution/balanced-binary-tree-di-gui-fang-fa-by-jin40789108/
 * 耗时
 *    执行用时：8 ms, 在所有 Go 提交中击败了90.30%的用户
 *    内存消耗：5.7 MB, 在所有 Go 提交中击败了59.46%的用户
 */
func main() {
	p := TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Left.Left = &TreeNode{Val: 4}
	p.Left.Left.Left = &TreeNode{Val: 44}
	// p.Left.Right = &TreeNode{Val: 5}

	p.Right = &TreeNode{Val: 3}
	p.Right.Left = &TreeNode{Val: 6}
	p.Right.Right = &TreeNode{Val: 7}

	fmt.Println(isBalanced(&p))
}

var a []int

func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}
func recur(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := recur(root.Left)
	if left == -1 {
		return -1
	}

	right := recur(root.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) < 2 {
		return int(math.Max(float64(left), float64(right))) + 1
	}

	return -1
}
