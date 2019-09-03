<?php
/**
 * 题目
 *     给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *     你的算法时间复杂度必须是 O(log n) 级别。
 *     如果数组中不存在目标值，返回 [-1, -1]。
 * 示例 1:
 *     输入: nums = [5,7,7,8,8,10], target = 8
 *     输出: [3,4]
 * 示例 2:
 *     输入: nums = [5,7,7,8,8,10], target = 6
 *     输出: [-1,-1]
 */
class Solution
{
    /**
     * 思路
     *     二分查找，在找到位置后，进行前后的数据校验比对
     * 耗时
     *     执行用时 :124 ms, 在所有 PHP 提交中击败了7.41%的用户
     *     内存消耗 :19.1 MB, 在所有 PHP 提交中击败了6.25%的用户
     */
    function searchRange($nums, $target)
    {
        if (!$nums) {
            return [-1, -1];
        }

        // 开始二分
        $hei    = count($nums);
        $min    = 0;
        $outPut = [];

        while ($hei > $min) {

            $middle = $min + floor(($hei - $min) / 2);

            if ($nums[$middle] == $target) {
                $outPut[] = $middle;

                $next = $middle + 1;
                $prev = $middle - 1;

                while (isset($nums[$next]) && $nums[$next] === $target) {
                    $outPut[] = $next;
                    $next++;
                }

                while (isset($nums[$prev]) && $nums[$prev] === $target) {
                    array_unshift($outPut, $prev);
                    $prev--;
                }

                break;
            }
            elseif ($nums[$middle] > $target) {
                $hei = $middle;
            }
            elseif ($nums[$middle] < $target) {
                $min = $middle + 1;
            }
        }

        $outPut = [
            $outPut ? $outPut[0] : -1,
            $outPut ? end($outPut) : -1,
        ];

        return $outPut;
    }
}

$s = new Solution();
var_dump($s->searchRange([2,2], 2));
