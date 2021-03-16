package main

import (
	"fmt"
)

//题目: https://leetcode-cn.com/problems/island-perimeter/
func main() {
	abc := [][]int{
		{0,1,0,0},
		{1,1,1,0},
		{0,1,0,0},
		{1,1,0,0},
	}

	fmt.Println(islandPerimeter(abc))
}

//思路
//	迭代
//  看每个land所占用的周长，上下进行对比
//耗时
//	执行用时： 76 ms , 在所有 Go 提交中击败了 52.07% 的用户
//	内存消耗： 6.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
func islandPerimeter(grid [][]int) int {

	all := 0

	for n, g := range grid {

		for k, v := range g {
			if v == 0 {
				continue
			}

			//上层
			if n == 0 {
				//没有上层则算一条边
				all++
			} else if grid[n-1][k] == 0 {
				//如果上册是海则算一条边
				all++
			}

			//下层
			if n == len(grid) - 1 {
				//如果是最下层则算一条边
				all++
			} else if grid[n+1][k] == 0 {
				//如果下层是海则算一条边
				all++
			}

			//如果左面是海(0)+1
			if k == 0 || g[k-1] == 0 {
				all++
			}
			//如果又面是海(0)+1
			if k == len(g)-1 || g[k+1] == 0 {
				all++
			}
		}

	}

	return all
}