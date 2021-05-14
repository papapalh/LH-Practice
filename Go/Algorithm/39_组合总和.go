package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}

//思路
//	回溯算法
//	参考 46 全排列算法逻辑
//耗时
//	执行用时： 4 ms , 在所有 Go 提交中击败了 83.03% 的用户
//	内存消耗： 3.1 MB , 在所有 Go 提交中击败了 50.51% 的用户
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	sum := 0
	node := []int{}
	dict := map[string]struct{}{}

	dfs(&res, candidates, sum, node, target, dict)

	return res
}

func dfs(res *[][]int, candidates []int, sum int, node []int, target int, dict map[string]struct{}) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(node))
		copy(tmp, node)

		*res = append(*res, tmp)
		return
	}

	for _, v := range candidates {

		if len(node) > 0 && v > node[len(node) - 1] {
			continue
		}

		node = append(node, v)
		sum += v

		dfs(res, candidates, sum, node, target, dict)

		node = node[:len(node)-1]
		sum -= v
	}
}