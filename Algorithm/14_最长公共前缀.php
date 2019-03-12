<?php
/**
 * 题目
 *     编写一个函数来查找字符串数组中的最长公共前缀。
 *     如果不存在公共前缀，返回空字符串 ""。
 * 用例
 *     输入: ["flower","flow","flight"]
 *     输出: "fl"
 *     输入: ["dog","racecar","car"]
 *     输出: ""
 *     解释: 输入不存在公共前缀。
 */
class Solution
{
    /**
     * 水平扫描法
     *     遍历 S1-Sn 中所有字符，找到最长的前缀，如果不满足，则算法退出
     * 思路
     *     前缀必须是每个单词都有的，所以取出第一个单词作为标准
     *     遍历剩下的数组元素作为比较
     * 耗时
     *     执行用时 : 16 ms, 在Longest Common Prefix的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.9 MB, 在Longest Common Prefix的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     时间复杂度：O(S)，S 是所有字符串中字符数量的总和。
     */
    function longestCommonPrefix($strs)
    {

        $flag   = array_shift($strs); // 取出第一个数
        $length = strlen($flag);      // 剩下的数组长度
        $str    = '';                 // 输出字符串

        for ($i = 0; $i < $length; $i++) {
            foreach ($strs as $s) {
                if ($s[$i] != $flag[$i]) {
                    goto no;
                }
            }

            $str .= $flag[$i];
        }

        no:
            return $str;
    }
}

$s = new Solution();
echo $s->longestCommonPrefix(["aca","cba"]);