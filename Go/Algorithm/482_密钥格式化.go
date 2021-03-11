package main

import (
	"fmt"
	"strings"
)

//有一个密钥字符串 S ，只包含字母，数字以及 '-'（破折号）。其中， N 个 '-' 将字符串分成了 N+1 组。
//给你一个数字 K，请你重新格式化字符串，使每个分组恰好包含 K 个字符。特别地，第一个分组包含的字符个数必须小于等于 K，但至少要包含 1 个字符。两个分组之间需要用 '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。
//给定非空字符串 S 和数字 K，按照上面描述的规则进行格式化。
//示例 1：
//
//输入：S = "5F3Z-2e-9-w", K = 4
//输出："5F3Z-2E9W"
//解释：字符串 S 被分成了两个部分，每部分 4 个字符；
//     注意，两个额外的破折号需要删掉。
//示例 2：
//
//输入：S = "2-5g-3-J", K = 2
//输出："2-5G-3J"
//解释：字符串 S 被分成了 3 个部分，按照前面的规则描述，第一部分的字符可以少于给定的数量，其余部分皆为 2 个字符。
func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}

//思路
//	倒序处理字符串,尾部后来满足K个一直向上
//耗时
//	执行用时： 4 ms , 在所有 Go 提交中击败了 89.61% 的用户
//	内存消耗： 4.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
func licenseKeyFormatting(S string, K int) string {

	res := ""

	cnt := len(S)
	flag := 0

	S = strings.ToUpper(string(S))

	for i := cnt - 1; i >= 0; i-- {

		if string(S[i]) == "-" {
			continue
		}

		if flag == K {
			res = "-" + res
			flag = 0
		}

		res = string(S[i]) + res
		flag++

	}

	return res
}