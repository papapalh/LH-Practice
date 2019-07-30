<?php

/**
 * 题目
 *     将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
 *     比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
 *     L   C   I   R
 *     E T O E S I I G
 *     E   D   H   N
 *     之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
 *     请你实现这个将字符串进行指定行数变换的函数：
 *         string convert(string s, int numRows);
 * 示例 1:
 *     输入: s = "LEETCODEISHIRING", numRows = 3
 *     输出: "LCIRETOESIIGEDHN"
 * 示例 2:
 *     输入: s = "LEETCODEISHIRING", numRows = 4
 *     输出: "LDREOEIIECIHNTSG"
 *     解释:
 *         L     D     R
 *         E   O E   I I
 *         E C   I H   N
 *         T     S     G
 */
class Solution
{
    /**
     * 解析
     *     Z字排列，也就是循环去跑出字符串，放入对应的组内
     * 耗时
     *     执行用时 :20 ms, 在所有 PHP 提交中击败了92.50%的用户
     *     内存消耗 :15.2 MB, 在所有 PHP 提交中击败了25.76%的用户
     */
    function convert($s, $numRows)
    {
        $dict = [];

        $flag = true;

        $depth = 1;

        $len = strlen($s);
        for ($i = 0; $i < $len; $i++) {

            $dict[$depth][] = $s[$i];

            if ($depth >= $numRows) {
                $depth--;
                $flag = false;
                continue;
            }

            if ($depth == 1) {
                $depth++;
                $flag = true;
                continue;
            }

            if ($flag) {
                $depth++;
            }
            else {
                $depth--;
            }
        }

        $res = '';

        foreach ($dict as $d) {
            $res = $res . implode('', $d);
        }

        return $res;
    }
}

$s = new Solution();
var_dump($s->convert('LEETCODEISHIRING', 3));