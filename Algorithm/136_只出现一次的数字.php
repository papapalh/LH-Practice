<?php
/*
 * 题目：https://leetcode-cn.com/problems/single-number/
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 * 输入: [2,2,1]     输出: 1
 * 输入: [4,1,2,1,2] 输出: 4
 */

class Solution
{
    /* 
     * HASH 表做法
     * 用时 76 ms
     */
    function singleNumberHash($nums)
    {
        foreach ($nums as $n) {
            if (isset($arrOutput[(string)$n])) {
                unset($arrOutput[(string)$n]);
                continue;
            }
            $arrOutput[(string)$n] = $n;
        }

        return current($arrOutput);
    }

    /* 
     * 位运算
     * 用时 40 ms
     */
    function singleNumber($nums)
    {
        foreach ($nums as $n) {
            $sum ^= $n;
            var_dump($sum);
        }

        return $sum;
    }
}

$arr = [2,2,1];
$s = new Solution();
echo $s->singleNumberHash($arr);
echo $s->singleNumber($arr);
