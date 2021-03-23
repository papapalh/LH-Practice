package main

import (
	"fmt"
	"strconv"
)


// 给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold Medal", "Silver Medal", "Bronze Medal"）。
// (注：分数越高的选手，排名越靠前。)
// 示例 1:
// 		输入: [5, 4, 3, 2, 1]
// 		输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
// 解释: 
// 		前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal" and "Bronze Medal").
// 		余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。
// [10,3,8,9,4]
// 输出：
// ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
// 预期结果：
// ["Gold Medal","5","Bronze Medal","Silver Medal","4"]

func main() {
	fmt.Println(findRelativeRanks([]int{5,4,3,2,1}))
}

// 思路
// 		map假排序存储位置,最后进行归位
// 耗时
// 		执行用时8 ms, 在所有 Go 提交中击败了100.00%的用户
// 		内存消耗：6.5 MB, 在所有 Go 提交中击败了14.00%的用户
func findRelativeRanks(nums []int) []string {

	res := make([]string, len(nums))

	dict := map[int]int{}
	max := 0
	s := 1

	for i, v := range nums {
		dict[v] = i
		if v > max {
			max = v
		}
	}

	for i := max; i >= 0; i-- {

		if _, ok := dict[i]; !ok {
			continue
		}

		if s == 1 {
			res[dict[i]] = "Gold Medal"
		} else if s == 2 {
			res[dict[i]] = "Silver Medal"
		} else if s == 3 {
			res[dict[i]] = "Bronze Medal"
		} else {
			res[dict[i]] = strconv.Itoa(s)
		}
		s++
	}

	return res
}
