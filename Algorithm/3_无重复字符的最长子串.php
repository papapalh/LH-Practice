<?php
/**
 * 题目
 *     给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 示例 1:
 *     输入: "abcabcbb"
 *     输出: 3 
 *     解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 示例 2:
 *     输入: "bbbbb"
 *     输出: 1
 *     解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 示例 3:
 *     输入: "pwwkew"
 *     输出: 3
 *     解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 */
class Solution
{
    /**
     * 思路
     *     滑动窗口
     *         详解请看力扣  https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/
     *     详解
     *         例如 abcabcbb
     *             1 : 小窗口到 abc 没有重复,小窗继续，并记录次数
     *             2 : 小窗到 a 时，发现重复，则对于 小窗口(list),出队列一个元素，使得整个小窗不出现重复元素
     *             3 : 目前小窗为 bca 
     *             4 : ......重复之前的操作
     *             5 : 登出过程中小窗的最大值
     * 耗时
     *     执行用时 : 44 ms, 在Longest Substring Without Repeating Characters的PHP提交中击败了81.76% 的用户
     *     内存消耗 : 15 MB, 在Longest Substring Without Repeating Characters的PHP提交中击败了31.06% 的用户
     */
    function lengthOfLongestSubstring($s)
    {
        $len = strlen($s);
        $list = [];
        $m = 0;

        for ($i = 0; $i < $len; $i++) {
 
            if (in_array($s[$i], $list)) {
                
                $count = count($list);
                $count > $m ? $m = $count : '';

                while (in_array($s[$i], $list)) {
                    array_shift($list);
                    $max--;
                }

                $list[] = $s[$i];
                $max++;
                continue;
            }

            $list[] = $s[$i];
            $max++;
        }

        return max($m, $max);
    }
}

class Solution1 
{
    /**
     * 思路
     *     暴力破解，把字符串的每一个字符当做不重复的开端
     *     O（n ^ 2） 时间查询
     * 耗时
     *     时间超时
     * 时间复杂度
     *     O（n ^ 2）
     */
    function lengthOfLongestSubstring($s)
    {
        $len  = strlen($s);
        $sum  = 0;

        for ($i = 0; $i < $len; $i++) {

            $dict = [];
            $max  = 0;
            for ($j = $i; $j < $len; $j++) {

                if ($dict[$s[$j]]) {
                    $max > $m ? $m = $max : '';
                    break;
                }

                $dict[$s[$j]] = true;
                $max++;

                $max > $m ? $m = $max : '';
            }
        }

        return max($m, $sum);
    }
}