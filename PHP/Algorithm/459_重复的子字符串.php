<?php

/**
 * 题目
 *     给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
 * 示例 1:
 *     输入: "abab"
 *     输出: True
 *     解释: 可由子字符串 "ab" 重复两次构成。
 * 示例 2:
 *     输入: "aba"
 *     输出: False
 * 示例 3:
 *     输入: "abcabcabcabc"
 *     输出: True
 *     解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 */
class Solution
{
    /**
     * 思路
     *     求解重复字符串问题改变为求解字符串长度的因子数
     * 详解
     *     abab 字符串 长度为4
     *         如果想为重复字符串，则必为 a组成 / ab 组成
     *         话句话来说也就是由长度的因子 1/2 组成
     *         这样的话可以求解因子数，在计算 aaaa/abab 是否等于原子串
     *     同理 abcabcabcabcabc 长度为 15
     *         因数为 1/3/5
     *         计算 aaaaaaaaaaaaaaa/abcabcabcabcabc/abcabcabcabcabcabc 是否等于子串即可
     * 耗时
     *     执行用时 : 140 ms, 在Repeated Substring Pattern的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.5 MB, 在Repeated Substring Pattern的PHP提交中击败了100.00% 的用户
     *
     */
    function repeatedSubstringPattern($s)
    {
        $len    = strlen($s);
        $factor = [];

        // 查找因子
        for ($i = 1; $i < $len; $i++) {

            if (!($len%$i)) {
                $factor[]= $i;
            }
        }

        // 拼装因子
        if (!$factor) {
            return false;
        }

        // 拼装因子字符串
        foreach ($factor as $f) {
            $count = $len/$f;
            $str = '';

            for ($i = 0; $i < $count; $i++) {
                $str .= substr($s, 0, $f);
            }

            if ($str == $s) {
                return true;
            }

        }

        return false;
    }
}


$s = new Solution();
var_dump($s->repeatedSubstringPattern('abcabcabcabc'));