/**
 * 题目
 *    给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *    你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 * 用例
 *     输入: nums = [2, 7, 11, 15], target = 9
 *     因为 nums[0] + nums[1] = 2 + 7 = 9
 *     输出: [0, 1]
 */
package main

import (
	"fmt"
)
/**
 * 思路
 *    GO 语言的第一道题，熟悉语法
 *    哈希map
 * 耗时
 *     执行用时 : 8 ms, 在Two Sum的Go提交中击败了96.53% 的用户
 *     内存消耗 : 3.9 MB, 在Two Sum的Go提交中击败了22.72% 的用户
 * 时间复杂度 
 *     O(n)
 */
func twoSum (nums []int, target int) []int {
	dict := make(map[int]int)

	for k,v := range nums {
		if inx, ok := dict[target-v]; ok {
			return []int{inx, k}
		}

		// map
		dict[v] = k
	}
	return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
	a := twoSum(nums, 9)
	fmt.Println("xxx", a)
}