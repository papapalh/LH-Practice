<?php

/**
 * 题目
 *    给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *    你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * 示例:
 *     给定 1->2->3->4, 你应该返回 2->1->4->3.
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
 *     遍历整个链表，进行两两反转
 * 耗时
 *     执行用时 :12 ms, 在所有 PHP 提交中击败了50.75%的用户
 *     内存消耗 :14.8 MB, 在所有 PHP 提交中击败了32.56%的用户
 */
class Solution
{
    function swapPairs($head)
    {
        if (!$head) {
            return [];
        }

        // 头结点挂载
        $l = new ListNode(9999);
        $l->next = $head;

        $res = $l;

        while ($l) {

            if ($l->next && $l->next->next) {

                $other = $l->next->next->next;

                $tmp = $l->next;

                $l->next = $l->next->next;

                $l->next->next = $tmp;

                $l->next->next->next = $other;
            }

            $l = $l->next->next;
        }

        return $res->next;
    }
}

$l = new ListNode(1);
$l->next = new ListNode(2);
$l->next->next = new ListNode(3);
$l->next->next->next = new ListNode(4);

$s = new Solution();
var_dump($s->swapPairs($l));