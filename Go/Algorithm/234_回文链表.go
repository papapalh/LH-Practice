package main

import "fmt"
/**
 * 题目
 *     请判断一个链表是否为回文链表。
 * 示例 1:
 *     输入: 1->2
 *     输出: false
 * 示例 2:
 *     输入: 1->2->2->1
 *     输出: true
 * 进阶：
 *     你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 */
func main () {
	list := ListNode{Val:1}
	list.Next = &ListNode{Val:2}
	list.Next.Next = &ListNode{Val:3}
	list.Next.Next.Next = &ListNode{Val:4}
	list.Next.Next.Next.Next = &ListNode{Val:5}
	list.Next.Next.Next.Next.Next = &ListNode{Val:6}

	fmt.Printf("-", isPalindrome(&list))
}
type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * 思路（大神思路）
 *    思想很很简单，用2个指针，一个low，一个hei，hei是low的2倍，所以可以达到2分链表的效果 
 *    在移动指针时同时对前半部分链表进行入栈。
 *    最后直接比较被分开的2个链表。
 * 耗时
 *    执行用时 :24 ms, 在所有 Go 提交中击败了84.65%的用户
 *    内存消耗 :6.1 MB, 在所有 Go 提交中击败了64.75%的用户
 */
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	if head.Next.Next == nil {
		if head.Val != head.Next.Val {
			return false
		}
		return true
	}
	dict := []int{}
	low  := head.Next
	hei  := head.Next.Next

	dict = append(dict, head.Val)

	// 找到前半部分
	for {

		if hei.Next == nil {
			break
		}
		
		if hei.Next.Next == nil {
			dict = append(dict, low.Val)
			break
		}

		dict = append(dict, low.Val)

		hei = hei.Next.Next
		low = low.Next
	}

	low = low.Next

	len := len(dict)
	for low != nil {

		if low.Val != dict[len - 1] {
			return false
		}

		len--
		low = low.Next
	}

	return true
}