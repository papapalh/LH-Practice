package main

import "fmt"
/**
 * 题目
 *    给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 * 说明:
 *    初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 *    你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 * 示例:
 *    输入:
 *        nums1 = [1,2,3,0,0,0], m = 3
 *        nums2 = [2,5,6],       n = 3
 *    输出: [1,2,2,3,5,6]
 */
func main() {

	a := []int{1,2,8,0,0,0}
	b := []int{2,5,6}

	merge(a, 3, b, 3)

	fmt.Printf("-", a)
}
/*
 * 思路
 *    归并排序
 *    归并两数组结果，赋予原数组，关于GO的底层知识运用还应该更进一步，现在还是太肤浅
 * 耗时
 *    执行用时 : 0 ms, 在Merge Sorted Array的Go提交中击败了100.00% 的用户
 *    内存消耗 : 3.9 MB, 在Merge Sorted Array的Go提交中击败了5.19% 的用户
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {

	if nums2 == nil {
		return
	}

	arr := make([]int, 0)

	nPoint := 0 // nums2 指针
	mPoint := 0 // nums1 指针

	for i := 0; i <= m+n; i++ {

		if nPoint >= n && mPoint >= m  {
			continue
		}

		if nPoint >= n {
			arr = append(arr, nums1[mPoint])
			mPoint++
			continue
		}

		if mPoint >= m {
			arr = append(arr, nums2[nPoint])
			nPoint++
			continue
		}

		if nums1[mPoint] >= nums2[nPoint] {
			arr = append(arr, nums2[nPoint])
			nPoint++
		} else {
			fmt.Printf("-", mPoint)
			arr = append(arr, nums1[mPoint])
			mPoint++
		}
	}

	for i, v := range arr {
		nums1[i] = v
	}
}