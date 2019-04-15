<?php

/**
 * 题目
 *     反转一个单链表。
 * 示例:
 *     输入: 1->2->3->4->5->NULL
 *     输出: 5->4->3->2->1->NULL
 * 进阶:
 *     你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 */
class Solution
{
    /**
     * 思路1
     *     先遍历出链表的元素入栈
     *     随后在遍历链表时赋值，完成反转
     * 耗时
     *    执行用时 : 16 ms, 在Reverse Linked List的PHP提交中击败了100.00% 的用户
     *    内存消耗 : 15.7 MB, 在Reverse Linked List的PHP提交中击败了100.00% 的用户
     */
     function reverseList1($head) {
     	$arr = [];
     	$node = $head;
     	while (isset($node->val)) {
     		$arr[] = $node->val;
     		$node = $node->next;
     	}

     	$node = $head;
     	while (isset($node->val)) {
     		$node->val = array_pop($arr);
     		$node = $node->next;
     	}

     	return $head;
     }

    /**
     * 思路2
     *     在遍历链表时候，记录当前链表节点，作为后一个节点的next值
     *     完成反转
     *     因为又开辟了空间，所以在内存消耗上会比上面的方式大一些
     * 耗时
     *    执行用时 : 16 ms, 在Reverse Linked List的PHP提交中击败了100.00% 的用户
     *    内存消耗 : 16.1 MB, 在Reverse Linked List的PHP提交中击败了81.73% 的用户
     */
    function reverseList2($head)
    {
    	$n = '';

    	while ($head) {
    		
    		$temp = $head->next;

        	$head->next = $n;
        	$n = $head;

        	$head = $temp;
    	}

    	return $n;
    }
}