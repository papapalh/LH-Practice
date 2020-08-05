package main

import "fmt"

// 编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
// 示例1:
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
// 示例2:
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
// 提示：
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
// 进阶：
// 如果不得使用临时缓冲区，该怎么解决？
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := ListNode{Val: 1}
	l.Next = &ListNode{Val: 2}
	l.Next.Next = &ListNode{Val: 3}
	l.Next.Next.Next = &ListNode{Val: 3}
	l.Next.Next.Next.Next = &ListNode{Val: 2}
	l.Next.Next.Next.Next.Next = &ListNode{Val: 1}

	fmt.Println(removeDuplicateNodes(&l))
}

//思路
//    O(n)的算法复杂度，O(n)的空间复杂度
//    不耗费额外空间的方法为双循环暴力破解(O(n2))
//耗时
//	  执行用时：20 ms, 在所有 Go 提交中击败了23.63%的用户
//    内存消耗：6 MB, 在所有 Go 提交中击败了75.00%的用户
func removeDuplicateNodes(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	//不循环当前节点，先添加进来
	m := map[int]bool{
		head.Val: true,
	}

	l := head

	for l.Next != nil {
		if m[l.Next.Val] == true {
			l.Next = l.Next.Next
		} else {
			m[l.Next.Val] = true
			l = l.Next
		}
	}

	return head
}
