package main

import "fmt"
/*
 * 题目
 *     给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 示例:
 *     输入: [-2,1,-3,4,-1,2,1,-5,4],
 *     输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 */
func main() {

	arr := []int{-2147483647}
	fmt.Printf("x", arr)

	x := maxSubArray(arr)

	fmt.Printf("x", x)
}
/*
 * 思路
 *     动态规划
 *     其实我感觉这种题更像是总结规律的题目
 * 耗时
 *     执行用时 : 12 ms, 在Maximum Subarray的Go提交中击败了64.13% 的用户
 *     内存消耗 : 3.3 MB, 在Maximum Subarray的Go提交中击败了73.67% 的用户
 * 时间复杂度
 *     O（n）
 */
func maxSubArray(nums []int) int {

	res := -2147483647
	sum := 0

	for _,v := range nums {
		sum = Max(sum + v, v)
		res = Max(sum, res)
	}

	return res
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}