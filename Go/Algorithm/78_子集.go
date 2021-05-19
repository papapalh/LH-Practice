package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1,2,3}))
}

//思路
//	还时回溯算法,把所有的路径结果全部拉出来
//耗时
//	执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//	内存消耗： 2.3 MB , 在所有 Go 提交中击败了 52.46% 的用户
func subsets(nums []int) [][]int {

	res := make([][]int, 0)

	if len(nums) == 0 {
		return res
	}

	used := map[int]bool{}
	xxx(&res, []int{}, used, nums)

	return res
}

func xxx(res *[][]int, node []int, used map[int]bool, nums []int) {

	tmp := make([]int, len(node))
	copy(tmp, node)
	*res = append(*res, tmp)
	if len(node) == len(nums) {
		return
	}

	for i, v := range nums {

		if len(node) > 0 && v < node[len(node) - 1]  {
			continue
		}

		if r, ok := used[i]; r == true && ok {
			continue
		}

		node = append(node, v)
		used[i] = true

		xxx(res, node, used, nums)

		used[i] = false
		node = node[:len(node)-1]
	}
}