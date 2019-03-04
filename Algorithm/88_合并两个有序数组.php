<?php
/**
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 * 说明
 *     初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 *     你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 * 输入 nums1 = [1,2,3,0,0,0], m = 3,nums2 = [2,5,6], n = 3
 * 输出: [1,2,2,3,5,6]
 */
class Solution
{
    /**
     * 归并排序 算法
     *     执行用时: 16 ms, 在Merge Sorted Array的PHP提交中击败了100.00% 的用户
     *     内存消耗: 14.7 MB, 在Merge Sorted Array的PHP提交中击败了100.00% 的用户
     * 时间复杂度 O(n)
     * 算法详解
     *     https://www.cnblogs.com/25-lH/p/10471966.html
     */
    function merge(&$nums1, $m, $nums2, $n)
    {
        $count = $m + $n; // 循环次数
        $n1    = $n2 = 0; // n1 n2 数组偏移量
        $arr   = [];      // 输出

        // 因为是 PHP ，所以切除数组长度之外的无用元素
        $nums1 = array_slice($nums1,0,$m);

        for ($i = 1; $i <= $count; $i++) {
            // 排除空元素的影响
            if (!isset($nums1[$n1])) {
                $arr[] = $nums2[$n2];
                $n2++;
                continue;
            }
            if (!isset($nums2[$n2])) {
                $arr[] = $nums1[$n1];
                $n1++;
                continue;
            }

            // 比较数组头
            if ($nums1[$n1] < $nums2[$n2]) {
                $arr[] = $nums1[$n1];
                $n1++;
            }
            elseif ($nums1[$n1] == $nums2[$n2]) {
                $arr[] = $nums1[$n1];
                $arr[] = $nums2[$n2];
                $n1++;
                $n2++;
                $i++;
            }
            else {
                $arr[] = $nums2[$n2];
                $n2++;
            }
        }

        $nums1 = $arr; 
    }
}

$nums1 = [-1,0,0,3,3,3,0,0,0];
$m = 6;
$nums2 = [1,2,2];
$n = 3;
$s = new Solution();
$s->merge($nums1, $m, $nums2, $n);

var_dump($nums1);