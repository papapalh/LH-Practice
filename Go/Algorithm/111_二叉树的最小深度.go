package main

import (
	"fmt"
)
/**
 * 题目
 *    给定一个二叉树，找出其最小深度。
 *    最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 * 说明: 叶子节点是指没有子节点的节点。
 * 示例:
 *    给定二叉树 [3,9,20,null,null,15,7],
 *        3
 *       / \
 *      9  20
 *        /  \
 *       15   7
 *     返回它的最小深度  2.
 */
func main() {
	p := TreeNode{Val:3}
	p.Left = &TreeNode{Val:9}

	p.Right = &TreeNode{Val:20}
	p.Right.Left = &TreeNode{Val:15}
	p.Right.Right = &TreeNode{Val:7}

	c := minDepth(&p)

	fmt.Printf("-", c)
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
 * 思路
 *    深度有点递归
 * 耗时
 *    执行用时 :12 ms, 在所有Go提交中击败了93.07%的用户
 *    内存消耗 :5.2 MB, 在所有Go提交中击败了73.74%的用户
 */
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	min := 20000

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil {

		min = minInt(minDepth(root.Left), min)
	}

	if root.Right != nil {
		min = minInt(minDepth(root.Right), min)
	}
	
	return min + 1
}

func minInt (x int, y int) int {
	if x > y {
		return y
	}

	return x
}