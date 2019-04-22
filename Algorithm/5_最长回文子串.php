<?php

/**
 * 题目
 *     给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 * 示例 1：
 *     输入: "babad"
 *     输出: "bab"
 *     注意: "aba" 也是一个有效答案。
 * 示例 2：
 *     输入: "cbbd"
 *     输出: "bb"
 */
class Solution
{
    /**
     * 思路
     *     暴力求解+动态规划
     *     将可能成为回文字符串的拉出来，进行判断
     *     如果已经处理过这个字符串，则不需要在处理
     * 耗时
     *     执行用时 : 28 ms, 在Longest Palindromic Substring的PHP提交中击败了70.00% 的用户
     *     内存消耗 : 14.9 MB, 在Longest Palindromic Substring的PHP提交中击败了44.44% 的用户
     * 时间复杂度
     *     O（n^2）
     */
    function longestPalindrome($s) {

        if (!$s) {
            return '';
        }

        if ($this->check($s)) {
            return $s;
        }

        $len = strlen($s);
        $dict = [];

        for ($i = 0; $i < $len; $i++) {

            $x = $s[$i];

            $dict[$x] = 1;

            if ($s[$i+1] == $s[$i]) {
                $dict[$x.$s[$i+1]] = 2;
            }

            $head   = $i;
            $footer = $i;

            for ($j = $i; $j < $len; $j++) {

                if (($head-1) < 0 || ($footer+1) > $len) {
                    break;
                }

                $x = $s[$head-1] . $x . $s[$footer+1];

                if (!$dict[$x]) {
                    $ret = $this->check($x);
                    if ($ret) {
                        $dict[$x] = $ret;
                    }
                }

                $head--;
                $footer++;
            }
        }

        if (!$dict) {
            return '';
        }

        return (string)array_search(max($dict), $dict);
    }

    /**
     * 检查是否为回文字符串
     */
    function check($str)
    {
        $length = strlen($str);

        $middle = floor(($length) / 2);

        for ($i=0; $i < $middle; $i++) {
            if ($str[$i] != $str[$length-1-$i]) {
                return 0;
            }
        }

        return $length;
    }
}

$s = new Solution();
var_dump($s->longestPalindrome("babaddtattarrattatddetartrateedredividerb"));