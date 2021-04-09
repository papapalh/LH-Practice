package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("aba"))
}

//思路
//	中心扩散法(牛B)
//	https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zhong-xin-kuo-san-fa-he-dong-tai-gui-hua-by-reedfa/
//耗时
//	执行用时： 8 ms , 在所有 Go 提交中击败了 85.48% 的用户
//	内存消耗： 2.6 MB , 在所有 Go 提交中击败了 100.00% 的用户
func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	left, right, length, maxLen, maxStart := 0, 0, 1, 0, 0

	for i := 0; i < len(s); i++ {
		left = i - 1  //左
		right = i + 1 //右

		for left >= 0 && s[i] == s[left] {
			left--
			length++
		}

		for right < len(s) && s[i] == s[right]{
			right++
			length++
		}

		for left >= 0 && right < len(s) && s[left] == s[right] {
			length += 2
			left--
			right++
		}

		if length > maxLen {
			maxLen = length
			maxStart = left
		}

		length = 1
	}

	return s[maxStart+1:maxStart+maxLen+1]
}

func abc(s string) bool {

	l := len(s)
	middle := l / 2

	for i := 0; i < middle; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}

func e(xx int, l *sync.Mutex){
	l.Lock()
	fmt.Println(xx)
	time.Sleep(time.Second * 1)
	defer func() {
		l.Unlock()
	}()
}