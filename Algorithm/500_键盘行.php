<?php
/**
 * 题目
 *     给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
 *     键盘不好画，详情看题目
 *         https://leetcode-cn.com/problems/keyboard-row/
 * 示例：
 *     输入: ["Hello", "Alaska", "Dad", "Peace"]
 *     输出: ["Alaska", "Dad"]
 * 注意：
 *     你可以重复使用键盘上同一字符。
 *     你可以假设输入的字符串将只包含字母。
 */
class Solution
{
    /**
     * 思路
     *     哈希表
     *     建立哈希表，进行查找
     * 耗时
     *     执行用时 : 16 ms, 在Keyboard Row的PHP提交中击败了80.00% 的用户
     *     内存消耗 : 14.7 MB, 在Keyboard Row的PHP提交中击败了100.00% 的用户
     */
    const DICT = [
        'Q' => 1,
        'W' => 1,
        'E' => 1,
        'R' => 1,
        'T' => 1,
        'Y' => 1,
        'U' => 1,
        'I' => 1,
        'O' => 1,
        'P' => 1,
        'A' => 2,
        'S' => 2,
        'D' => 2,
        'F' => 2,
        'G' => 2,
        'H' => 2,
        'J' => 2,
        'K' => 2,
        'L' => 2,
        'Z' => 3,
        'X' => 3,
        'C' => 3,
        'V' => 3,
        'B' => 3,
        'N' => 3,
        'M' => 3
    ];
    function findWords($words)
    {
        $outPut = [];

        if (!$words) {
            return $outPut;
        }

        foreach ($words as $w) {
            $tmp = strtoupper($w); // 临时处理字符串，使得符合哈希表规则
            $len = strlen($w);

            $flag = self::DICT[$tmp[0]];

            for ($i = 1; $i < $len; $i++) {
                if (self::DICT[$tmp[$i]] != $flag) {
                    goto skip;
                }
            }

            $outPut[] = $w;

            skip:
                continue;

        }
        return $outPut;
    }
}

$s = new Solution();
var_dump($s->findWords(["Hello", "Alaska", "Dad", "Peace"]));