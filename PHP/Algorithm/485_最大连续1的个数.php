<?php
/**
 * 题目
 *     给定一个二进制数组， 计算其中最大连续1的个数。
 * 示例 1:
 *     输入: [1,1,0,1,1,1]
 *     输出: 3
 *     解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
 * 注意：
 *     输入的数组只包含 0 和1。
 *     输入数组的长度是正整数，且不超过 10,000。
 */
class Solution
{
    /**
     * 思路
     *     通过O(n)的循环，记录出现1的最大次数
     *     最后和总数比较，得到最大的1
     * 耗时
     *     执行用时 : 152 ms, 在Max Consecutive Ones的PHP提交中击败了78.95% 的用户
     *     内存消耗 : 15.7 MB, 在Max Consecutive Ones的PHP提交中击败了81.82% 的用户
     * 时间复杂度
     *     O(n)
     */
    function findMaxConsecutiveOnes($nums) {
        
        $sum = 0;
        $tmp = 0;

        foreach ($nums as $n) {
            // 如果为1，则记录次数
            if (1&$n) {
                $tmp++;
                if ($tmp > $sum) {
                    $sum = $tmp;
                }
            }
            else {
                $tmp = 0;
            }
        }

        return $sum;
    }
}

$s = new Solution();
var_dump($s->findMaxConsecutiveOnes([1,1,0,1,1,1]));