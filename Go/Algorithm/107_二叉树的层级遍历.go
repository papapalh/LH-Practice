package main

import (
	"fmt"
)
/**
 * 题目
 *    给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 * 例如：
 *    给定二叉树 [3,9,20,null,null,15,7],
 *       3
 *      / \
 *     9  20
 *       /  \
 *      15   7
 * 返回其自底向上的层次遍历为：
 * [
 *    [15,7],
 *    [9,20],
 *    [3]
 * ]
 */
func main() {
	p := TreeNode{Val:3}
	p.Left = &TreeNode{Val:9}

	p.Right = &TreeNode{Val:20}
	p.Right.Left = &TreeNode{Val:15}
	p.Right.Right = &TreeNode{Val:7}

	c := levelOrderBottom(&p)

	fmt.Printf("-", c)
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
 * 思路
 *     层级遍历，数组翻转
 * 耗时
 *    执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
 *    内存消耗 :6.2 MB, 在所有Go提交中击败了45.98%的用户
 */
var tree = [][]int{}
func levelOrderBottom(root *TreeNode) [][]int {

	tree = [][]int{}

	if root == nil {
		return tree
	}
	x(root, 0)

	return reverse(tree)
}

func x (root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if depth >= len(tree) {
		tree = append(tree, []int{})
	}

	tree[depth] = append(tree[depth], root.Val)

	depth++

	x(root.Left, depth)
	x(root.Right, depth)
}

func reverse(trees [][]int) [][]int {
	n := len(trees)
	reversed := make([][]int, 0, n)
	for i:=n-1;i>=0;i-- {
		reversed = append(reversed, trees[i])
	}
	return reversed
}