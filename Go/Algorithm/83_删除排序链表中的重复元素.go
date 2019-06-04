package main

/**
 * 题目
 *    给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 * 示例 1:
 *    输入: 1->1->2
 *    输出: 1->2
 * 示例 2:
 *    输入: 1->1->2->3->3
 *    输出: 1->2->3
 */
func main() {
	l := ListNode{Val:1}
	l.Next = &ListNode{Val:1}
	l.Next.Next = &ListNode{Val:2}

	deleteDuplicates(&l)
}
type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * 思路
 *    链表操作，过滤重复链表
 * 耗时
 *    执行用时 : 4 ms, 在Remove Duplicates from Sorted List的Go提交中击败了100.00% 的用户
 *    内存消耗 : 3.2 MB, 在Remove Duplicates from Sorted List的Go提交中击败了27.01% 的用户
 */
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil{
		return nil
	}

	res := head

	for res.Next != nil {

		if res.Val == res.Next.Val {
			res.Next = res.Next.Next
			continue
		}
		res = res.Next
	}

	return head
}