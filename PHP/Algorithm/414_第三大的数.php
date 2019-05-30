<?php
/**
 * 题目
 *     给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
 * 示例 1:
 *     输入: [3, 2, 1]
 *     输出: 1
 *     解释: 第三大的数是 1.
 * 示例 2:
 *     输入: [1, 2]
 *     输出: 2
 *     解释: 第三大的数不存在, 所以返回最大的数 2 .
 * 示例 3:
 *     输入: [2, 2, 3, 1]
 *     输出: 1
 *     解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。存在两个值为2的数，它们都排第二。
 */
class Solution
{
    /**
     * 思路
     *     队列
     *     类似队列，维护一个队列，当元素队列需要改变时候，移动这个队列
     *     最后输出123大的数
     * 耗时
     *     执行用时 : 36 ms, 在Third Maximum Number的PHP提交中击败了77.78% 的用户
     *     内存消耗 : 16.1 MB, 在Third Maximum Number的PHP提交中击败了66.67% 的用户
     * 时间复杂
     *     O（n）
     */
    function thirdMax($nums)
    {
        $dict = [];

        foreach ($nums as $n) {

            // 对于相等的，不处理
            if ($n === $dict[1] || $n === $dict[2] || $n === $dict[3]) {
                continue;
            }

            if ($n > $dict[1]) {
                $dict[3] = $dict[2];
                $dict[2] = $dict[1];
                $dict[1] = $n;
            }
            elseif ($n > $dict[2]) {
                $dict[3] = $dict[2];
                $dict[2] = $n;
            }
            elseif ($n >= $dict[3]) {
                $dict[3] = $n;
            }
        }

        return isset($dict[3]) ? $dict[3] : $dict[1];
    }
}

$s = new Solution();
var_dump($s->thirdMax([3, 1, 0]));