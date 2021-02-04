package main

import (
	"fmt"
	"sort"
)

//题目
//给定一个长度为 n 的 非空 整数数组，每次操作将会使 n - 1 个元素增加 1。找出让数组所有元素相等的最小操作次数。
//输入：
//	[1,2,3]
//输出：
//	3
//解释：
//	只需要3次操作（注意每次操作会增加两个元素的值）：
//	[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
func main() {
	a := []int{1,2147483647}
	fmt.Println(minMoves(a))
	fmt.Println(a)
}


//思路
//	排序切片数据，直接可以对比每个值和初始值的差距
//耗时
//	执行用时： 72 ms , 在所有 Go 提交中击败了 20.43% 的用户
//	内存消耗： 6.6 MB , 在所有 Go 提交中击败了 45.16% 的用户
func minMoves(nums []int) int {

	sort.Ints(nums)

	count := 0

	for i := len(nums) - 1; i > 0; i-- {
		count += nums[i] - nums[0]
	}

	return count
}