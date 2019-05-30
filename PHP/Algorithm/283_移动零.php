<?php

/**
 * 题目
 *     给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 * 示例:
 *     输入: [0,1,0,3,12]
 *     输出: [1,3,12,0,0]
 * 说明:
 *     必须在原数组上操作，不能拷贝额外的数组。
 *     尽量减少操作次数。
 */
class Solution
{
    /**
     * 思路1
     *     碰到是0的元素，直接删除，并在数组末推入0
     *     站在PHP的角度上是可行的，但是站在数组的立场上，这样等于动态扩充了数组的长度
     * 耗时
     *     执行用时 : 96 ms, 在Move Zeroes的PHP提交中击败了27.78% 的用户
     *     内存消耗 : 16.2 MB, 在Move Zeroes的PHP提交中击败了12.07% 的用户
     * 复杂度
     *     O（n）
     */
    function moveZeroes1(&$nums) {

        $count = count($nums);

        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] == 0) {
                unset($nums[$i]);
                $nums[] = 0;
            }
        }
    }

    /**
     * 思路2
     *     将非0元素统一放在数组左边，之后再进行一次循环处理0元素
     *     比上一个方法时间和空间都快了不少，在PHP中，变量的销毁和扩充数组缺失需要消耗不少的资源
     * 耗时
     *     执行用时 : 36 ms, 在Move Zeroes的PHP提交中击败了74.44% 的用户
     *     内存消耗 : 15.9 MB, 在Move Zeroes的PHP提交中击败了58.62% 的用户
     * 复杂度
     *     O（n）
     */
    function moveZeroes2(&$nums)
    {
        $count = count($nums);
        $left  = 0;

        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] != 0) {
                $nums[$left] = $nums[$i];
                $left++;
            }
        }

        for ($i = $left; $i < $count; $i++) {
            $nums[$i] = 0;
        }
    }
}

$arr = [0,1,0,3,0,12];
$s = new Solution();
$s->moveZeroes($arr);
var_dump($arr);