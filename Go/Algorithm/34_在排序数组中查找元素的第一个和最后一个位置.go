package main

import (
	"fmt"
)

func main() {

	fmt.Println(searchRange([]int{0,0,2,3,4,4,4,5}, 5))

}

//思路
//	先二分法找到节点位置(在哪个位置找到都行),在中心扩散找到边界
//	最好log2n最差n
//耗时
//	执行用时： 8 ms , 在所有 Go 提交中击败了 88.67% 的用户
//  内存消耗： 3.9 MB , 在所有 Go 提交中击败了 60.62% 的用户
func searchRange(nums []int, target int) []int {

	left, right := -1, -1

	low, hei := 0, len(nums) - 1

	for low <= hei {
		middle := (hei - low) / 2 + low

		if nums[middle] > target {
			hei--
		} else if nums[middle] < target {
			low = low + 1
		} else {

			left, right = middle, middle

			//确定左边界
			left_tag := 1
			for {
				if middle == 0 {
					left = middle
					break
				}

				if middle - left_tag < 0 {
					break
				}

				if nums[middle-left_tag] != nums[middle] {
					break
				}

				left = middle - left_tag
				left_tag++
			}

			//确定右边界
			right_tag := 1
			for {
				if middle == len(nums) - 1 {
					right = middle
					break
				}

				if middle + right_tag > len(nums) - 1 {
					break
				}

				if nums[middle+right_tag] != nums[middle] {
					break
				}

				right = middle + right_tag
				right_tag++
			}

			break
		}
	}


	return []int{left, right}
}