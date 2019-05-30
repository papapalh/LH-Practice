<?php
/**
 * 题目
 *     给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *     你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 * 示例
 *     输入: [2,2,1]
 *     输出: 1
 *     输入: [4,1,2,1,2]
 *     输出: 4
 */
class Solution
{
    /**
     * 思路
     *     Hash 表
     *     使用 hash 表进行数字的判断和删除
     *     由于使用了额外空间，所以时间会消耗多
     * 耗时
     *     76 ms
     * 时间复杂度
     *     O(n)
     */
    function singleNumber1($nums)
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

    /**
     * 思路
     *     位运算
     *     直接操作 二进制位 有更好的效果
     * 耗时
     *     40 ms
     * 时间复杂度
     *     O(n)
     */
    function singleNumber($nums)
    {
        $sum = 0;

        foreach ($nums as $n) {
            $sum ^= $n;
        }

        return $sum;
    }
}

$arr = [2,2,1];
$s = new Solution();
echo $s->singleNumberHash($arr);
echo $s->singleNumber($arr);
