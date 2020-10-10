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
 * 思路1
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

/**
 * 思路2
 *    在方法一中，我们对从根节点开始，通过遍历找出到达节点 pp 和 qq 的路径，一共需要两次遍历。我们也可以考虑将这两个节点放在一起遍历。
 *    整体的遍历过程与方法一中的类似：
 *    我们从根节点开始遍历；
 *        如果当前节点的值大于 pp 和 qq 的值，说明 pp 和 qq 应该在当前节点的左子树，因此将当前节点移动到它的左子节点；
 *        如果当前节点的值小于 pp 和 qq 的值，说明 pp 和 qq 应该在当前节点的右子树，因此将当前节点移动到它的右子节点；
 *        如果当前节点的值不满足上述两条要求，那么说明当前节点就是「分岔点」。此时，pp 和 qq 要么在当前节点的不同的子树中，要么其中一个就是当前节点。
 */
func lowestCommonAncestor2(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}
