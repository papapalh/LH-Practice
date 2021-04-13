package main

import (
	"fmt"
)


func main() {
	fmt.Println(maxArea([]int{1,3,2,5,25,24,5}))
}

//思路
//	双指针-在区间的左右两边建立指针,逐渐缩小区间,找到最大的区间
//	https://leetcode-cn.com/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode-solution/
//耗时
//	执行用时： 104 ms , 在所有 Go 提交中击败了 30.89% 的用户
//	内存消耗： 8.3 MB , 在所有 Go 提交中击败了 38.34% 的用户
func maxArea(height []int) int {

	low, hei, res := 0, len(height) - 1, 0

	for hei > low {

		m := min(height[low], height[hei]) * (hei - low)
		if m > res {
			res = m
		}

		if height[low] >= height[hei] {
			hei--
		} else {
			low++
		}
	}

	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}