<?php
/**
 * 题目
 *     给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *     如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
 *     注意你不能在买入股票前卖出股票。
 * 示例 1:
 *     输入: [7,1,5,3,6,4]
 *     输出: 5
 *     解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 *     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 * 示例 2:
 *     输入: [7,6,4,3,1]
 *     输出: 0
 *     解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 */

class Solution
{
    /**
     * 思路
     *     记录每一次当股票价格达到低点时，到达最高点的差价(推入)
     * 耗时
     *     执行用时 : 64 ms, 在Best Time to Buy and Sell Stock的PHP提交中击败了44.44% 的用户
     *     内存消耗 : 16.6 MB, 在Best Time to Buy and Sell Stock的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    function maxProfit($prices)
    {

        if (!$prices) {
            return 0;
        }

        $low       = $prices[0];
        $hei       = 0;
        $arrOutput = [];

        foreach ($prices as $p) {

            if ($p <= $low) {
                $low = $p;
                $hei = $p;
            }

            if ($p >= $hei) {
                $hei = $p;
            }

            $arrOutput[] = $hei - $low;
        }

        return max($arrOutput);
    }
}

$prices = [7,1,5,3,6,4,7];
$s = new Solution();
$s->maxProfit($prices);