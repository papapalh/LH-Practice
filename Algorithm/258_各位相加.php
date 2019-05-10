<?php
/**
 * 题目
 *     给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。
 * 示例:
 *     输入: 38
 *     输出: 2 
 * 解释
 *     各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。
 * 进阶:
 *     你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？
 */
class Solution
{
    /**
     * 思路
     *     呃......这道题没动脑子，属于每天必练，但是不动脑子的类型
     *     最近除了算法，也好好整理下其他的知识点
     * 耗时
     *     执行用时 : 24 ms, 在Add Digits的PHP提交中击败了94.44% 的用户
     *     内存消耗 : 14.2 MB, 在Add Digits的PHP提交中击败了100.00% 的用户
     */
    function addDigits($num)
    {
        while (true) {
            $sum = array_sum(str_split($num));
            if (strlen($sum) == 1) {
                return $sum;
            }

            $num = $sum;
        }
    }
}

$s = new Solution();
var_dump($s->addDigits(38));