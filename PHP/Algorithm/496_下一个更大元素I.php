<?php
/**
 * 题目
 *     给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
 *     找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
 *     nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。
 *     如果不存在，对应位置输出-1。
 * 示例 1:
 *     输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
 *     输出: [-1,3,-1]
 *     解释:
 *         对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
 *         对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
 *         对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
 * 示例 2:
 *     输入: nums1 = [2,4], nums2 = [1,2,3,4].
 *     输出: [3,-1]
 *     解释:
 *         对于num1中的数字2，第二个数组中的下一个较大数字是3。
 *         对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。
 *     注意:
 *         nums1和nums2中所有元素是唯一的。
 *         nums1和nums2 的数组大小都不超过1000。
 */
class Solution
{
    /**
     * 思路
     *     哈希表
     *     使用 n(n log n) 时间计算每个数组元素中下一个最大的数，建立哈希映射关系
     *     在遍历查找数组，通过映射关系，实现对最大数的查找
     * 耗时
     *     执行用时 : 20 ms, 在Next Greater Element I的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.9 MB, 在Next Greater Element I的PHP提交中击败了44.44% 的用户
     */
    function nextGreaterElement($nums1, $nums2)
    {
        $dict   = [];
        $outPut = [];

        // 创建字典
        $len = count($nums2);
        for ($i = 0; $i < $len; $i++) {
            for ($j = $i; $j < $len; $j++) {
                if ($nums2[$i] < $nums2[$j]) {
                    $dict[$nums2[$i]] = $nums2[$j];
                    break;
                }
            }
        }

        foreach ($nums1 as $n) {
            if (isset($dict[$n])) {
                $outPut[] = $dict[$n];
            }
            else {
                $outPut[] = -1;
            }
        }

       return $outPut;
    }
}

$s = new Solution();
var_dump($s->nextGreaterElement([1,3,5,2,4], [6,5,4,3,2,1,7]));