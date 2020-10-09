package main

func main() {
	countPrimes(499979)
}

var primesMap = map[int]bool{}

/**
 * 思路1
 *    暴力解决，每次-1判断数字是否是素数
 * 耗时
 *    超时
 */
func countPrimes1(n int) int {

	num := 0

	if n <= 2 {
		return num
	}

	for i := n - 1; i > 0; i-- {

		r, ok := primesMap[i]

		//不存在则开始计算
		if !ok {

			if isPrimes(i) {
				num += 1
				primesMap[i] = true
			} else {
				primesMap[i] = false
			}

			continue
		}

		//是质数则+1
		if r {
			num += 1
			continue
		}

		//不是是质数则不变
		// if !r {
		// 	continue
		// }
	}

	return num
}
func isPrimes(n int) bool {

	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n&1 == 0 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

/**
 * 思路2
 *    厄拉多塞筛法
 *        当发现一个素数时，排除这个质数的所有可能
 *        比如输入 64 那么发现2是质数，则排除 2.4.6.8.10.12.14....... < 64 的可能
 *        最后一遍统计
 *    总结一句话
 *        是一个质数倍数的数字，不可能是一个质数
 * 耗时
 *    执行用时：8 ms, 在所有 Go 提交中击败了95.69%的用户
 *    内存消耗：4.9 MB, 在所有 Go 提交中击败了46.07%的用户
 */
func countPrimes(n int) int {
	count := 0
	signs := make([]bool, n)
	for i := 2; i < n; i++ {
		if signs[i] {
			continue
		}
		count++
		for j := 2 * i; j < n; j += i {
			signs[j] = true
		}
	}
	return count
}
