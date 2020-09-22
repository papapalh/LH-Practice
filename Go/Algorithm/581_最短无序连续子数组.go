package main

import "fmt"

func main() {

	fmt.Print(findUnsortedSubarray([]int{2, 3, 3, 2, 4}))
}

/**
 * 参考官方题解
 *    https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/solution/zui-duan-wu-xu-lian-xu-zi-shu-zu-by-leetcode/
 * 思路1
 *     先排序，在比较。
 *         [2, 6, 4, 8, 10, 9, 15] 排序后 [2, 4, 6, 8, 9, 10, 15]
 *         在比较不同
 * 思路二
 *     确定左右边界，循环获取边界值(大小)，再次确定边界
 * 耗时
 *     执行用时：28 ms, 在所有 Go 提交中击败了95.07%的用户
 *     内存消耗：6.4 MB, 在所有 Go 提交中击败了28.70%的用户
 */

func findUnsortedSubarray(nums []int) int {

	errMinKey := 0
	errMaxKey := 0

	//找到错误开始最小的值
	for i := 0; i < len(nums); i++ {
		if len(nums)-1 != i && nums[i] > nums[i+1] {
			errMinKey = i
			break
		}
	}

	//找到错误结束最大值
	for i := len(nums) - 1; i >= 0; i-- {
		if 0 != i && nums[i] < nums[i-1] {
			errMaxKey = i
			break
		}
	}

	if errMinKey == 0 && errMaxKey == 0 {
		return 0
	}

	min := nums[errMinKey]
	max := nums[errMaxKey]

	//找到区间内的最大最小
	for i := errMinKey; i <= errMaxKey; i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	minKey := 0
	maxKye := 0

	//两次循环找到区间
	for i := 0; i < len(nums); i++ {
		if nums[i] > min {
			minKey = i
			break
		}
	}

	//找到错误结束最大值
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			maxKye = i + 1
			break
		}
	}

	return maxKye - minKey
}
