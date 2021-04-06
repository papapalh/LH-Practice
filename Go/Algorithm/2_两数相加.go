package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

//题目
//	给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//	请你将两个数相加，并以相同形式返回一个表示和的链表。
//	你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//示例 1：
//	输入：l1 = [2,4,3], l2 = [5,6,4]
//	输出：[7,0,8]
//	解释：342 + 465 = 807.
//示例 2：
//	输入：l1 = [0], l2 = [0]
//	输出：[0]
//示例 3：
//	输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//	输出：[8,9,9,9,0,0,0,1]
func main() {

	l1 := ListNode{Val:9}
	l1.Next = &ListNode{Val:9}
	l1.Next.Next = &ListNode{Val:9}
	l1.Next.Next.Next = &ListNode{Val:9}
	l1.Next.Next.Next.Next = &ListNode{Val:9}
	l1.Next.Next.Next.Next.Next = &ListNode{Val:9}
	l1.Next.Next.Next.Next.Next.Next = &ListNode{Val:9}

	l2 := ListNode{Val:9}
	l2.Next = &ListNode{Val:9}
	l2.Next.Next = &ListNode{Val:9}
	l2.Next.Next.Next = &ListNode{Val:9}

	addTwoNumbers(&l1,&l2)
}

//思路
//	两个链表一直向后, >10 则进位
//耗时
//	执行用时：16 ms, 在所有 Go 提交中击败了39.81%的用户
//	内存消耗：4.6 MB, 在所有 Go 提交中击败了90.64%的用户
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//头节点
	res := ListNode{Val:0}
	f := 0

	tmp := &res


	for l1 != nil || l2 != nil {

		l1Val := 0
		l2Val := 0

		if l1 != nil  {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 != nil  {
			l2Val = l2.Val
			l2 = l2.Next
		}

		//处理进位
		//fmt.Println(l1Val + l2Val + f)
		if l1Val + l2Val + f >= 10 {
			tmp.Next = &ListNode{Val:l1Val + l2Val + f - 10}
			f = 1
		} else {
			tmp.Next = &ListNode{Val:l1Val + l2Val + f}
			f = 0
		}

		tmp = tmp.Next
	}

	if f != 0 {
		tmp.Next = &ListNode{Val:f}
	}

	return res.Next
}