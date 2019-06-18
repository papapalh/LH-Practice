package main

import "fmt"
/**
 * 题目
 *    给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 * 说明：
 *    你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 * 示例 1:
 *    输入: [2,2,1]
 *    输出: 1
 * 示例 2:
 *    输入: [4,1,2,1,2]
 *    输出: 4
 */
func main () {
	x := singleNumber([]int{2,2,1})
	fmt.Println("x", x)
}
/**
 * 思路
 *    位操作
 * 耗时
 *    执行用时 :24 ms, 在所有Go提交中击败了43.36%的用户
 *    内存消耗 :4.8 MB, 在所有Go提交中击败了82.17%的用户
 */
func singleNumber(nums []int) int {

	sum := 0

	for _,v := range nums {
		sum = sum ^ v
	}

	return sum
}