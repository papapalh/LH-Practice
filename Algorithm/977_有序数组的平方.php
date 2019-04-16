<?php

/**
 * 题目
 *     给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
 * 示例 1：
 *     输入：[-4,-1,0,3,10]
 *     输出：[0,1,9,16,100]
 * 示例 2：
 *     输入：[-7,-3,2,3,11]
 *     输出：[4,9,9,49,121]
 */

class Solution {

    /**
     * 思路
     *     冒泡排序
     *     这可能是刷 Leet Cdoe 最不认真的一次了
     *     先重组，后冒泡排序
     *     果不其然，冒泡的时间复杂度直接超时
     *     后来改用快排做了一下，看题目的话使用双指针的，明天研究先吧。
     * 耗时
     *     快排耗时
     *     执行用时 : 192 ms, 在Squares of a Sorted Array的PHP提交中击败了52.94% 的用户
     *     内存消耗 : 17.8 MB, 在Squares of a Sorted Array的PHP提交中击败了18.97% 的用户
     */
    function sortedSquares($A)
    {
        $arr = [];
        $count = count($A);
        foreach ($A as $a) {
            $arr[] = $a*$a;
        }

        for ($i = 1; $i < $count; $i++) {
            for ($j = 0; $j < $count-$i; $j++) {
                if ($arr[$j] > $arr[$j+1]) {
                    list($arr[$j], $arr[$j+1]) = [$arr[$j+1], $arr[$j]];
                }
            }
        }

        return $arr;
    }
}

$s = new Solution();
var_dump($s->sortedSquares([-3,0,2]));
//[0,1,9,16,100]