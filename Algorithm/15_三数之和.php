<?php
/**
 * 题目
 *     给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
 * 注意
 *     答案中不可以包含重复的三元组。
 * 例如
 *     给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *     满足要求的三元组集合为：
 *     [
 *         [-1, 0, 1],
 *         [-1, -1, 2]
 *      ]
 */
class Solution
{
    /**
     * 思路
     *     大神思路
     *     先对数组排序，之后把三个数的整合问题变换成 左右两数合的问题
     *     更多的是对思想的改变
     * 耗时
     *     执行用时 : 544 ms, 在3Sum的PHP提交中击败了53.16% 的用户
     *     内存消耗 : 23.9 MB, 在3Sum的PHP提交中击败了83.87% 的用户
     * 时间复杂度
     *     O（n^3）
     */
    function threeSum($nums)
    {
        sort($nums);
        $count = count($nums);
        $dict  = [];
        $arr   = [];

        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] < 0) {
                for ($j = $i; $j < $count; $j++) {
                    if ($nums[$i] < 0 && $nums[$j+1] <= 0) {
                        $dict[abs($nums[$i]+$nums[$j+1])][] = [$nums[$i], $nums[$j+1]];
                        continue;
                    }
                    break;
                }
            }
            elseif ($nums[$i] > 0) {
                if ($dict[$nums[$i]]) {
                    foreach ($dict[$nums[$i]] as $d) {
                        $d[] = $nums[$i];
                        $arr[implode('/', $d)] = $d;
                    }
                }
            }
            else {
                $dict[0]++;
            }
        }

        if ($dict[0] >= 3) {
            $arr = [0,0,0];
        }

        return array_values($arr);
    }
}

$s = new Solution();
var_dump($s->threeSum([0,0,0]));