<?php

/**
 * 题目
 *     给定一个整数数组，判断是否存在重复元素。
 *     如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
 * 示例 1:
 *     输入: [1,2,3,1]
 *     输出: true
 * 示例 2:
 *     输入: [1,2,3,4]
 *     输出: false
 * 示例 3:
 *     输入: [1,1,1,3,3,4,3,2,4,2]
 *     输出: true
 */
class Solution
{
    /**
     * 思路
     *     哈希表
     * 耗时
     *     执行用时 : 96 ms, 在Contains Duplicate的PHP提交中击败了78.15% 的用户
     *     内存消耗 : 20.9 MB, 在Contains Duplicate的PHP提交中击败了71.57% 的用户
     * 时间复杂度
     *     O（n）
     */
    function containsDuplicate($nums)
    {
        $dict = [];
        foreach ($nums as $n) {

            if (isset($dict[$n])) {
                return true;
            }
            $dict[$n] = $n;
        }

        return false;
    }
}

$s = new Solution();
var_dump($s->containsDuplicate([1,2,3,1]));