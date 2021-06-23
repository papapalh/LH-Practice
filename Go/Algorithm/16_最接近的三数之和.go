package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1))
	fmt.Println(threeSumClosest2([]int{-1,2,1,-4}, 1))
}

//思路
//	1.三重循环暴力破解(居然过了也是不可思议)
//	2.先排序,取固定的中间值,双指针取合适的数据集合
//耗时
//	执行用时： 152 ms , 在所有 Go 提交中击败了 5.27% 的用户
//  内存消耗： 2.8 MB , 在所有 Go 提交中击败了 13.23% 的用户
func threeSumClosest(nums []int, target int) int {

	res := 0
	path := 999.00

	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			for k := j+1; k < len(nums); k++ {
				if done := nums[i] + nums[j] + nums[k]; math.Abs(float64(target - done)) < path  {
					res = done
					path = math.Abs(float64(target - done))
				}
			}
		}
	}

	return res
}

//思路
//	2.先排序,取固定的中间值,双指针取合适的数据集合
//
//耗时
//	执行用时： 4 ms , 在所有 Go 提交中击败了 94.63% 的用户
//  内存消耗： 2.7 MB , 在所有 Go 提交中击败了 100.00% 的用户
func threeSumClosest2(nums []int, target int) int {

	res := 0
	path := 999.99 //距离

	//排序之后通过双指正左右迁移
	sort.Ints(nums)

	//从1号执行(0号没有左指针,所以忽略)
	for i := 1; i < len(nums); i++ {

		left := i-1
		right := i+1

		for left >= 0 && right < len(nums){
			done := nums[left] + nums[right] + nums[i]

			if math.Abs(float64(target - done)) < path {
				res = done
				path = math.Abs(float64(target - done))
			}

			if done > target {
				//大于向左
				left -= 1
			} else if done < target {
				//小于向右
				right += 1
			} else {
				return done
			}
		}
	}

	return res
}

