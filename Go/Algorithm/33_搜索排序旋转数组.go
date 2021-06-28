package main

import "fmt"

func main() {

	fmt.Println(search([]int{5,1,3}, 5))
}

//思路
//	二分(需要特殊条件区分段)
//耗时
//
func search(nums []int, target int) int {

	res := -1

	if len(nums) == 0 {
		return res
	}

	left := 0
	right := len(nums)-1
	middle := 0

	//开始二分
	for right >= left {

		middle = (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > nums[left] {

			if nums[middle] > target && target >= nums[left] {
				right = middle - 1
			} else {
				left = middle + 1
			}

		} else {

			if nums[middle] < target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return res
}
