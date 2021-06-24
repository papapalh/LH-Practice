package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-2147483648,-3))
}

//思路
//	这题真是费脑子，各种超限问题, int32-1 | -int32 各种处理
//	本质还是二分, 结果集每次 *2(翻倍) 看是否达达到了数据要求上线, 如果超过了说明在这个中间，在循环找到这个数
//	https://leetcode-cn.com/problems/divide-two-integers/solution/po-su-de-xiang-fa-mei-you-wei-yun-suan-mei-you-yi-/
//耗时
//	执行用时： 4 ms , 在所有 Go 提交中击败了 53.72% 的用户
//	内存消耗： 2.4 MB , 在所有 Go 提交中击败了 63.91% 的用户
func divide(dividend int, divisor int) int {

	if dividend == 0 {
		return 0
	}

	res := 1
	fu := 0

	if dividend < 0 {
		fu ^= 1
		dividend *= -1
	}

	if divisor < 0 {
		fu ^= 1
		divisor *= -1
	}

	if divisor > dividend {
		return 0
	}

	add := divisor

	for {

		left, right := divisor, divisor * 2

		if right >= dividend {
			for {
				if left <= dividend && (left+add) <= dividend  {
					res += 1
					left += add
				} else {
					break
				}
			}
			break
		}

		res *= 2
		divisor *= 2
	}

	if fu&1 == 1 {
		res *= -1
	}

	if res > math.MaxInt32 {
		res = math.MaxInt32
	}
	if res < -2147483648 {
		res = -2147483648
	}

	return res
}