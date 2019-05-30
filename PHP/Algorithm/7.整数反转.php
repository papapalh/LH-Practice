<?php
 /**
 * 题目
 *    给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *    假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31, 2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回0
 * 用例
 *    输入: 123
 *    输出: 321
 *    输入: -123
 *    输出: -321
 */
class Solution
{
    /**
     * 思路
     *    如果数字为负，则取绝对值处理，最后加符号
     *    超出范围，则爆 0
     * 耗时
     *     执行用时 : 28 ms, 在Reverse Integer的PHP提交中击败了64.44% 的用户
     *     内存消耗 : 14.5 MB, 在Reverse Integer的PHP提交中击败了100.00% 的用
     * 时间复杂度 
     *     O(n)
     */
    function reverse($x)
    {
        $num = $x < 0 ? abs($x) : $x;

        $reverseNum = 0;
        // 整数反转
        while($num >= 1) {
            $reverseNum = ($reverseNum * 10) + ($num % 10);
            $num /= 10;
        }

        if ($reverseNum>=2147483647 || $reverseNum<=-2147483648) {
            return 0;
        }

        if ($x < 0) {
            $reverseNum *= -1;
        }

        return $reverseNum;
    }
}
