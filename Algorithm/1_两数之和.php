<?php
/**
 * 题目
 *    给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *    你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 * 用例
 *     输入: nums = [2, 7, 11, 15], target = 9
 *     因为 nums[0] + nums[1] = 2 + 7 = 9
 *     输出: [0, 1]
 */
class Solution
{
    /**
     * 思路
     *    暴力破解 O(n^2) 的复杂度首先被抛弃
     *    Hash 表做法是肯定的，问题是进行几次 Hash 循环
     *    可以进行两次 Hash ,找出对应的键 时间复杂度为 O(n+n) 也就是 O(n)
     *    其实完全可以进行一次循环，在循环时判断是否存在，抛出
     * 耗时
     *     执行用时 : 28 ms, 在Two Sum的PHP提交中击败了92.54% 的用户
     *     内存消耗 : 15.5 MB, 在Two Sum的PHP提交中击败了100.00% 的用户
     * 时间复杂度 
     *     O(n)
     */
    function twoSum($nums, $target)
    { 
        $hash = [];
        $count = count($nums);
        for ($i = 0; $i <= $count; $i++) {

            if (isset($hash[(string)($target - $nums[$i])])) {
                if ($hash[(string)($target - $nums[$i])] != $i) {
                    return [$hash[(string)($target - $nums[$i])], $i];
                }
            }

            $hash[(string)$nums[$i]] = (string)$i;
        }
    } 
}

$arr = [3,3];
$tag = 6;

$s = new Solution();
print_r($s->twoSum($arr, $tag));