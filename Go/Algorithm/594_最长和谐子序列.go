package main

import "fmt"

func main(){

	a := []int{1,3,2,2,5,2,3,7}

	fmt.Println(findLHS(a))
	fmt.Println(findLHS2(a))
}

//题目: https://leetcode-cn.com/problems/longest-harmonious-subsequence/

/*
 *  方法
 *      哈希映射
 *      先哈希映射一遍所有数字的次数，第二次在进行 +-1 的潘墩
 *  复杂度
 *      O(N2)
 *  耗时
 *      执行用时 : 64 ms , 在所有 Go 提交中击败了 92.00% 的用户
 *      内存消耗 : 6.7 MB , 在所有 Go 提交中击败了 100.00% 的用户
 */
func findLHS(nums []int) int {

	recode := map[int]int{}
	max := 0

	for _, v := range nums {
		recode[v]++
	}

	for value, count := range recode {

		next := recode[value + 1]
		prev := recode[value - 1]

		if next > 0 && (next + recode[value]) > max {
			max = next + count
		}

		if prev > 0 && (prev + recode[value]) > max {
			max = prev + count
		}
	}

	return max
}

/*
 *  方法
 *      哈希映射 + 一次遍历
 *      这个和上面的方法其实差不多，就是这个题可以这个弄，有点取巧的意思
 *      哈希映射次数的同时进行比对，肯定在遍历到最后时候会出现最大值
 *  复杂度
 *      O(N2)
 *  耗时
 *      执行用时 : 72 ms , 在所有 Go 提交中击败了 69.33% 的用户
 *      内存消耗 : 6.7 MB , 在所有 Go 提交中击败了 100.00% 的用户
 */
func findLHS2(nums []int) int {

	recode := map[int]int{}
	max := 0

	for _, v := range nums {

		recode[v]++

		next := recode[v + 1]
		prev := recode[v - 1]

		if next > 0 && (next + recode[v]) > max {
			max = next + recode[v]
		}

		if prev > 0 && (prev + recode[v]) > max {
			max = prev + recode[v]
		}
	}

	return max
}
