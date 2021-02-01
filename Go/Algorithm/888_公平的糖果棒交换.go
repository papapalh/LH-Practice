package main

import (
	"fmt"
)

//题目
//	爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 根糖果棒的大小，B[j] 是鲍勃拥有的第 j 根糖果棒的大小。
//  因为他们是朋友，所以他们想交换一根糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）
//	返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。
//如果有多个答案，你可以返回其中任何一个。保证答案存在。
//示例 1：
//	输入：A = [1,1], B = [2,2]
//	输出：[1,2]
//示例 2：
//	输入：A = [1,2], B = [2,3]
//	输出：[1,2]
//示例 3：
//	输入：A = [2], B = [1,3]
//	输出：[2,3]
//示例 4：
//	输入：A = [1,2,5], B = [2,4]
//	输出：[5,4]
//提示：
//	答案肯定存在。

func main() {

	a := []int{35,17,4,10,24}
	b := []int{63,21}

	fmt.Println(fairCandySwap(a, b))
}


//题解
//	其实这个更像是数学题
//  要让 A总量=B总量, 假设结果为{x,y}, 那么公式为 A-x+y = B+x-y
//  化简 x = y + (A-B)/2
//  那么我们循环A，让A的x满足这个公式就OK了
//	https://leetcode-cn.com/problems/fair-candy-swap/solution/gong-ping-de-tang-guo-jiao-huan-by-leetc-tlam/
//耗时
//	执行用时： 76 ms , 在所有 Go 提交中击败了 71.25% 的用户
//	内存消耗： 7.6 MB , 在所有 Go 提交中击败了 15.19% 的用户
func fairCandySwap(A []int, B []int) []int {
	aTotal := 0
	bTotal := 0
	bMap := map[int]int{}

	for _, v := range A {
		aTotal += v
	}

	for _, v := range B {
		bMap[v] = v
		bTotal += v
	}

	only := (aTotal-bTotal) / 2

	for _, v := range A {
		if bMap[v-only] > 0 {
			return []int{v, bMap[v-only]}
		}
	}

	return []int{}
}
