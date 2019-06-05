package main

import "fmt"
/**
 * 题目
 *     假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *     每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 注意：给定 n 是一个正整数。
 * 示例 1：
 *     输入： 2
 *     输出： 2
 *     解释： 有两种方法可以爬到楼顶。
 *     1.  1 阶 + 1 阶
 *     2.  2 阶
 * 示例 2：
 *     输入： 3
 *     输出： 3
 *     解释： 有三种方法可以爬到楼顶。
 *     1.  1 阶 + 1 阶 + 1 阶
 *     2.  1 阶 + 2 阶
 *     3.  2 阶 + 1 阶
 */
func main () {
	c := climbStairs(10)
	fmt.Println(c)
}
/**
 * 思路
 *     动态规划
 *     更像是找规律的题目
 * 耗时
 *     执行用时 : 0 ms, 在Climbing Stairs的Go提交中击败了100.00% 的用户
 *     内存消耗 : 2 MB, 在Climbing Stairs的Go提交中击败了26.96% 的用户
 */
func climbStairs(n int) int {

	dict := make(map[int]int)

	dict[1] = 1 // 1阶梯只有一种方法
	dict[2] = 2 // 2阶梯只有两种方法

	for i := 3; i <= n; i++ {
		dict[i] = dict[i-1] + dict[i-2]
	}

	return dict[n]
}