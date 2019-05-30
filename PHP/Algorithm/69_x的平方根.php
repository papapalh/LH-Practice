<?php

/**
 * 题目
 *     实现 int sqrt(int x) 函数。
 *     计算并返回 x 的平方根，其中 x 是非负整数。
 *     由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 示例 1:
 *     输入: 4
 *     输出: 2
 * 示例 2:
 *     输入: 8
 *     输出: 2
 *     说明: 8 的平方根是 2.82842...,
 *     由于返回类型是整数，小数部分将被舍去。
 */
class Solution
{
    /**
     * 思路
     *     二分法
     *     根据题目可知，我们要找出的数是不满足的平方条件的最后一个小于 $x 的数
     *     所以在二分时候记录每一个小于 $x 的数，赋予 $flag
     *     如果刚好数字刚好有平方根，则返回
     *     如果没有，则返回小的一个数，也就是 $flag
     * 耗时
     *     执行用时 : 24 ms, 在Sqrt(x)的PHP提交中击败了94.59% 的用户
     *     内存消耗 : 14.8 MB, 在Sqrt(x)的PHP提交中击败了100.00% 的用户
     */
    function mySqrt($x)
    {
        $low  = 0;
        $hei  = $x;
        $flag = 0;

        if (!$x) {
            return 0;
        }
        if ($x == 1) {
            return 1;
        }

        while ($hei > $low) {
            $middle = $low + floor(($hei - $low) / 2);

            if ($middle * $middle > $x) {
                $hei = $middle;
            }
            elseif ($middle * $middle < $x) {
                $flag = $middle;
                $low = $middle+1;
            }
            else {
                return $middle;
            }
        }

        return $flag;
    }
}

$s = new Solution();
var_dump($s->mySqrt(8));