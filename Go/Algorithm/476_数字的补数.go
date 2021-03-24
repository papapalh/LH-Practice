package main

import (
	"fmt"
)

// 给你一个 正 整数 num ，输出它的补数。补数是对该数的二进制表示取反。
// 示例 1：
// 		输入：num = 5
// 		输出：2
// 		解释：5 的二进制表示为 101（没有前导零位），其补数为 010。所以你需要输出 2 。
// 示例 2：
// 		输入：num = 1
// 		输出：0
// 		解释：1 的二进制表示为 1（没有前导零位），其补数为 0。所以你需要输出 0 

func main() {
	fmt.Println(findComplement(7))
}

// 思路
// 		本质是通过抑或处理(相同为0相异为1)
// 		构造一个与num相同位的全是1的数字，和num抑或，产生补数
// 		https://leetcode-cn.com/problems/number-complement/solution/yi-huo-by-im-me/
// 耗时
// 		执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 		内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func findComplement(num int) int {
	tmp := 1;
	for tmp < num {
		tmp <<= 1;
		tmp += 1;
	}

	return tmp^num;
}
