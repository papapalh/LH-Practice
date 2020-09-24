package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	tree := Node{Val: 1}
	tree.Children = []*Node{
		&Node{Val: 2, Children: []*Node{&Node{Val: 3, Children: []*Node{}}}},
		&Node{Val: 4, Children: []*Node{}},
		&Node{Val: 6, Children: []*Node{}},
	}

	fmt.Print(preorder(&tree))
}

/**
 * 递归方法
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：4 MB, 在所有 Go 提交中击败了33.65%的用户
 */
var res []int

func preorder1(root *Node) []int {
	res = []int{}
	xx(root)
	return res
}

func xx(root *Node) {

	if root == nil {
		return
	}

	res = append(res, root.Val)
	for _, n := range root.Children {
		xx(n)
	}
}

/**
 * 迭代方法(https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/solution/ncha-shu-de-qian-xu-bian-li-by-leetcode/)
 * 思路
 *    我们使用一个栈来帮助我们得到前序遍历，需要保证栈顶的节点就是我们当前遍历到的节点。
 *    我们首先把根节点入栈，因为根节点是前序遍历中的第一个节点。
 *    随后每次我们从栈顶取出一个节点 u，它是我们当前遍历到的节点，并把 u 的所有子节点逆序推入栈中。
 *    例如 u 的子节点从左到右为 v1, v2, v3，那么推入栈的顺序应当为 v3, v2, v1，这样就保证了下一个遍历到的节点（即 u 的第一个子节点 v1）出现在栈顶的位置。。
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：3.9 MB, 在所有 Go 提交中击败了78.28%的用户
 */
func preorder(root *Node) []int {
	res = []int{}

	if root == nil {
		return res
	}

	stack := []*Node{}

	stack = append(stack, root)

	for {
		if len(stack) == 0 {
			break
		}

		//栈顶元素出栈，如果有子节点，逆序推入栈
		top := stack[len(stack)-1]

		//栈顶出栈
		stack = stack[:len(stack)-1]

		//推如结果
		res = append(res, top.Val)

		//逆序入栈
		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack, top.Children[i])
		}
	}

	return res
}
