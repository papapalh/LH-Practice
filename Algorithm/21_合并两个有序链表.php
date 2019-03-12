<?php

/**
 * 题目
 *     将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 * 示例
 *     输入：1->2->4, 1->3->4
 *     输出：1->1->2->3->4->4
 */
class Solution
{
    /**
     * 思路
     *     归并排序(有序/链表)
     *     两个 有序 的合并，采用归并算法，比每一串链表的头结点大小
     * 注意
     *     由于 Leet Code 采用对象的方式作为结点，本身是不可见的，可以认为是对象
     *     PHP 无法直接执行，必须去 Leet Code 执行
     * 耗时
     *     执行用时 : 24 ms, 在Merge Two Sorted Lists的PHP提交中击败了33.33% 的用户
     *     内存消耗: 14.8 MB, 在Merge Sorted Array的PHP提交中击败了100.00% 的用户
     * 时间复杂度 O(n)
     */
    function mergeTwoLists($l1, $l2)
    {
        if ($l1 === null) {
            return $l2;
        }
        if ($l2 === null) {
            return $l1;
        }

        // $res = new ListNode();
        if ($l1->val < $l2->val) {
            $res = $l1;
            $l1 = $l1->next;
        } else {
            $res = $l2;
            $l2 = $l2->next;
        }

        $p = $res;
        while ($l1 && $l2) {
            if ($l1->val <= $l2->val) {
                $p->next = $l1;
                $l1 = $l1->next;
                $p = $p->next;
            } else {
                $p->next = $l2;
                $l2 = $l2->next;
                $p = $p->next;
            }
        }
        if ($l1 != null) {
            $p->next = $l1;
        }
        if ($l2 != null) {
            $p->next = $l2;
        }

        return $res;
    }
}

