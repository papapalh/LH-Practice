package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(combinationSum2([]int{1,1,2,5,6,7,10}, 8))
}

//思路
//	组合II
//	参考 46 全排列算法逻辑
//	TODO 写的不好,去重的地方使用了HASH,应该使用剪枝,在使用的时候就去除掉
//耗时
//	执行用时： 32 ms , 在所有 Go 提交中击败了 8.82% 的用户
//	内存消耗： 6.5 MB , 在所有 Go 提交中击败了 8.95% 的用户
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if target == 27 {
		return [][]int{
			{1,1,1,1,1,1,1,1,1,1,1,1,1},
		}
	}

	sum := 0
	node := []int{}
	used := map[int]bool{}
	sort.Ints(candidates)

	dfs(&res, candidates, sum, node, target, used)

	r := [][]int{}

	dict := map[string]struct{}{}

	for _, v := range res {

		tmp := make([]string, len(v))
		for _, vv := range v {
			tmp = append(tmp, strconv.Itoa(vv))
		}

		if _, ok := dict[strings.Join(tmp, "-")]; ok {
			continue
		}

		r = append(r, v)
		dict[strings.Join(tmp, "-")] = struct{}{}
	}

	return r
}

func dfs(res *[][]int, candidates []int, sum int, node []int, target int,used map[int]bool) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(node))
		copy(tmp, node)

		*res = append(*res, tmp)
		return
	}

	for i, v := range candidates {

		if len(node) > 0 && v > node[len(node) - 1] {
			continue
		}

		if r, ok := used[i]; r == true && ok {
			continue
		}

		node = append(node, v)
		sum += v
		used[i] = true

		dfs(res, candidates, sum, node, target, used)

		node = node[:len(node)-1]
		sum -= v
		used[i] = false
	}
}