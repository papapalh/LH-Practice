package main

import "fmt"

/**
 * 题目
 *    两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
 *    给出两个整数 x 和 y，计算它们之间的汉明距离。
 * 注意：
 *    0 ≤ x, y < 231.
 * 示例:
 *    输入: x = 1, y = 4
 *    输出: 2
 * 解释:
 *    1   (0 0 0 1)
 *    4   (0 1 0 0)
 *           ↑   ↑
 *    上面的箭头指出了对应二进制位不同的位置。
 */

func main() {
	fmt.Println(hammingDistance(1, 4))
}

/**
 * 思路
 *    逐位比较。
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func hammingDistance(x int, y int) int {

	res := 0

	for {

		//跳出条件
		if x == 0 && y == 0 {
			break
		}

		//位相等判断
		if x&1 != y&1 {
			res++
		}

		//整体向后推一位
		x >>= 1
		y >>= 1
	}

	return res
}
