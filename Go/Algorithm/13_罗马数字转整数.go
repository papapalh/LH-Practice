package main

import (
	"fmt"
)
// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//     字符          数值
//     I             1
//     V             5
//     X             10
//     L             50
//     C             100
//     D             500
//     M             1000
// 示例
//     输入: "III"     输出: 3
//     输入: "IV"      输出: 4
//     输入: "IX"      输出: 9
//     输入: "LVIII"   输出: 58    解释: L = 50, V= 5, III = 3.
//     输入: "MCMXCIV" 输出: 1994  解释: M = 1000, CM = 900, XC = 90, IV = 4.
func main() {
	x := romanToInt("MCMXCIV")
	fmt.Printf("echo", x)
}
// 思路
//     由于GO的强类型限制，不能像PHP一样直接去判断数组的不存的位
//     这里使用了 flag 标志位标记前一位，判断后一位是否大于前一位，则减，否则加
//耗时
//     执行用时 : 8 ms, 在Roman to Integer的Go提交中击败了98.41% 的用户
//     内存消耗 : 3 MB, 在Roman to Integer的Go提交中击败了57.68% 的用户
//时间复杂度
//     O(n)
func romanToInt(s string) int {

	dict := map[string]int{
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}

	flag := 0
	sum := 0

	for _,v := range s {

		if flag == 0 {
			flag = int(v)
			sum += dict[string(v)]
			continue
		}

		if dict[string(v)] > dict[string(flag)] {
			sum += dict[string(v)] - (dict[string(flag)] * 2)
			flag = int(v)
			continue
		}

		flag = int(v)
		sum += dict[string(v)]
	}

	return sum
}
