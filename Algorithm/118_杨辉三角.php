<?php
/**
 * 题目
 *     给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *     在杨辉三角中，每个数是它左上方和右上方的数的和。
 *     杨辉三角是个动图，具体的看下链接
 *         https://leetcode-cn.com/problems/pascals-triangle/
 * 示例:
 *     输入: 5
 *     输出:
 *        [
 *             [1],
 *            [1,1],
 *           [1,2,1],
 *          [1,3,3,1],
 *         [1,4,6,4,1]
 *        ]
 */
class Solution
{
    /**
     * 思路
     *     没啥思路，就计算上层数据和，放入本层数据中
     * 耗时
     *     执行用时 : 12 ms, 在Pascal's Triangle的PHP提交中击败了98.55% 的用户
     *     内存消耗 : 14.3 MB, 在Pascal's Triangle的PHP提交中击败了100.00% 的用户
     */
    function generate($numRows)
    {
        $outPut = [];

        for ($i = 0; $i <= $numRows; $i++) {
            switch ($i) {
                case 0:
                    break;
                case 1:
                    $outPut[$i] = [1]; 
                    break;
                case 2:
                    $outPut[$i] = [1,1]; 
                    break;
                default:
                    $outPut[$i][] = 1;

                    $len = count($outPut[$i-1]) - 1;
                    for ($j = 0; $j < $len; $j++) {
                        $outPut[$i][] = $outPut[$i-1][$j] +  $outPut[$i-1][$j+1];
                    }

                    $outPut[$i][] = 1;
                break;
            }
        }

        return $outPut;
    }
}

$s = new Solution();
var_dump($s->generate(5));