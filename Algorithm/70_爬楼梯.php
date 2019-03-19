<?php

/**
 * 题目
 *    假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *    每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *    注意：给定 n 是一个正整数。
 * 用例
 *    输入： 2
 *    输出： 2
 *    解释： 有两种方法可以爬到楼顶。
 *        1.  1 阶 + 1 阶
 *        2.  2 阶
 *    输入： 3
 *    输出： 3
 *    解释： 有三种方法可以爬到楼顶。
 *        1.  1 阶 + 1 阶 + 1 阶
 *        2.  1 阶 + 2 阶
 *        3.  2 阶 + 1 阶
 */
class Solution
{
    public static $dict = [];

    /**
     * 思路1
     *     动态规划
     *     记录上一次爬楼梯的转态，只计算新楼层
     *     需要有字典表记录爬楼状态
     * 详解
     *     爬楼梯的行为，必然是从 1 层爬起，而且也只有两种方法 2阶/1阶
     *         1 层 = 1(一种方法)
     *         2 层 = 2(两种方法)
     *     从第3层开始
     *         3 层 = (二层方法数目) + (一层方法数目)
     *         4 层 = (三层方法数目) + (二层方法数目)
     *         ...
     *         $n = $n-1 + $n-2
     * 耗时
     *     执行用时 : 16 ms, 在Climbing Stairs的PHP提交中击败了45.45% 的用户
     *     内存消耗 : 14.8 MB, 在Climbing Stairs的PHP提交中击败了100.00% 的用户
     */
    function climbStairs ($n)
    {
        // 从第三阶开始循环
        self::$dict[1] = 1;
        self::$dict[2] = 2;

        $i = 3;
        while ($n >= $i) {
            self::$dict[$i] = self::$dict[$i-1] + self::$dict[$i-2];
            $i++;
        }

        return self::$dict[$n];
    }

    /**
     * 思路
     *     递归
     *     原理同上，但是递归更耗资源，所以更慢
     * 耗时
     *     执行用时 : 24 ms, 在Climbing Stairs的PHP提交中击败了18.18% 的用户
     *     内存消耗 : 14.7 MB, 在Climbing Stairs的PHP提交中击败了100.00% 的用户
     */
    function climbStairs1($n)
    {
        if (self::$dict[$n]) {
            return self::$dict[$n];
        }

        if ($n == 0 || $n == 1) {
            return 1;
        }
        else {
            self::$dict[$n] = $this->climbStairs($n - 1) + $this->climbStairs($n - 2);
        }

        return self::$dict[$n];
    }
}

$s = new Solution();
var_dump($s->climbStairs(3));