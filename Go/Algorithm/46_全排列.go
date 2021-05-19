package main

import (
	"fmt"
)


func main() {
	fmt.Println(permute([]int{1,2,3,4}))
}

//思路
//	回溯算法
//	虽然写出来了，但是也是看的参考,后面几天继续练习回溯
//	https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
//耗时
//	执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//	内存消耗： 2.7 MB , 在所有 Go 提交中击败了 62.74% 的用户
func permute(nums []int) [][]int {
	res := [][]int{}

	//确定哪些头节点已经被选择
	used := map[int]bool{}

	//为空返回
	if len(nums) == 0 {
		return res
	}

	node := make([]int, 0)

	dfs(&res, 0, used, node, nums)

	return res
}

func dfs(res *[][]int, depth int, used map[int]bool, node []int, nums []int) {
	if depth == len(nums) {
		tmp := make([]int, len(node))
		copy(tmp, node)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++  {

		if r, ok := used[i]; r == true && ok {
			continue
		}

		node = append(node, nums[i])
		used[i] = true
		depth++

		dfs(res, depth, used, node, nums)

		used[i] = false
		node = node[:len(node)-1]
		depth--
	}
}