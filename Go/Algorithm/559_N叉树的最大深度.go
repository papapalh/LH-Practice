package main

/**
 *
 */
func main() {
	tree := Node{Val: 1}
	tree.Children = []*Node{
		&Node{Val: 2, Children: []*Node{&Node{Val: 3, Children: []*Node{}}}},
		&Node{Val: 4, Children: []*Node{}},
		&Node{Val: 6, Children: []*Node{}},
	}

	maxDepth(&tree)

}

type Node struct {
	Val      int
	Children []*Node
}

/**
 * 题目
 *    给定一个 N 叉树，找到其最大深度。
 *    最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
 * 思路
 *    广度优先遍历
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：3.6 MB, 在所有 Go 提交中击败了6.41%的用户
 */
func maxDepth(root *Node) int {

	if root == nil {
		return 0
	}

	depth := 0
	treeMap := map[int][]*Node{}
	treeMap[depth] = []*Node{
		root,
	}

	for {

		if len(treeMap[depth]) == 0 {
			break
		}

		for _, n := range treeMap[depth] {
			treeMap[depth+1] = append(treeMap[depth+1], n.Children...)
		}

		depth++

	}

	return depth
}
