package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

/**
 * 题目
 *    斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
 *    F(0) = 0,   F(1) = 1
 *    F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
 *    给定 N，计算 F(N)。
 * 思路
 *    不动脑子的循环处理
 * 耗时(这就100%了？)
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：1.9 MB, 在所有 Go 提交中击败了55.74%的用户
 */
func fib(N int) int {

	//记录每个波菲那契数的值
	d := map[int]int{
		0: 0,
		1: 1,
	}

	if v, ok := d[N]; ok {
		return v
	}

	i := 2
	for {

		d[i] = d[i-1] + d[i-2]

		if i == N {
			return d[i]
		}

		i++
	}
}
