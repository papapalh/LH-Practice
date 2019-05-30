<?php
/**
 * 题目
 *     给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
 *     不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
 *     元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
 *     注意这五个元素可为任意顺序。
 * 用例
 *     给定 nums = [3,2,2,3], val = 3,
 *     函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
 *
 *     给定 nums = [0,1,2,2,3,0,4,2], val = 2,
 *     函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
 */

class Solution
{
    /**
     * 思路
     *    双指针(官方解法)
     *    通过快指针的循环数组，找到不同，由慢指针进行数组的替换，这样可以达到题目要求的(前几个元素的概念)
     *    和 26 类似
     * 耗时
     *    执行用时 : 16 ms, 在Remove Element的PHP提交中击败了95.74% 的用户
     *    内存消耗 : 14.6 MB, 在Remove Element的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *    O(n)
     */
    function removeElement(&$nums, $val)
    {
        $count = count($nums);
        $j = 0; // 慢指针

        // $i 快指针
        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] != $val) {
                $nums[$j] = $nums[$i];
                $j++;
            }
        }

        return $j;
    }
}

$nums = [0,1,2,2,3,0,4,2];
$val = 2;

$s = new Solution();
var_dump($s->removeElement($nums, $val));
var_dump($nums);