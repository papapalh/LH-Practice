<?php

/**
 * 题目
 *    给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *    如果不存在最后一个单词，请返回 0 。
 *    说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 * 用例
 *    输入: Hello World
 *    输出: 5
 */
class Solution
{
    /**
     * 思路
     *     双指针
     *     前置指针+后置指针 存储有效字符串
     *     这样，后置 - 前置 可以得到有效的字符串/数量
     * 耗时
     *     执行用时 : 16 ms, 在Length of Last Word的PHP提交中击败了55.56% 的用户
     *     内存消耗 : 14.8 MB, 在Length of Last Word的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    function lengthOfLastWord($s)
    {
        $low    = 0;  // 慢指针
        $hei    = 0;  // 快指针
        $length = strlen($s);

        for ($i = 0; $i < $length; $i++) {

            // 如果不是空字符串，则推入快指针
            if ($s[$i] != ' ') {
                // 由于是从 0 开始循环,每次加一，用作计算
                $hei = $i+1;
            }
            // 否则看下一个指针是不是有字符，如果有，则推入指针
            else {
                // 判断，不要超限
                if ($s[$i+1] != ' ' && ($i+1) <= $length) {
                    $low = $i+1;
                }
            }
        }

        return $hei - $low;
    }
}

$s = new Solution();
var_dump($s->lengthOfLastWord('a'));