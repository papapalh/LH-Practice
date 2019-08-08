<?php

/**
 * 题目
 *     给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 * 示例 1:
 *     输入: 1->2->3->3->4->4->5
 *     输出: 1->2->5
 * 示例 2:
 *     输入: 1->1->1->2->3
 *     输出: 2->3
 */
class ListNode
{
    public $val = 0;
    public $next = null;
    function __construct($val)
    {
        $this->val = $val;
    }
}
/**
 * 思路
 *     有序链表，如果发现有两个结点的值相同
 *     则立flag，删除所有flag值结点
 * 耗时
 *     执行用时 :16 ms, 在所有 PHP 提交中击败了38.46%的用户
 *     内存消耗 :14.8 MB, 在所有 PHP 提交中击败了42.86%的用户
 */
class Solution
{
    function deleteDuplicates($head)
    {
        if (!$head) {
            return null;
        }

        // 构造头结点 随便链表不可能存在的数，为了避免头结点影响链表操作
        $l = new ListNode(999999);
        $l->next = $head;
        $flag = 999999;

        $res = $l;

        while ($l) {

            if ($l->next && $l->next->next && ($l->next->val == $l->next->next->val)) {
                $flag = $l->next->val;
            }

            if ($l->next && $l->next->val == $flag) {
                $l->next = $l->next->next;
                continue;
            }

            $l = $l->next;
        }

        return $res->next;
    }
}

$l = new ListNode(0);
$l->next = new ListNode(0);
$l->next->next = new ListNode(0);
$l->next->next->next = new ListNode(0);
$l->next->next->next->next = new ListNode(0);
$l->next->next->next->next->next = new ListNode(4);
$l->next->next->next->next->next->next = new ListNode(5);

$s = new Solution();
var_dump($s->deleteDuplicates($l));