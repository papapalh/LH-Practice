<?php

/**
 * 题目
 *    给定两个二进制字符串，返回他们的和（用二进制表示）。
 *    输入为非空字符串且只包含数字 1 和 0。
 * 实例
 *    输入: a = "11", b = "1"
 *    输出: "100"
 *    输入: a = "1010", b = "1011"
 *    输出: "10101"
 */
class Solution
{
    /**
     * 思路1
     *     转换数字方法
     *     先把2进制转换为10进制，求和，在转换为2进制
     *     问题
     *         数字溢出，测试用例失败
     * 耗时
     *     无
     */
    function addBinary1($a, $b)
    {
        $sum     = 0;
        $binary  = '';
        $aLength = strlen($a);
        $bLength = strlen($b);

        $count  = 1;
        for ($i = $aLength-1; $i >= 0; $i--) {
            if ($a[$i]) {
                $sum += $count;
            }

            $count *= 2;
        }

        $count = 1;
        for ($i = $bLength-1; $i >= 0; $i--) {
            if ($b[$i]) {
                $sum += $count;
            }
            $count *= 2;
        }

        if (!$sum) {
            return 0;
        }

        // $sum 为 ab 两数十进制总和
        while ($sum >= 1) {

            $binary = $sum % 2 . $binary;

            $sum = floor($sum / 2);
        }

        return $binary;
    }

    /**
     * 思路2
     *     直接操作二进制字符串
     *     先把字符串补0，预2进1，求字符串
     * 耗时
     *     执行用时 : 36 ms, 在Add Binary的PHP提交中击败了12.50% 的用户
     *     内存消耗 : 14.9 MB, 在Add Binary的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    function addBinary($a, $b)
    {
        $binary  = '';
        $aLength = strlen($a);
        $bLength = strlen($b);
        $length = $aLength > $bLength ? $aLength-1 : $bLength-1;

        // 对于 $a $b 字符不足的,补0处理
        $a = str_pad($a,$length+1,"0",STR_PAD_LEFT);
        $b = str_pad($b,$length+1,"0",STR_PAD_LEFT);

        // 处理位字符串 - 进位
        for ($i = $length; $i >= 0; $i--) {

            $binary[$i] = (int)$a[$i] + (int)$b[$i] + (int)$binary[$i];

            switch ($binary[$i]) {
                case 0:
                    continue;
                    break;
                case 1:
                    continue;
                    break;
                case 2:
                    $binary[$i] = '0';
                    if ($i == 0) {
                        $binary = '1'.$binary;
                    }
                    else {
                        $binary[$i-1] = (int)$binary[$i-1] + 1;
                    }
                    break;
                case 3:
                    $binary[$i] = '1';
                    if ($i == 0) {
                        $binary = '1'.$binary;
                    }
                    else {
                        $binary[$i-1] = (int)$binary[$i-1] + 1;
                    }
                    break;
            }
        }

        return $binary;
    }
}

$s = new Solution();

$a = "11";
$b = '1';
var_dump($s->addBinary($a, $b));