<?php

/**
 * 题目
 *     给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 * 示例 1:
 *     输入: "Let's take LeetCode contest"
 *     输出: "s'teL ekat edoCteeL tsetnoc"
 * 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 */
class Solution
{
    /**
     * 思路
     *     字符串操作
     *     对于不是空格的子串倒序处理
     *     遇到空格/最后一位时，将临时字符串赋予给最后的记过字符串
     * 耗时
     *     执行用时 : 56 ms, 在Reverse Words in a String III的PHP提交中击败了42.86% 的用户
     *     内存消耗 : 15 MB, 在Reverse Words in a String III的PHP提交中击败了83.33% 的用户
     * 时间复杂度
     *     O（n）
     */
    function reverseWords($s)
    {
        if (!$s) {
            return $s;
        }

        $len = strlen($s);
        $tmp = '';
        $str = '';

        for ($i = 0; $i < $len; $i++) {

            if ($s[$i] == ' ') {
                $str .=  $tmp . ' ';
                $tmp = '';
                continue;
            }

            $tmp = $s[$i] . $tmp;

            if ($i == $len-1) {
                $str .=  $tmp;
            }
        }

        return $str;

    }
}

$s = new Solution();
var_dump($s->reverseWords('Let\'s take LeetCode contest'));
