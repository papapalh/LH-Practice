<?php
/**
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 说明
 *     本题中，我们将空字符串定义为有效的回文串。
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 输入: "race a car"
 * 输出: false
 */
class Solution
{

    /**
     * 思路
     *     如果是回文字符串，则比必满足 折半。也就是对折折后两字符串相同
     *     由于测试用例给的字符串并不是标准字符串，所以要先处理一遍
     *     由于使用了正则匹配，所以时间会消耗多
     * 耗时
     *     执行用时: 16 ms, 在Merge Sorted Array的PHP提交中击败了100.00% 的用户
     *     内存消耗: 14.7 MB, 在Merge Sorted Array的PHP提交中击败了100.00% 的用户
     * 时间复杂度 O(n)
     * 详解
     *     https://www.cnblogs.com/25-lH/p/9565117.html
     */
    function isPalindrome($s)
    {
        $length = strlen($s); // 处理特殊符号
        $str    = '';

        for ($i = 0; $i < $length; $i++) {
            if (preg_match('/[a-zA-Z0-9]/', $s[$i])) {
                $str .= strtolower($s[$i]);
            }
        }

        // 回文验证
        $length  = strlen($str);
        $middle = floor($length / 2);
        for ($i=0; $i < $middle; $i++) { 
            if ($str[$i] != $str[$length - $i - 1]) {
                return FALSE;
            }
        }
    
        return TRUE;
    }
}

$s = new Solution();
$result = $s->isPalindrome("A man, a plan, a canal: Panama");
var_dump($result);
