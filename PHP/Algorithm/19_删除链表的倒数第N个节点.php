<?php

/**
 * 题目
 *     给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 * 示例：
 *     给定一个链表: 1->2->3->4->5, 和 n = 2.
 *     当删除了倒数第二个节点后，链表变为 1->2->3->5.
 * 说明：
 *     给定的 n 保证是有效的。
 * 进阶：
 *     你能尝试使用一趟扫描实现吗？
 */
class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}
class Solution
{
    /**
     * 思路
     *     一次遍历，设置头结点，快慢指针，拉开删除距离，指定节点删除
     * 耗时
     *     执行用时 :8 ms, 在所有 PHP 提交中击败了89.89%的用户
     *     内存消耗 :14.8 MB, 在所有 PHP 提交中击败了32.73%的用户
     * 时间复杂度
     *     O（n）
     */
    function removeNthFromEnd($head, $n)
    {
        $hei = 0;
        $low = 0;

        $l = new ListNode(0);
        $l->next = $head;
        $res = $l;
        $dict = [];

        if ($n == 1 && !$head->next) {
            return null;
        }

        while (true) {

            if (!$l) {
                $dict[$low]->next = $dict[$low]->next->next;
                break;
            }

            if ($hei > $n) $low++;

            $dict[$hei] = $l;
            $hei++;
            $l = $l->next;
        }

        return $res->next;
    }
}

$l = new ListNode(1);
$l->next = new ListNode(2);
$l->next->next = new ListNode(3);
$l->next->next->next = new ListNode(4);
$l->next->next->next->next = new ListNode(5);

$s = new Solution();
var_dump($s->removeNthFromEnd($l, 2));