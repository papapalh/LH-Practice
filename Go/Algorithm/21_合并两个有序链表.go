package main

import "fmt"
/**
 * 题目
 *     将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 * 示例
 *     输入：1->2->4, 1->3->4
 *     输出：1->1->2->3->4->4
 */
/**
 * 思路
 *     归并算法
 *     递归处理
 * 耗时
 *     执行用时 : 4 ms, 在Merge Two Sorted Lists的Go提交中击败了96.50% 的用户
 *     内存消耗 : 2.6 MB, 在Merge Two Sorted Lists的Go提交中击败了33.97% 的用户
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func main () {

	l1 := ListNode{Val : 1}
	l1.Next = &ListNode{Val : 2}
	l1.Next.Next = &ListNode{Val : 5}

	l2 := ListNode{Val : 1}
	l2.Next = &ListNode{Val : 3}
	l2.Next.Next = &ListNode{Val : 5}

	l3 := mergeTwoLists(&l1, &l2)

	fmt.Println(l3)
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var res *ListNode

	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}

	return res
}