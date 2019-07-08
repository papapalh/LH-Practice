package main

import (
    "fmt"
)
/**
 * 题目
 *     给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
 * 示例 1:
 *    输入: nums = [1,2,3,1], k = 3
 *    输出: true
 * 示例 2:
 *    输入: nums = [1,0,1,1], k = 1
 *    输出: true
 * 示例 3:
 *    输入: nums = [1,2,3,1,2,3], k = 2
 *    输出: false
 */
func main () {
	a := containsNearbyDuplicate([]int{1,0,1,1}, 1)
	fmt.Println("-", a)
}
/*
 * 思路
 *    第一个方法由于对GO的数据结构还是太薄弱，确实是想的太多了。
 *    整体写的太复杂了，总体来说就是建立 哈希表 存储上一个元素的值，进行比对
 * 耗时
 *    执行用时 :64 ms, 在所有 Go 提交中击败了23.53%的用户
 *    内存消耗 :12.6 MB, 在所有 Go 提交中击败了7.32%的用户
*/
func containsNearbyDuplicate1(nums []int, k int) bool {
	
	if len(nums) <= 1 {
		return false
	}

	relation := make(map[int]int) // 关联
	dict := [][]int{}

	point := 0
	for i,v := range nums {

		if relation[v] == 0 {
			dict = append(dict, []int{})
			dict[point] = append(dict[point], i)
			relation[v] = point+1
			point++
			continue
		}

		dict[relation[v]-1] = append(dict[relation[v]-1], i)

		if dict[relation[v]-1][len(dict[relation[v]-1]) - 1] - dict[relation[v]-1][len(dict[relation[v]-1]) - 2] <= k {
			return true
		}
	}

	return false
}

/*
 * 思路
 *    修正上一个思路，果然快了很多
 * 耗时
 *    执行用时 :24 ms, 在所有 Go 提交中击败了93.53%的用户
 *    内存消耗 :8.8 MB, 在所有 Go 提交中击败了48.78%的用户
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	
	if len(nums) <= 1 {
		return false
	}

	dict := make(map[int]int) // 关联

	for i,v := range nums {

		if dict[v] == 0 {
			dict[v] = i + 1
			continue
		}

		if i + 1 - dict[v] <= k {
			return true
		}

		dict[v] = i + 1
	}

	return false
}