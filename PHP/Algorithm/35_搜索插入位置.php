<?php

/**
 * 题目
 *     给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
 *     如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *     你可以假设数组中无重复元素。
 * 用例
 *     输入: [1,3,5,6], 5
 *     输出: 2
 *     输入: [1,3,5,6], 2
 *     输出: 1
 *     输入: [1,3,5,6], 7
 *     输出: 4
 */
class Solution
{
    /**
     * 思路
     *     二分法
     *     条件（排序数组/无重复值）
     * 耗时
     *     执行用时 : 72 ms, 在Search Insert Position的PHP提交中击败了13.64% 的用户
     *     内存消耗 : 15.8 MB, 在Search Insert Position的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(log2n)
     */
    function searchInsert($nums, $target) {

        $low    = 0;            // 二分 - 前置
        $height = count($nums); // 二分 - 后置

        while ($low < $height) {
            $middle = $low + floor(($height - $low) / 2);
            if ($nums[$middle] > $target) {
                $height = $middle;
            }
            elseif($nums[$middle] < $target) {
                $low = $middle + 1;
            }
            else {
                return $middle;
            }

            echo $low,$height,$middle,"\n";
        }

        return $low;
    }
}

$s = new Solution();
var_dump($s->searchInsert([1,3,5,6], 7));