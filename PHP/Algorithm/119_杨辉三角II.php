<?php
/**
 * 题目
 *     给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *     杨辉三角是个动图，具体的看下链接
 *         https://leetcode-cn.com/problems/pascals-triangle/
 * 示例:
 *     输入: 3
 *     输出: [1,3,3,1]
 */
class Solution
{
    /**
     * 思路
     *     和杨辉三角I类似
     *     没啥思路，就计算上层数据和，放入本层数据中
     * 耗时
     *     执行用时 : 20 ms, 在Pascal's Triangle II的PHP提交中击败了70.00% 的用户
     *     内存消耗 : 14.7 MB, 在Pascal's Triangle II的PHP提交中击败了56.25% 的用户
     */
    function getRow($numRows)
    {
        $outPut = [];

        for ($i = 0; $i <= $numRows+1; $i++) {
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

            if ($i-1 == $numRows) {
                return $outPut[$i];
            }
        }
    }
}

$s = new Solution();
var_dump($s->getRow(3));