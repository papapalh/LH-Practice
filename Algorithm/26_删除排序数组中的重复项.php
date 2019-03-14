<?php
/**
 * 题目
 *     给定数组 nums = [1,1,2]
 *     函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
 *     你不需要考虑数组中超出新长度后面的元素。
 * 用例
 *     给定 nums = [0,0,1,1,1,2,2,3,3,4],
 *     函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
 *
 *     给定数组 nums = [1,1,2],
 *     函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
 */

class Solution
{
    /**
     * 思路1
     *    针对排序数组来说，前一个数与后一个数必然是相同的
     *    所以针对这个思路，如果发现前一个数与后一个数相同，则删除当前数，循环会得到一个唯一数组
     *    虽然也能通过，但是貌似和题目中的（前几个元素）的概念不想同
     * 耗时
     *    执行用时 : 68 ms, 在Remove Duplicates from Sorted Array的PHP提交中击败了44.37% 的用户
     *    内存消耗 : 16.8 MB, 在Remove Duplicates from Sorted Array的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *    O(n)
     */
    function removeDuplicates1(&$nums)
    {
        $count = count($nums);

        for ($i = 0; $i < $count; $i++) {
            if (isset($nums[$i+1])) {
                if ($nums[$i] != $nums[$i+1]) {
                    unset($nums[$i]);
                }
            }
        }

        return count($nums);
    }

    /**
     * 思路2
     *    双指针(官方解法)
     *    通过快指针的循环数组，找到不同，由慢指针进行数组的替换，这样可以达到题目要求的(前几个元素的概念)
     *    但是时间并没有快多少，只是更满足题目要求
     * 耗时
     *    执行用时 : 104 ms, 在Remove Duplicates from Sorted Array的PHP提交中击败了29.14% 的用户
     *    内存消耗 : 16.8 MB, 在Remove Duplicates from Sorted Array的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *    O(n)
     */
    function removeDuplicates(&$nums)
    {
        $count = count($nums);
        $j = 0; // 慢指针

        // $i 快指针
        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] != $nums[$j]) {
                $j++;
                $nums[$j] = $nums[$i];
            }
        }

        return $j+1;
    }
}

$s = new Solution();
$a = [0,0,1,1,1,2,2,3,3,4];
var_dump($s->removeDuplicates($a));
var_dump($a);
