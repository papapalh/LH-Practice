package main
/**
 * 题目
 *     翻转一棵二叉树。
 * 示例：
 *     输入：
 *             4
 *           /   \
 *          2     7
 *         / \   / \
 *        1   3 6   9
 *     输出：
 *             4
 *           /   \
 *          7     2
 *         / \   / \
 *        9   6 3   1
 * 备注:
 *     这个问题是受到 Max Howell 的 原问题 启发的 ：
 *     谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
 */
 func main() {
	tree := TreeNode{Val:4}
	tree.Left = &TreeNode{Val:2}
	tree.Left.Left = &TreeNode{Val:1}
	tree.Left.Right = &TreeNode{Val:3}
	tree.Right = &TreeNode{Val:7}
	tree.Right.Left = &TreeNode{Val:6}
	tree.Right.Right = &TreeNode{Val:9}

	invertTree(&tree)

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
 * 思路
 *     前序遍历
 *     每个结点的左右节点互换，完成反转
 * 耗时
 *     执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 *     内存消耗 :2.3 MB, 在所有 Go 提交中击败了45.87%的用户
 */
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}