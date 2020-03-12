package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 题目
 *    请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
 *    题目太难懂
 * 耗时
 *    执行用时 :4 ms, 在所有 Go 提交中击败了83.12%的用户
 *    内存消耗 :2.9 MB, 在所有 Go 提交中击败了53.48%的用户
 */
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
