package main

import (
	"fmt"
	"strconv"
)

//字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
// 示例1:
// 输入："aabcccccaaa"
// 输出："a2b1c5a3"
// 示例2:
// 输入："abbccd"
// 输出："abbccd"
// 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
func main() {
	fmt.Println(compressString("abbccd"))
}

//双指针
//执行用时：80 ms, 在所有 Go 提交中击败了40.80%的用户
//内存消耗：7.9 MB, 在所有 Go 提交中击败了20.00%的用户
func compressString(S string) string {

	if len(S) == 0 {
		return S
	}

	outPut := ""
	low := 0
	hei := 0
	flag := ""

	for hei = 0; hei < len(S); hei++ {

		//处理首个字符
		if flag == "" {
			flag = string(S[hei])
			continue
		}

		if flag != string(S[hei]) {
			outPut = outPut + flag + strconv.Itoa(hei-low)
			low = hei
			flag = string(S[hei])
		}
	}

	//处理尾巴
	outPut = outPut + flag + strconv.Itoa(hei-low)

	//比较处理
	if len(outPut) >= len(S) {
		return S
	}

	return outPut
}
