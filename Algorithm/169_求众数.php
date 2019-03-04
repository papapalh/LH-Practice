<?php
/**
 * 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 * 你可以假设数组是非空的，并且给定的数组总是存在众数。
 *     输入: [3,2,3]         输出: 3
 *     输入: [2,2,1,1,1,2,2] 输出: 2
 */
class Solution
{
    /**
     * HASH 算法
     *     执行用时: 232 ms, 在Majority Element的PHP提交中击败了0.00% 的用户
     *     内存消耗: 20.4 MB, 在Majority Element的PHP提交中击败了100.00% 的用户
     * 时间复杂度 O(n)
     * 同时间复杂度，但是因为额外维护了一个数组，所以会比 摩尔投票算法 差一些
     */
    function majorityElementByHash($nums)
    {
        $arrOutput = [];
        $target    = [];

        foreach ($nums as $n) {
            $arrOutput[$n] += 1;

            if ($arrOutput[$n] > $target['count']) {
                $target['count'] = $arrOutput[$n];
                $target['taget'] = $n;
            }
        }
        return $target['taget'];
    }

    /**
     * 摩尔投票算法
     *     执行用时: 96 ms, 在Majority Element的PHP提交中击败了46.15% 的用户
     *     内存消耗: 20.6 MB, 在Majority Element的PHP提交中击败了100.00% 的用户
     * 时间复杂度 O(n)
     * 算法详解
     *     https://www.cnblogs.com/25-lH/p/10469639.html
     */
    function majorityElementByVote($nums)
    {
        $major = $count = 0;

        foreach ($nums as $n) {
            if ($count == 0) {
                $major = $n;
                $count++;
            }
            else {
                $n == $major ? $count++ : $count--;
            }
        }

        return $major;
    }
}

$s = new Solution();
echo $s->majorityElementByHash([3,2,3]);
echo $s->majorityElementByVote([3,2,3]);