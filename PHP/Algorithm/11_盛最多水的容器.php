<?php

/**
 * 题目
 *     具体题目请查看：https://leetcode-cn.com/problems/container-with-most-water/
 *     图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 * 示例 1:
 *     输入: [1,8,6,2,5,4,8,3,7]
 *     输出: 49
 */
class Solution
{
    /**
     * 思路
     *    双指针 (https://leetcode-cn.com/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode/)
     *    双指针，向中间计算面积，获取最大面积
     * 耗时
     *    执行用时 :40 ms, 在所有 PHP 提交中击败了96.81%的用户
     *    内存消耗 :16.2 MB, 在所有 PHP 提交中击败了78.82%的用户
     * 时间复杂度
     *    O（n）
     */
    function maxArea($height)
    {
    	$len      = count($height);
    	$hei      = $len-1; // 高指针
    	$low      = 0;      // 低指针
    	$area     = 0;

    	while ($low != $hei) {
    		
    		$countHeight = $height[$hei] > $height[$low] ? $height[$low] : $height[$hei];
    		$countWide   = $hei - $low;
    		$countArea   = $countHeight * $countWide;

    		// 和已经存在的面积作比较
    		$area = $countArea > $area ? $countArea : $area;

    		if ($height[$hei] >= $height[$low]) {
    			$low++;
    		}
    		else {
    			$hei--;
    		}
    	}

        return $area;
    }
}

$s = new Solution();
var_dump($s->maxArea([1,2,1]));

