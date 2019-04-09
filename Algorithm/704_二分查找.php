<?php
/**
 * 题目
 *     给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
 * 示例 1:
 *     输入: nums = [-1,0,3,5,9,12], target = 9
 *     输出: 4
 *     解释: 9 出现在 nums 中并且下标为 4
 * 示例 2:
 *     输入: nums = [-1,0,3,5,9,12], target = 2
 *     输出: -1
 *     解释: 2 不存在 nums 中因此返回 -1
 * 提示：
 *     你可以假设 nums 中的所有元素是不重复的。
 *     n 将在 [1, 10000]之间。
 *     nums 的每个元素都将在 [-9999, 9999]之间。
 */

class Solution
{
    /**
     * 思路
     *     二分查找
     *     还是二分查找，二分查找的关键要素
     *     有序！
     * 耗时
     *     执行用时 : 204 ms, 在Binary Search的PHP提交中击败了12.50% 的用户
     *     内存消耗 : 15.9 MB, 在Binary Search的PHP提交中击败了100.00% 的用户
     */
    function search($nums, $target)
    {
        $low = 0;
        $hei = count($nums);
        $i = 0;

        if (!$hei) {
            return -1;
        }

        while ($low < $hei) {
            $middle = $low + floor(($hei - $low) / 2);

            if ($target < $nums[$middle]) {
                $hei = $middle;
            }
            elseif ($target > $nums[$middle]) {
                $low = $middle + 1;
            }
            else {
                return $middle;
            }

            $i++;
        }

        return -1;
    }
}

$s = new Solution ();
var_dump($s->search([-1,0,3,5,9,12], 9));
