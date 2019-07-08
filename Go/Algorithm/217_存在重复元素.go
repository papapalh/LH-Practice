package main

import "fmt"
/**
 * 题目
 *     给定一个整数数组，判断是否存在重复元素。
 *     如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
 * 示例 1:
 *     输入: [1,2,3,1]
 *     输出: true
 * 示例 2:
 *     输入: [1,2,3,4]
 *     输出: false
 * 示例 3:
 *     输入: [1,1,1,3,3,4,3,2,4,2]
 *     输出: true
 */
func main () {
	a := containsDuplicate([]int{1,2,3,1})
	fmt.Println("-", a)
}
/**
 * 思路
 *     哈希表
 *     也是典型的空间换时间算法
 * 耗时
 *     执行用时 :28 ms, 在所有 Go 提交中击败了97.12%的用户
 *     内存消耗 :10.2 MB, 在所有 Go 提交中击败了21.95%的用户
 * 时间复杂度
 *     O（n）
 */
func containsDuplicate(nums []int) bool {
	
	dict := make(map[int]int)

	for _,v := range nums {

		if dict[v] > 0 {
			return true
		}
		
		dict[v]++
	}

	return false	
}