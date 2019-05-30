<?php
/**
 * 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
 *     字符          数值
 *     I             1
 *     V             5
 *     X             10
 *     L             50
 *     C             100
 *     D             500
 *     M             1000
 * 示例
 *     输入: "III"
 *     输出: 3
 *     输入: "IV"
 *     输出: 4
 *     输入: "IX"
 *     输出: 9
 *     输入: "LVIII"
 *     输出: 58
 *     解释: L = 50, V= 5, III = 3.
 *     输入: "MCMXCIV"
 *     输出: 1994
 *     解释: M = 1000, CM = 900, XC = 90, IV = 4.
 */
class Solution
{
    /**
     * 思路
     *    根据实例即可找到规律，发现后一位字符代表的数目大于前一位，则减，否则加
     *    记得考虑等于的情况
     * 耗时
     *     执行用时 : 88 ms, 在Roman to Integer的PHP提交中击败了27.40% 的用户
     *     内存消耗 : 14.9 MB, 在Reverse Integer的PHP提交中击败了100.00% 的用
     * 时间复杂度
     *     O(n)
     */
    function romanToInt($s)
    {
        $hash = [
            'I' => 1,
            'V' => 5,
            'X' => 10,
            'L' => 50,
            'C' => 100,
            'D' => 500,
            'M' => 1000,
        ];

        $count = strlen($s);
        $num   = 0;

        for ($i = 0; $i < $count; $i++) {

            if ($hash[$s[$i]] >= $hash[$s[$i+1]]) {
                $num += $hash[$s[$i]];
            }
            else {
                $num -= $hash[$s[$i]];
            }

        }

        return $num;
    }
}

$s = new Solution();
echo $s->romanToInt('III');