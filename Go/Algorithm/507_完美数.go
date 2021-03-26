package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(1))
}

//思路
//	只需要算出一半的正因数就行。
//  比如 28 = 1 + 2 + 4 + 7 + 14 = 28
//	那么 2*14、1*28、4*6 都是一组
//	那么我们只需要计算一半的数据就行
//耗时
//	执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//	内存消耗： 1.9 MB , 在所有 Go 提交中击败了 33.33% 的用户
func checkPerfectNumber(num int) bool {

	sum := 0

	//1的话单独处理,直接返回false
	if num == 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {

		if num % i == 0 {
			sum += i
			sum += num / i
		}
	}

	//1的话单独处理,如果在循环中则会处理两次
	sum += 1

	return sum == num
}