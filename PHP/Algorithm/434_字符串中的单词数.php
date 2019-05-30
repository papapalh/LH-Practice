<?php
/**
 * 题目
 *     统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 *     请注意，你可以假定字符串里不包括任何不可打印的字符。
 * 示例:
 *     输入: "Hello, my name is John"
 *     输出: 5
 */
class Solution
{
    /**
     * 思路
     *     哈希表
     *     记录每个单词
     *     由于额外开辟了空间，内存消耗过大
     * 耗时
     *     执行用时 : 12 ms, 在Number of Segments in a String的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.8 MB, 在Number of Segments in a String的PHP提交中击败了25.00% 的用户
     * 时间复杂度
     *     O（n）
     */
    function countSegments($s) {
        $len   = strlen($s);
        $dict  = [];
        $count = 0;

        if ($len <= 0) {
            return 0;
        }

        for ($i = 0; $i < $len; $i++) {
            if ($s[$i] == ' ') {
                $count++;
                continue;
            }

            $dict[$count] .= $s[$i];
        }

        return count($dict);
    }
}

$s = new Solution();
var_dump($s->countSegments('Hello, my name is John'));