package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//进阶：你能尝试使用一趟扫描实现吗？
//示例 1：
//	输入：head = [1,2,3,4,5], n = 2
//	输出：[1,2,3,5]
//示例 2：
//	输入：head = [1], n = 1
//	输出：[]
//示例 3：
//	输入：head = [1,2], n = 1
//	输出：[1]
func main() {

	l := ListNode{Val:1}
	//l.Next = &ListNode{Val:2}
	//l.Next.Next = &ListNode{Val:3}
	//l.Next.Next.Next = &ListNode{Val:4}
	//l.Next.Next.Next.Next = &ListNode{Val:5}

	removeNthFromEnd(&l, 1)

	xx := &l

	for xx != nil  {
		fmt.Println(xx.Val)
		xx = xx.Next
	}
}

//思路
//	空间换时间, 使用切片记录链表指针位置，最后统一处理
//耗时
//	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//	内存消耗：2.3 MB, 在所有 Go 提交中击败了9.22%的用户
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	res := head

	l := head

	dict := []*ListNode{}

	for l != nil  {
		dict = append(dict, l)
		l = l.Next
	}

	if len(dict) == 1 {
		return nil
	} else if n == 1 {
		dict[len(dict)-n-1].Next = nil
	} else if n == len(dict) {
		res = head.Next
	} else {
		dict[len(dict)-n-1].Next = dict[len(dict)-n+2-1]
	}

	return res
}