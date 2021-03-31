package main

import (
	"fmt"
	"math"
)

//题目
//	给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//	找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//	您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//输入:
//	[4,3,2,7,8,2,3,1]
//输出:
//	[5,6]
func main() {
	fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
}

//思路
//	标记: 对数组每个出现的位置打-1标记，再次遍历，没打标记的就是没出现的数字
//	桶排序: 把数组每个出现的数字放入桶内，没出现的桶就是没出现的数据(需要额外空间)
//耗时
//	执行用时： 60 ms , 在所有 Go 提交中击败了 69.30% 的用户
//	内存消耗： 7.7 MB , 在所有 Go 提交中击败了 22.72% 的用户
func findDisappearedNumbers(nums []int) []int {

	res := []int{}

	for _, v := range nums {
		if nums[int(math.Abs(float64(v)))-1] > 0 {
			nums[int(math.Abs(float64(v)))-1] *= -1
		}
	}

	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}

	return res
}