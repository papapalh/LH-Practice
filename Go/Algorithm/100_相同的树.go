package main

import "fmt"
/**
 * 题目
 *     给定两个二叉树，编写一个函数来检验它们是否相同。
 *     如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 * 示例 1:
 *     输入:       1         1
 *                / \       / \
 *               2   3     2   3
 *              [1,2,3],   [1,2,3]
 *     输出: true
 * 示例 2:
 *     输入:      1          1
 *               /           \
 *              2             2
 *            [1,2],     [1,null,2]
 *     输出: false
 * 示例 3:
 *     输入:       1         1
 *               / \       / \
 *              2   1     1   2
 *             [1,2,1],   [1,1,2]
 *     输出: false
 */
func main() {
	p := TreeNode{Val:1}
	p.Left = &TreeNode{Val:2}
	p.Right = &TreeNode{Val:3}

	q := TreeNode{Val:1}
	q.Left = &TreeNode{Val:2}
	q.Right = &TreeNode{Val:2}

	c := isSameTree(&p, &q)

	fmt.Printf("-", c)
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/*
 * 思路
 *     前序遍历树
 * 耗时
 *     执行用时 : 0 ms, 在Same Tree的Go提交中击败了100.00% 的用户
 *     内存消耗 : 2.1 MB, 在Same Tree的Go提交中击败了79.31% 的用户
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

	// 都到了 nil 结点，说明已经到了叶子结点
	if p == nil && q == nil {
		return true
	}

	// 判断结点是否存在
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	// 判断结点值
	if p.Val != q.Val {
		return false
	}

	// 递归判断
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}