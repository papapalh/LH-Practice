<?php

/**
 * 题目
 *     给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
 * 示例 1:
 *     输入: [3,0,1]
 *     输出: 2
 * 示例 2:
 *     输入: [9,6,4,2,3,5,7,0,1]
 *     输出: 8
 * 说明:
 *     你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
 */
class Solution
{
    /**
     * 思路
     *    ^ 异或运算
     * 耗时
     *     执行用时 : 104 ms, 在Missing Number的PHP提交中击败了20.00% 的用户
     *     内存消耗 : 15.9 MB, 在Missing Number的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    function missingNumber($nums)
    {
        $sum = 0;
        $i   = 0;

        foreach ($nums as $n) {
            $sum ^= $n;
            $sum ^= $i;
            $i++;
        }

        $sum ^= $i;

        return $sum;
    }
}

$s = new Solution();
var_dump($s->missingNumber([3,0,1]));