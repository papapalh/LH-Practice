package main

import "fmt"

/**
 * 题目
 *    如果两个链表没有交点，返回 null 。
 *    在返回结果后，两个链表仍须保持原有的结构。
 *    可假定整个链表结构中没有循环。
 *    程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
 * 图示
 *     https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
 */
func main() {

	l := ListNode{Val: 7}
	l.Next = &ListNode{Val: 8}
	l.Next.Next = &ListNode{Val: 9}

	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &l

	l2 := ListNode{Val: 4}
	l2.Next = &ListNode{Val: 5}
	l2.Next.Next = &l

	x := getIntersectionNode(&l1, &l2)
	fmt.Println(x)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 思路（Leet Code 思路）
 *    开两个指针分别遍历这两个链表，在第一次遍历到尾部的时候，指向另一个链表头部继续遍历，这样会抵消长度差。
 *    如果链表有相交，那么会在中途相等，返回相交节点；
 *    如果链表不相交，那么最后会 nil == nil，返回 nil；
 * 耗时
 *    执行用时：44 ms, 在所有 Go 提交中击败了92.16%的用户
 *    内存消耗：8.1 MB, 在所有 Go 提交中击败了39.13%的用户
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	A, B := headA, headB

	for A != B {

		if A == nil {
			A = headB
		} else {
			A = A.Next
		}

		if B == nil {
			B = headA
		} else {
			B = B.Next
		}
	}

	return B
}
