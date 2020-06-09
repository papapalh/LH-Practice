package main

import (
	"fmt"
)

/**
 * 题目
 *    给定一个二叉树，找出其最大深度。
 *    二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 * 说明:
 *    叶子节点是指没有子节点的节点。
 * 示例：
 *    给定二叉树 [3,9,20,null,null,15,7]
 *        3
 *       / \
 *      9  20
 *        /  \
 *       15   7
 *     返回它的最大深度 3 。
 */
func main() {
	p := TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Left.Left = &TreeNode{Val: 3}
	p.Left.Right = &TreeNode{Val: 4}

	p.Right = &TreeNode{Val: 2}
	p.Right.Left = &TreeNode{Val: 4}
	p.Right.Right = &TreeNode{Val: 3}

	c := maxDepth(&p)

	fmt.Printf("-", c)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 思路
 *     深度遍历，碰到最后结点时候计算深度
 * 耗时
 *    执行用时 :8 ms, 在所有Go提交中击败了94.38%的用户
 *    内存消耗 :4.6 MB, 在所有Go提交中击败了36.96%的用户
 */

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxL := maxDepth(root.Left)
	maxR := maxDepth(root.Right)
	max := 0

	if maxL > maxR {
		max = maxL
	} else {
		max = maxR
	}

	return max + 1
}
