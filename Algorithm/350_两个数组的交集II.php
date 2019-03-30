<?php

/**
 * 题目
 *     给定两个数组，编写一个函数来计算它们的交集。
 * 示例 1:
 *     输入: nums1 = [1,2,2,1], nums2 = [2,2]
 *     输出: [2,2]
 * 示例 2:
 *     输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 *     输出: [4,9]
 * 说明：
 *     输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
 *     我们可以不考虑输出结果的顺序。
 * 进阶:
 *     如果给定的数组已经排好序呢？你将如何优化你的算法？
 *     如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 *     如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 */
class Solution
{
    /**
     * 思路
     *    哈希表
     *    mums1建立哈希表，并记录出现次数，在遍历mums2时进行哈希查找，输出相同
     * 耗时
     *     执行用时 : 52 ms, 在Intersection of Two Arrays II的PHP提交中击败了20.00% 的用户
     *     内存消耗 : 15 MB, 在Intersection of Two Arrays II的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    function intersect($nums1, $nums2)
    {
        $dict    = [];
        $aOutput = [];

        foreach ($nums1 as $n) {
            if ($dict[$n]) {
                $dict[$n]++;
                continue;
            }

            $dict[$n] = 1;
        }

        foreach ($nums2 as $n) {
            if ($dict[$n]) {
                $aOutput[] = $n;
                $dict[$n]--;
            }
        }

        return $aOutput;
    }
}

$s = new Solution();
var_dump($s->intersect([4,9,5], [9,4,9,8,4]));
