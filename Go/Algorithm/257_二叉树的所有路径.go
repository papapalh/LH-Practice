package main

import (
	"fmt"
	"strconv"

	// "strconv"
)
/**
 * 题目
 *     给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *     说明: 叶子节点是指没有子节点的节点。
 * 示例:
 *      输入:
 *          1
 *        /   \
 *       2     3
 *        \
 *        5
 * 输出: ["1->2->5", "1->3"]
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
 */
func main() {
	p := TreeNode{Val:1}
	p.Left = &TreeNode{Val:2}
	p.Left.Right = &TreeNode{Val:5}

	p.Right = &TreeNode{Val:3}
	// p.Right.Left = &TreeNode{Val:5}
	// p.Right.Right = &TreeNode{Val:6}

	c := binaryTreePaths(&p)

	fmt.Printf("-", c)
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
 * 思路
 *     深度优先遍历，输出路径
 * 耗时
 *     执行用时 :4 ms, 在所有 Go 提交中击败了72.53%的用户
 *     内存消耗 :2.5 MB, 在所有 Go 提交中击败了97.83%的用户
 */
var dict []string

func binaryTreePaths(root *TreeNode) []string {

	dict = []string{}

	if root == nil {
		return dict
	}

	x(root, "")

    return dict
}

func x(root *TreeNode, path string)  {

	if root.Left == nil && root.Right == nil {
		dict = append(dict, path + strconv.Itoa(root.Val))
		return
	}

	path = path + strconv.Itoa(root.Val) + "->"

	if root.Left != nil {
		x(root.Left, path)
	}
	if root.Right != nil {
		x(root.Right, path)
	}
}