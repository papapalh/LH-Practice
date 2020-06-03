package main

import "fmt"

func main() {
	fmt.Println(TwoFind())
}

/**
 * 思路
 *     二分查找
 *     适用于已经有序数据的查找。O(logn)
 * 步骤
 *     查找中位数，如果不满足，向前后推进指针。
 *	   直至满足或退出
 */
func TwoFind() bool {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tag := 2

	min := 0
	max := len(arr) - 1

	for i := 0; i <= 10; i++ {
		if min > max {
			break
		}

		middle := ((max - min) / 2) + min

		if tag > arr[middle] {
			min = middle + 1
		} else if tag < arr[middle] {
			max = middle - 1
		} else {
			return true
		}
	}

	return false
}
