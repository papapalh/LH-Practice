<?php

/**
 * 题目
 *     给定两个数组，编写一个函数来计算它们的交集。
 * 示例 1:
 *     输入: nums1 = [1,2,2,1], nums2 = [2,2]
 *     输出: [2]
 * 示例 2:
 *     输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 *     输出: [9,4]
 * 说明:
 *     输出结果中的每个元素一定是唯一的。
 *     我们可以不考虑输出结果的顺序。
 */

class Solution
{
    /**
     * 思路
     *     哈希表
     *     $nums1 和 $nums2 都需要遍历一次，获取HASH值相同的部分
     * 耗时
     *     执行用时 : 32 ms, 在Intersection of Two Arrays的PHP提交中击败了80.00% 的用户
     *     内存消耗 : 14.9 MB, 在Intersection of Two Arrays的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O（n）
     */
    function intersection($nums1, $nums2) {
        $hash = [];
        $arrOutput = [];

        foreach ($nums1 as $n) {
            $hash[$n] = true;
        }

        foreach ($nums2 as $n) {
            if ($hash[$n]) {
                $arrOutput[$n] = $n;
            }
        }

        return $arrOutput;
    }
}

$s = new Solution();
var_dump($s->intersection([1,2,2,1], [2,2]));