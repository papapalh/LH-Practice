package main

import "fmt"
/**
 * 题目
 *     删除链表中等于给定值 val 的所有节点。
 * 示例:
 *     输入: 1->2->6->3->4->5->6, val = 6
 *     输出: 1->2->3->4->5
 */
func main () {
	list := ListNode{Val:1}
	list.Next = &ListNode{Val:1}
	list.Next.Next = &ListNode{Val:3}
	list.Next.Next.Next = &ListNode{Val:5}
	list.Next.Next.Next.Next = &ListNode{Val:5}

	x := removeElements(&list, 1)
	fmt.Printf("-", x)

}
type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * 思路
 *     循环处理节点
 * 耗时
 *      执行用时 :12 ms, 在所有 Go 提交中击败了91.19%的用户
 *      内存消耗 :4.8 MB, 在所有 Go 提交中击败了58.16%的用户
*/
func removeElements(head *ListNode, val int) *ListNode {
	
	if head == nil {
		return head
	}

	res := head

	// 处理其他结点
	for head != nil {

		// 头结点等于目标值并且没有下一个结点，直接返回 nil
		if head.Val == val && head.Next == nil {
			return nil
		}

		// 头结点等于目标值，下结点赋值
		if head.Val == val && head.Next != nil {
			head.Val = head.Next.Val
			head.Next = head.Next.Next
			continue	
		}
		
		// 下一个结点是目标值
		if head.Next != nil && head.Next.Val == val {
			head.Next = head.Next.Next
			continue
		}

		head = head.Next
	}

	return res
}