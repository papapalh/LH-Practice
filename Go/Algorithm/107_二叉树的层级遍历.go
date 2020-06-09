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
	p := TreeNode{Val: 3}
	p.Left = &TreeNode{Val: 9}

	p.Right = &TreeNode{Val: 20}
	p.Right.Left = &TreeNode{Val: 15}
	p.Right.Right = &TreeNode{Val: 7}

	c := levelOrderBottom(&p)

	fmt.Printf("-", c)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 思路
 *     层级遍历，数组翻转
 * 耗时
 *    执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
 *    内存消耗 :6.2 MB, 在所有Go提交中击败了45.98%的用户
 */
func levelOrderBottom(root *TreeNode) [][]int {

	res := make([][]int, 0)

	if root == nil {
		return nil
	}

	depth := 0 //深度
	treePool := map[int][]*TreeNode{}

	treePool[depth] = []*TreeNode{root}

	for {
		//数组超限判断
		if len(treePool[depth]) == 0 {
			break
		}

		for _, node := range treePool[depth] {

			if node.Left != nil {
				treePool[depth+1] = append(treePool[depth+1], node.Left)
			}

			if node.Right != nil {
				treePool[depth+1] = append(treePool[depth+1], node.Right)
			}
		}
		depth++
	}

	//反转数组
	for i := len(treePool) - 1; i >= 0; i-- {

		tmp := []int{}

		for _, node := range treePool[i] {
			tmp = append(tmp, node.Val)
		}

		res = append(res, tmp)
	}

	return res
}
