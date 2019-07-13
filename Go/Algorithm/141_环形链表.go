package main

import "fmt"
/**
 * 题目
 *     给定一个链表，判断链表中是否有环。
 *     为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
 * 图示
 *     https://leetcode-cn.com/problems/linked-list-cycle
 */
func main () {
	list := ListNode{Val:3}
	list.Next = &ListNode{Val:2}
	list.Next.Next = &ListNode{Val:0}
	list.Next.Next.Next = &ListNode{Val:4}
	list.Next.Next.Next.Next = list.Next

	x := hasCycle(&list)
	fmt.Printf("-", x)

}
type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * 思路
 *    哈希表
 * 耗时
 *    执行用时 :32 ms, 在所有 Go 提交中击败了5.70%的用户
 *    内存消耗 :6.3 MB, 在所有 Go 提交中击败了5.40%的用户
 */
func hasCycle(head *ListNode) bool {

	dict := make(map[*ListNode]int)

	for head != nil {
		
		if _, ok := dict[head]; ok {
			return true
		}

		dict[head] = head.Val
		head = head.Next

	}

	return false
}

