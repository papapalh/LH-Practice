<?php
/**
 * 题目
 *    判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *    你能不将整数转为字符串来解决这个问题吗？
 * 用例
 *    输入: 输入: 121
 *    输出: true
 *    输入: -121
 *    解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *    输出: false
 */
class Solution
{
    /**
     * 思路
     *    回文数，不用字符串处理，这里采用的是官方的做法
     *    首先 负数 因为有(-)，所以肯定不是回文数，首先排除
     *    对于一个数，首先 %10 获取它的最后一位，最后一位当做新数的首位，同时 旧数/10 新数*10 在进行，可以得到这个数的回文数
     *    最后比对两个数，回文数肯定是相同的
     * 耗时
     *     执行用时 : 164 ms, 在Palindrome Number的PHP提交中击败了26.17% 的用户
     *     内存消耗 : 14.8 MB, 在Palindrome Number的PHP提交中击败了100.00% 的用户
     * 时间复杂度 
     *     O(n)
     */
    function isPalindrome($x)
    {
        // 所有负数都不可能是回文
        if ($x < 0) {
            return false;
        }

        $initNum    = $x;
        $reverseNum = 0;
        while ($x >= 1) {
            $reverseNum =  ($reverseNum * 10) + ($x % 10);
            $x /= 10;
        }

        if ($reverseNum == $initNum) {
            return true;
        }
        return false;
       
    }
}

$s = new Solution();
echo $s->isPalindrome(121);