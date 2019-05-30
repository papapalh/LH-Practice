<?php
/**
 * 题目
 *     给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 * 示例 1:
 *     输入: 1->1->2
 *     输出: 1->2
 * 示例 2:
 *     输入: 1->1->2->3->3
 *     输出: 1->2->3
 */
class Solution
{
    /**
     * 思路
     *     一次遍历链表
     *     发现当前结点与下一个结点相同，则修改下一个结点位置
     *     如果不相同，则认为是不同值，则略过
     * 耗时
     *     执行用时 : 24 ms, 在Remove Duplicates from Sorted List的PHP提交中击败了76.09% 的用户
     *     内存消耗 : 14.7 MB, 在Remove Duplicates from Sorted List的PHP提交中击败了73.68% 的用户
     * 时间复杂度
     *     O（n）
     */
    function deleteDuplicates($head)
    {
        $res = $head;

        while ($head) {

            if ($head->val === $head->next->val) {
                $head->next = $head->next->next;
                continue;
            }

            $head = $head->next;
        }

        return $res;
    }
}