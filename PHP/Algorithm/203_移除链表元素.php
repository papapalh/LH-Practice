<?php
/**
 * 题目
 *     删除链表中等于给定值 val 的所有节点。
 * 示例:
 *     输入: 1->2->6->3->4->5->6, val = 6
 *     输出: 1->2->3->4->5
 */
class Solution
{
    /**
     * 思路
     *     循环处理节点
     * 耗时
     *     执行用时 : 36 ms, 在Remove Linked List Elements的PHP提交中击败了87.65% 的用户
     *     内存消耗 : 17.8 MB, 在Remove Linked List Elements的PHP提交中击败了48.98% 的用户
     */
    function removeElements($head, $val)
    {
        if (!$head) {
            return $head;
        }

        if ($head->val == $val) {
            $head = $head->next;
        }

        $res = $head;

        while ($head) {

            if ($head->next->val == $val) {
                $head->next = $head->next->next;
            }

            $head = $head->next;
        }

        return $res;
    }
}