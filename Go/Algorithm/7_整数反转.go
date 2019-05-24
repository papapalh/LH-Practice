 /**
 * 题目
 *    给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *    假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31, 2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回0
 * 用例
 *    输入: 123
 *    输出: 321
 *    输入: -123
 *    输出: -321
 */
package main

import (
	"fmt";
)
/**
 * 思路
 *    如果数字为负，则取绝对值处理，最后加符号
 *    超出范围，则爆 0
 * 耗时
 *     执行用时 : 4 ms, 在Reverse Integer的Go提交中击败了96.15% 的用户
 *     内存消耗 : 2.2 MB, 在Reverse Integer的Go提交中击败了55.10% 的用户
 * 时间复杂度 
 *     O(n)
 */
func main() {
	i := reverse(-123)

	fmt.Printf("xxx", i)
}

func reverse(x int) int {

	result := 0
	flag := false

	if x < 0 {
		x *= -1
		flag = true
	}

    for ;x >= 1; {

		result = (result * 10) + (x % 10)

		x = x / 10
	}

	if result >= 2147483647 || result <= -2147483648  {
		return 0
	}

	if flag {
		result *= -1
	}

	return result
}