<?php
/**
 * 题目
 *      给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 *      函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 * 说明:
 *      返回的下标值（index1 和 index2）不是从零开始的。
 *      你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
 * 示例:
 *      输入: numbers = [2, 7, 11, 15], target = 9
 *      输出: [1,2]
 *      解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 */
class Solution
{
    /**
     * 思路
     *     哈希表做法
     *     先遍历一遍数组，建立 哈希表 , 对于相同建，放入数组的顺序
     *     搜索时先确定等于 $target 的值是否存在，在判断是否有重复
     *     缺点在于当数组过大时，占用空间，和遍历复杂度过大，导致超时
     */
    function twoSum1($numbers, $target)
    {
        $dict      = [];
        $i = 1;

        // 建立 哈希表
        foreach ($numbers as $n) {
            $dict[$n][] = $i;
            $i++;
        }

        foreach ($numbers as $n) {

            // 如果相等
            if ($n == ($target-$n)) {
                if ($dict[$n][0] && $dict[$n][1]) {
                    return [$dict[$n][0], $dict[$n][1]];
                }
            }

            // 如果存在匹配
            if ($dict[$n] && $dict[$target-$n]) {
                return [$dict[$n][0], $dict[$target-$n][0]];
            }
        }

        return [];
    }

    /**
     * 思路
     *     双指针
     *     题目要求描述，有序数组， index1 必大于 index2
     *     赋予高低指针，一直向中间推进，如果符合则返回，不符合则退出
     * 耗时
     *     执行用时 : 312 ms, 在Two Sum II - Input array is sorted的PHP提交中击败了46.67% 的用户
     *     内存消耗 : 16.1 MB, 在Two Sum II - Input array is sorted的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O（n）
     */
    function twoSum($numbers, $target)
    {
        $low = 0;
        $hei = count($numbers)-1;

        while($low < $hei) {
            if ($numbers[$low] + $numbers[$hei] == $target) {
                return [$low+1, $hei+1];
            }
            elseif ($numbers[$low] + $numbers[$hei] < $target) {
                $low++;
            }
            else {
                $hei--;
            }
        }

        return [];
    }
}

$s = new Solution();
var_dump($s->twoSum([2, 7, 11, 15], 9));