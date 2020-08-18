package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 9, 8}))
}

/**
 * 题目
 *    给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *    最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *    你可以假设除了整数 0 之外，这个整数不会以零开头。
 * 示例 1:
 *    输入: [1,2,3]
 *    输出: [1,2,4]
 *    解释: 输入数组表示数字 123。
 * 示例 2:
 *    输入: [4,3,2,1]
 *    输出: [4,3,2,2]
 *    解释: 输入数组表示数字 4321。
 * 思路
 *     1.变成数字+1 => 变成数组
 *        经过试验，不行(会超过数字类型的最大限制)
 *     2.进位处理
 *        循环处理，需要特殊处理高位
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：2 MB, 在所有 Go 提交中击败了64.44%的用户
 */
func plusOne(digits []int) []int {

	//检查指标
	checkPoint := len(digits) - 1

	//递归进位
	for {
		//特殊情况-已经到达了高位
		if checkPoint == -1 {
			digits = append([]int{1}, digits...)
			break
		}

		//9则进位
		if digits[checkPoint] == 9 {
			digits[checkPoint] = 0
			checkPoint--
		} else {
			digits[checkPoint] += 1
			break
		}
	}

	return digits
}
