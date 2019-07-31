<?php
/**
 * 题目
 *     给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *     给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 * 示例:
 *     输入："23"
 *     输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 说明:
 *     尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 */

/**
 * 思路
 *     类似树的全路径
 *     构造一个类似树的递归，一层一层的获得全排列
 * 耗时
 *     执行用时 :16 ms, 在所有 PHP 提交中击败了24.05%的用户
 *     内存消耗 :14.9 MB, 在所有 PHP 提交中击败了13.79%的用户
 */
class Solution
{
    public static $dict = [
        0 => "",
        1 => '*',
        2 => 'abc',
        3 => 'def',
        4 => 'ghi',
        5 => 'jkl',
        6 => 'mno',
        7 => 'pqrs',
        8 => 'tuv',
        9 => 'wxyz',
    ];

    public $res = [];

    function letterCombinations($digits)
    {
        if (!$digits) {
            return [];
        }

        $digits = str_split($digits);

        $this->x($digits);

        return $this->res;
    }

    function x($digits, $path = '')
    {
        $d = self::$dict[array_shift($digits)];
        if (!$d) {
            $this->res[] = $path;
            return false;
        }

        foreach (str_split($d) as $v) {
            $this->x($digits, $path.$v);
        }
    }
}

$s = new Solution();

var_dump($s->letterCombinations('23'));