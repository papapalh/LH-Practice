package main

import "fmt"
/**
 * 题目
 *     给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
 *     不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
 *     元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
 *     注意这五个元素可为任意顺序。
 * 用例
 *     给定 nums = [3,2,2,3], val = 3,
 *     函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
 *
 *     给定 nums = [0,1,2,2,3,0,4,2], val = 2,
 *     函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
 */
func main () {

	arr := []int{0,1,2,2,3,0,4,2}

	c := removeElement(arr, 2)

	fmt.Printf("x", c)
}
/*
 * 思路
 *     类似双指针，原地算法
 * 耗时
 *     执行用时 : 0 ms, 在Remove Element的Go提交中击败了100.00% 的用户
 *     内存消耗 : 2.5 MB, 在Remove Element的Go提交中击败了30.24% 的用户
 * 时间复杂度
 *     O（n）
 */
func removeElement(nums []int, val int) int {

	point := 0

	for i := 0; i < len(nums);i++  {
		if nums[i] != val {
			nums[point] = nums[i]
			point++
		}
	}

	return point
}