package main

import (
	"fmt"
)
/**
 * 题目
 *    给定一个二叉树，检查它是否是镜像对称的。
 *    例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *         1
 *        / \
 *       2   2
 *      / \ / \
 *     3  4 4  3
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *         1
 *        / \
 *       2   2
 *        \   \
 *        3    3
 * 说明:
 *     如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
 */
func main() {
	p := TreeNode{Val:1}
	p.Left = &TreeNode{Val:2}
	p.Left.Left = &TreeNode{Val:3}
	p.Left.Right = &TreeNode{Val:4}

	p.Right = &TreeNode{Val:2}
	p.Right.Left = &TreeNode{Val:4}
	p.Right.Right = &TreeNode{Val:3}

	c := isSymmetric(&p)

	fmt.Printf("-", c)
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
 * 思路
 *    递归
 *    通过镜像的概念，对左子树和右子数进行方向不同的递归进行比较
 * 耗时
 *    执行用时 :0 ms, 在所有Go提交中击败了 100.00% 的用户
 *    内存消耗 :3 MB, 在所有Go提交中击败了 52.17% 的用户
 */
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return x(root.Left, root.Right)
}

func x(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 ==nil && node2 ==nil {
		return true
	}
	if (node1==nil && node2!=nil) || (node1!=nil && node2==nil) || (node1.Val!=node2.Val) {
		return false
	}

	return x(node1.Left, node2.Right) && x(node1.Right, node2.Left)
}