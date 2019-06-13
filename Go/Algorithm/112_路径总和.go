package main

import (
	"fmt"
)
/**
 * 题目
 *    给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 *    说明: 叶子节点是指没有子节点的节点。
 * 示例: 
 *    给定如下二叉树，以及目标和 sum = 22，
 *            5
 *           / \
 *          4   8
 *         /   / \
 *        11  13  4
 *       /  \      \
 *      7    2      1
 *     返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
 */
func main() {
	p := TreeNode{Val:5}
	p.Left = &TreeNode{Val:4}
	p.Left.Left = &TreeNode{Val:11}
	p.Left.Left.Left = &TreeNode{Val:7}
	p.Left.Left.Right = &TreeNode{Val:2}

	p.Right = &TreeNode{Val:8}
	p.Right.Left = &TreeNode{Val:13}
	p.Right.Right = &TreeNode{Val:4}
	p.Right.Right.Right = &TreeNode{Val:1}

	c := hasPathSum(&p, 22)

	fmt.Printf("-", c)
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 思路
 *    前序遍历，抓取所有路径，取和判断
 * 耗时
 *    执行用时 :4 ms, 在所有Go提交中击败了99.53%的用户
 *    内存消耗 :4.9 MB, 在所有Go提交中击败了84.00%的用户
 */
func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && sum == 1 {
		return true
	}

	return xx(root, 0, sum)
}

func xx(root *TreeNode, v int, sum int) bool {

	if root.Left == nil && root.Right == nil {

		if (v + root.Val) == sum {
			return true
		}

		return false
	}

	if root.Left != nil {

		if xx(root.Left, v + root.Val, sum) {
			return true
		}
	}

	if root.Right != nil {
		if xx(root.Right, v + root.Val, sum) {
			return true
		}
	}

	return false
}

