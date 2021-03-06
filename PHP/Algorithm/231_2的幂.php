<?php

/**
 * 题目
 *     给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
 * 示例 1:
 *     输入: 1
 *     输出: true
 *     解释: 20 = 1
 * 示例 2:
 *     输入: 16
 *     输出: true
 *     解释: 24 = 16
 * 示例 3:
 *     输入: 218
 *     输出: false
 */
class Solution
{
    /**
     * 思路(大神解释)
     *     重点在于对位运算符的理解
     * 解法：&运算，同1则1。 return (n > 0) && (n & -n) == n;
     *     2的幂次方在二进制下，只有1位是1，其余全是0。
     *     例如:8---00001000。负数的在计算机中二进制表示为补码(原码->正常二进制表示，原码按位取反(0-1,1-0)，最后再+1。
     *     然后两者进行与操作，得到的肯定是原码中最后一个二进制的1。
     *     例如8&(-8)->00001000 & 11111000 得 00001000，即8。
     *     建议自己动手算一下，按照这个流程来一遍，加深印象。
     * 耗时
     *     执行用时 : 24 ms, 在Power of Two的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.9 MB, 在Power of Two的PHP提交中击败了100.00% 的用户
     */
    function isPowerOfTwo($n) {
        return ($n > 0) && ($n & -$n) == $n;
    }
}

$s = new Solution();
var_dump($s->isPowerOfTwo(6));