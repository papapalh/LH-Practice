package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := TreeNode{Val: 6}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 0}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Left.Right.Left = &TreeNode{Val: 3}
	tree.Left.Right.Right = &TreeNode{Val: 5}

	tree.Right = &TreeNode{Val: 8}
	tree.Right.Left = &TreeNode{Val: 7}
	tree.Right.Right = &TreeNode{Val: 9}

	// fmt.Println(lowestCommonAncestor(&tree, tree.Left, &tree))
	fmt.Println(lowestCommonAncestor(&tree, tree.Left.Right.Left, tree.Left.Right.Right))

}

/**
 * 思路
 *    两次遍历，获取节点路径，取交点
 * 耗时
 *    执行用时：24 ms, 在所有 Go 提交中击败了85.04%的用户
 *    内存消耗：7.7 MB, 在所有 Go 提交中击败了13.50%的用户
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	ancestor := &TreeNode{}

	pPath := getPath(root, p)
	qPath := getPath(root, q)

	for i := 0; i < len(pPath) && i < len(qPath); i++ {
		if pPath[i].Val == qPath[i].Val {
			ancestor = pPath[i]
		}
	}

	return ancestor
}

func getPath(root, point *TreeNode) []*TreeNode {
	res := []*TreeNode{}

	for {
		if root == nil {
			break
		}

		// fmt.Println(root.Val)

		res = append(res, root)

		if point.Val < root.Val {
			root = root.Left
		} else if point.Val > root.Val {
			root = root.Right
		} else {
			break
		}
	}

	return res
}
