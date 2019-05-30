<?php

/**
 * 题目
 *     给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 * 说明：不要使用任何内置的库函数，如  sqrt。
 * 示例 1：
 *     输入：16
 *     输出：True
 * 示例 2：
 *     输入：14
 *     输出：False
 */
class Solution
{
    /**
     * 思路
     *     二分法
     *     用中间值计算，检查是不是完全平方
     * 耗时
     *     执行用时 : 12 ms, 在Valid Perfect Square的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.6 MB, 在Valid Perfect Square的PHP提交中击败了100.00% 的用户
     */
    function isPerfectSquare($num)
    {
        if ($num == 1) {
            return true;
        }

        $low    = 0;
        $hei    = $num;

        while ($low < $hei) {
            $middle = $low + floor(($hei - $low) / 2);

            $result = $middle * $middle;

            if ($result > $num) {
                $hei = $middle;
            }
            elseif($result < $num) {
                $low = $middle + 1;
            }
            else {
                return true;
            }
        }

        return false;
    }
}
$s = new Solution();

var_dump($s->isPerfectSquare(14));
