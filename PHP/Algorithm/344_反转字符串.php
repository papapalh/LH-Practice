<?php
/**
 * 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
 * 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
 * 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
 *    输入：["h","e","l","l","o"]
 *    输出：["o","l","l","e","h"]
 *    输入：["H","a","n","n","a","h"]
 *    输出：["h","a","n","n","a","H"]
 */
class Solution
{
    /**
     * 思路1
     *    新开辟一个数组，让原数组倒叙输出，则完成反转字符效果
     * 时间复杂度
     *    O(n)
     * 耗时
     *    执行用时: 96 ms, 在Reverse String的PHP提交中击败了77.27% 的用户
     *    内存消耗: 35.4 MB, 在Reverse String的PHP提交中击败了100.00% 的用户
     */
    function reverseString1(&$s)
    {
        $count  = count($s) - 1;
        $result = []; 

        for ($i = $count; $i >= 0; $i--) {
            $result[] = $s[$i];
        }

        $s = $result;
    }

    /**
     * 思路2
     *    在原数组上操作
     *    在折半的位置上，替换收尾字符串，只循环 n/2 次便可以达到效果
     * 时间复杂度
     *    O(n)
     * 耗时
     *    执行用时: 124 ms, 在Reverse String的PHP提交中击败了63.64% 的用户
     *    内存消耗: 35.3 MB, 在Reverse String的PHP提交中击败了100.00% 的用户
     */
    function reverseString(&$s)
    {
        $length = count($s);
        $count  = floor($length / 2);

        for ($i = 0; $i < $count; $i++) {
            $tmp                 = $s[$i];
            $s[$i]               = $s[$length - 1 - $i];
            $s[$length - 1 - $i] = $tmp;
        }
    }
}


$s = ["h","e","l","l","o"];
$Object = new Solution();
$Object->reverseString($s);
var_dump($s);
