package main

import "fmt"
/**
 * 题目
 *    给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *    你可以假设数组是非空的，并且给定的数组总是存在众数。
 * 示例 1:
 *    输入: [3,2,3]
 *    输出: 3
 * 示例 2:
 *    输入: [2,2,1,1,1,2,2]
 *    输出: 2
 */
func main () {
	x := majorityElement([]int{3,2,3})
	fmt.Println("x", x)
}
/*
 * 思路
 *    哈希表
 *    主要想看看哈希表是否可以像PHP一样往上加
 * 耗时
 *    执行用时 :32 ms, 在所有 Go 提交中击败了74.24%的用户
 *    内存消耗 :6 MB, 在所有 Go 提交中击败了20.78%的用户
 */
func majorityElement(nums []int) int {
	dict := make(map[int]int)

	macCount, max := 0
	
	for _,v := range nums {

		dict[v] += 1

		if dict[v] > macCount {
			macCount = dict[v]
			max = v
		}
	}

	return max
}