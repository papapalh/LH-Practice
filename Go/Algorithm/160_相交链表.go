package main

import "fmt"

/**
 * 题目
 *     编写一个程序，找到两个单链表相交的起始节点。
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
	fmt.Printf("-", x)

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
 *    执行用时 :60 ms, 在所有 Go 提交中击败了78.35%的用户
 *    内存消耗 :6.4 MB, 在所有 Go 提交中击败了83.42%的用户
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

	return A
}
