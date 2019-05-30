<?php
/**
 * 题目
 *     给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
 * 示例：
 *     给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
 *     sumRange(0, 2) -> 1
 *     sumRange(2, 5) -> -1
 *     sumRange(0, 5) -> -3
 * 说明:
 *     你可以假设数组不可变。
 *     会多次调用 sumRange 方法。
 */
class NumArray
{
    /**
     * 思路
     *     循环计算
     *     不计算每次的结果，在计算时候直接计算开始和结尾
     *     问题是需要做大量重复计算，时间耗时太长,不是一个好方法
     * 耗时
     *     执行用时 : 4268 ms, 在Range Sum Query - Immutable的PHP提交中击败了5.88% 的用户
     *     内存消耗 : 21.9 MB, 在Range Sum Query - Immutable的PHP提交中击败了10.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    public static $nums = [];
    function __construct($nums) {
        self::$nums = $nums;
    }
    function sumRange($i, $j)
    {
        $sum = 0;

        // 如果小于，等于参数错误
        if ($i > $j) {
            return 0;
        }

        // 如果等于，则输出当前位置和就可以
        if ($i == $j) {
            return self::$nums[$i];
        }

        while ($i <= $j) {
            $sum += self::$nums[$i];
            $i++;
        }

        return $sum;
    }
}

class NumArray1
{
    /**
     * 思路
     *     哈希表
     *     使用O（n^2）的时间处理可能存在的所有的集合，最后通过哈希表查找处理
     *     问题是O (n^2)的处理时间导致超时
     */
    public static $dict = [];
    function __construct($nums)
    {
        $len = count($nums);

        for ($i = 0; $i < $len; $i++) {

            self::$dict[$i] = $nums[$i];
            $sum = $nums[$i];

            for ($j = $i+1; $j < $len; $j++) {
                $sum += $nums[$j];
                self::$dict[$i.$j] = $sum;
            }
        }
    }
    function sumRange($i, $j)
    {
        // 如果小于，等于参数错误
        if ($i > $j) {
            return 0;
        }

        // 如果等于，则输出当前位置和就可以
        if ($i == $j) {
            return self::$dict[$i];
        }

        return self::$dict[$i.$j];
    }
}

class NumArray2
{
    /**
     * 思路
     *     动态规划，哈希表，存储已经算过的路径
     *     大神思想
     *     使用O(n)时间处理数组，计算每个数组的和
     *     在返回时，减去对应的数值
     */
    public static $sums = [];
    function __construct($nums)
    {
        $len = count($nums);

        if ($len <= 0) {
            return 0;
        }

        self::$sums[0] = $nums[0];

        for ($i = 1; $i < $len; $i++) {
            self::$sums[$i] += self::$sums[$i - 1] + $nums[$i];
        }
    }
    function sumRange($i, $j)
    {
        if ($i == 0) {
            return self::$sums[$j];
        } else {
            return (self::$sums[$j] - self::$sums[$i - 1]);
        }
    }
}