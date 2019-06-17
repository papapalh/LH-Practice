package main

import "fmt"
/**
 * 题目
 *    给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *    如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
 *    注意你不能在买入股票前卖出股票。
 * 示例 1:
 *    输入: [7,1,5,3,6,4]
 *    输出: 5
 *    解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 *    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 * 示例 2:
 *    输入: [7,6,4,3,1]
 *    输出: 0
 *    解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 */
func main () {

	a := []int{2,1,2,1,0,1,2}

	x := maxProfit(a)

	fmt.Println("x", x)

}
/**
 * 思路
 *    感觉这种题，真的是找规律
 * 耗时
 *    执行用时 :4 ms, 在所有Go提交中击败了98.94%的用户
 *    内存消耗 :3.2 MB, 在所有Go提交中击败了47.01%的用户
 */
func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	min, max := prices[0], 0

	for _, money := range prices {

		if money < min {
			min = money
		}

		if max < money - min {
			max = money - min
		}
	}

	return max
}