package main

import "fmt"

/**
 * 题目
 *     反转一个单链表。
 * 示例:
 *     输入: 1->2->3->4->5->NULL
 *     输出: 5->4->3->2->1->NULL
 * 进阶:
 *     你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 */
func main() {
	list := ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Printf("-", reverseList(&list))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 思路
 *     在遍历链表时候，记录当前链表节点，作为后一个节点的next值
 *     建立头结点来维护整个链表
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：2.7 MB, 在所有 Go 提交中击败了23.86%的用户
 */
func reverseList(head *ListNode) *ListNode {
	var l *ListNode

	for head != nil {

		tmp := &ListNode{Val: head.Val}

		tmp.Next = l

		l = tmp

		head = head.Next
	}

	return l
}
