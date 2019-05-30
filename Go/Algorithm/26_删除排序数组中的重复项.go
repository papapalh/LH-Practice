package main

import "fmt"
/**
 * 题目
 *     给定数组 nums = [1,1,2]
 *     函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
 *     你不需要考虑数组中超出新长度后面的元素。
 * 用例
 *     给定 nums = [0,0,1,1,1,2,2,3,3,4],
 *     函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
 *
 *     给定数组 nums = [1,1,2],
 *     函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
 */
func main () {

	arr := []int{0,0,1,1,1,2,2,3,3,4}

	c := removeDuplicates(arr)

	fmt.Printf("x", c)
}
/**
 * 思路2
 *    双指针(官方解法)
 *    对于不同的数进行位置交换，同时记录指针移动
 * 耗时
 *    执行用时 : 92 ms, 在Remove Duplicates from Sorted Array的Go提交中击败了95.90% 的用户
 *    内存消耗 : 8.2 MB, 在Remove Duplicates from Sorted Array的Go提交中击败了36.75% 的用户
 * 时间复杂度
 *    O(n)
 */
func removeDuplicates(nums []int) int {

	len := len(nums)

	if len <= 1 {
		return len
	}

	flag  := nums[0]
	point := 1

	for i := 1; i < len; i++ {
		if flag != nums[i] {

			nums[point] = nums[i]

			point++ // 指针移动
			flag = nums[i] // flag 变化
		}
	}

	return point
}