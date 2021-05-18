package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4,2))
}

//思路
//	依旧是回溯算法
//	参考全排列逻辑
//耗时
//	执行用时： 100 ms , 在所有 Go 提交中击败了 5.13% 的用户
//	内存消耗： 9.6 MB , 在所有 Go 提交中击败了 16.29% 的用户
func combine(n int, k int) [][]int {

	res := [][]int{}

	used := map[int]bool{}
	node := []int{}

	dfs(&res, n, k, node, used)

	return res
}

func dfs(res *[][]int, n int, k int, node []int, used map[int]bool) {
	if len(node) == k {
		tmp := make([]int, len(node))
		copy(tmp, node)
		*res = append(*res, tmp)
		return
	}

	for i := 1; i <= n; i++ {

		if len(node) > 0 && i <= node[len(node) - 1] {
			continue
		}

		if r, ok := used[i]; r == true && ok {
			continue
		}

		node = append(node, i)
		used[i] = true

		xxx(res, n, k, node, used)

		node = node[:len(node)-1]
		used[i] = false
	}
}
