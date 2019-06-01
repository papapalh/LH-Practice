package main

import (
	"fmt"
)
/**
 * 题目
 *     给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
 *     如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *     你可以假设数组中无重复元素。
 * 用例
 *     输入: [1,3,5,6], 5
 *     输出: 2
 *     输入: [1,3,5,6], 2
 *     输出: 1
 *     输入: [1,3,5,6], 7
 *     输出: 4
 */
func main() {
	arr := []int{1,3,5,6}
	x := searchInsert(arr, 2)
	fmt.Println("=",x)
}
/**
 * 思路
 *     二分查找，确定区间
 * 耗时
 *     执行用时 : 8 ms, 在Search Insert Position的Go提交中击败了91.37% 的用户
 *     内存消耗 : 3.2 MB, 在Search Insert Position的Go提交中击败了30.95% 的用户
 */
func searchInsert(nums []int, target int) int {

	min := 0
	max := len(nums)
	mid := 0

	// 个数少的情况

	for min != max  {

		mid = (max + min) / 2

		if target > nums[mid] {
			min = mid + 1
		} else if target < nums[mid] {
			max = mid
		} else {
			return mid
		}
	}

	return min
}