package main

import "fmt"

/**
 *   https://www.cnblogs.com/fivestudy/p/10212306.html
 *   排序算法   时间复杂度   最优          最差            稳定性
 *   冒泡排序   O(n^2)      O(n)         O(n^2)          稳定
 *   选择排序   O(n^2)      O(n^2)       O(n^2)          不稳定
 *   插入排序   O(n^2)      O(n)         O(n^2)          稳定
 *   归并排序   O(n log n)  O(n log n)   O(n log n)      稳定
 *   快速排序   O(n log n)  O(n log n)   O(n^2)          不稳定
 *   堆排序     O(n log n)  O(n log n)   O(n log n)      不稳定
 *   计数排序   O(n+k)      O(n+k)       O(n+k)          稳定
 *   桶排序     O(n^2)      O(n+k)       O(n^2)          稳定
 */

func main() {
	BubbleSort()
	SelectSort([]int{1,2,3,4})
	InsertSort([]int{1, 11, 12, 4, 2, 6, 9, 0, 3, 7, 8, 2})
	MergeSort([]int{1, 2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9})
	QuickSort([]int{1, 2, 3, 4})
	CountSort()
	BucketSort()
}

/**
 * 思路
 *     冒泡排序
 *     暴力破解的方法，在需要排序数很多的情况下，时间会很长。
 *     不适用实际的排序
 * 步骤
 *     比较相邻的元素。如果第一个比第二个大，就交换他们两个。
 *     对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
 *     针对所有的元素重复以上的步骤，除了最后一个。
 *     持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
 */
func BubbleSort() {
	// 常规的冒泡排序
	// 每个元素都与他的前一个元素作比较
	// 时间复杂度O(N2)/最好O(N2)
	arr := [...]int{4, 1, 6, 2, 9, 6, 7, 8}

	//外层定义循环次数
	for i := 0; i < len(arr)-1; i++ {

		//内存定义冒泡顺序
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}

	}
	fmt.Println(arr)

	// 改进
	// 增加flag参数，使在有序的情况下，获得更好的执行效率。
	// 在后面已经完全有序的情况下,减少计算
	// 时间复杂度为O(N2)/最好为O(N)
	arr2 := [...]int{4, 1, 6, 2, 9, 6, 7, 8}
	for i := 1; i < len(arr2); i++ {
		flag := 1
		for j := 0; j < len(arr2)-1; j++ {
			if arr2[j] > arr2[j+1] {
				arr2[j+1], arr2[j] = arr2[j], arr2[j+1]
				flag = 0
			}
		}
		if flag > 0 {
			break
		}
	}
	fmt.Println(arr2)
}

/**
 * 思路
 *     选择排序
 *     和冒泡排序类似，不过冒泡是在每个元素中交互，是稳定的。
 *     但是选择排序是根据遍历这个数组获取最大最小来重新操作这个数组。数组中本身的位置都已经改变了
 *     在自己做的时候用了系统的内置函数，其实和自己写差不多
 * 详解
 *     首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
 *     再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
 *     重复第二步，直到所有元素均排序完毕。
 */
func SelectSort(arr []int) []int {

	res := []int{}
	repeat := map[int]struct{}{} //避免重复处理

	for i := 0; i < len(arr)-1; i++ {

		minIndex := -999

		for j := 0; j < len(arr)-1; j++ {

			if _, ok := repeat[j]; ok {
				continue
			}

			if minIndex == -999 {
				minIndex = j
				continue
			}

			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		repeat[minIndex] = struct{}{}

		res = append(res, arr[minIndex])
	}

	return res
}

/**
 * 思路
 *     插入排序
 *     和冒泡排序类似，不过冒泡排序是不管之前是否已经排序完成的
 *     而插入排序，会认为前面的排序是已经排好序的序列
 *     是直接对排好序的序列进行操作
 * 详解
 *     将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
 *     从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
 *     如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。
 */
func InsertSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	//外层确定循环次数(从1开始,前面默认为排好的)
	for i := 1; i < len(arr); i++ {

		//排好序的数据移动会导致I的值发生变化,所以使用临时变量保存
		tmp := arr[i]

		//内层移动数据(j代表的是已经排好序的数据)
		for j := i - 1; j >= 0; j-- {

			if arr[j] > tmp {
				arr[j+1] = arr[j] //遍历前序数据，需要变更的数据向后推进一个
				arr[j] = tmp      //元素向前找地方插入
			} else {
				break //如果碰到不需要移动的元素，由于是已经排序好是数组，则前面的就不需要再次比较了。
			}
		}
	}

	return arr
}

/**
 * 思路
 *     归并排序
 *     对已经排序好的两个数组进行归并，效果最好
 * 详解
 *     比较两个数组的头元素，将比较小的数组头元素出栈，放入新数组
 *     一直循环这个数组，完成数组的排序
 */
func MergeSort(arr1, arr2 []int) []int {
	arr := []int{}

	leftPoint := 0
	rightPoint := 0

	for {
		if leftPoint == len(arr1) {
			arr = append(arr, arr2[rightPoint:]...)
			break
		}

		if rightPoint == len(arr2) {
			arr = append(arr, arr1[leftPoint:]...)
			break
		}

		if arr1[leftPoint] > arr2[rightPoint] {
			arr = append(arr, arr2[rightPoint])
			rightPoint++
		} else if arr1[leftPoint] < arr2[rightPoint] {
			arr = append(arr, arr1[leftPoint])
			leftPoint++
		} else {
			arr = append(arr, arr2[leftPoint], arr2[rightPoint])
			leftPoint++
			rightPoint++
		}
	}

	return arr
}

/**
 * 思路
 *     快速排序
 *     本质是通过 二分法不断的对数列进行排序
 *     快速排序基本上被认为在相同数量级中，平均性能最好的。
 * 详解
 *     在对数列进行排序时候，挑选一个基本数[flag][一般是第一个]把大于这个数的放右边，小于的放左边。
 *     之后再不停的递归对数列排序，直到最后一个。
 */
func QuickSort(arr []int) []int {

	//限界条件
	if len(arr) <= 1 {
		return arr
	}

	//随便选择一个，进行二分
	flag := arr[0]
	left := []int{}
	right := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] > flag {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	res := []int{}
	res = append(res, QuickSort(left)...)
	res = append(res, flag)
	res = append(res, QuickSort(right)...)

	return res
}

/**
 * 思路
 *     计数排序
 *     类似哈希表
 * 详解
 *     首先通过O（n）计算数组出先最大最小，以确定循环次数
 *     开辟额外空间存储每个元素出现位置和次数。
 *     根据最大最小区间循环输出，最后排序完成
 * 问题
 *     不适合离散度过大的数组
 */
func CountSort() {

	arr := []int{1, 2, 4, 6, 8, 10, 3, 1, 0, 4, 6}

	min := 0
	max := 0
	bucket := map[int][]int{}

	res := []int{}

	//O(N)遍历最大最小
	for _, v := range arr {
		if v > max {
			max = v
		}

		if min > v {
			min = v
		}
		//并放入对应的桶内
		bucket[v] = append(bucket[v], v)
	}

	//确定最大最小之后，遍历桶
	for i := min; i <= max; i++ {
		if len(bucket[i]) > 0 {
			res = append(res, bucket[i]...)
		}
	}

	fmt.Println(res)
}

/**
 * 思路
 *     桶排序
 *     使用于排序区间不大的排序
 * 步骤
 *     把每个元素放入对应的桶内(类型哈希表)
 *     最后按照顺序把每个桶的元素拿出来，完成排序
 */
func BucketSort() {
	arr := []int{1, 2, 4, 6, 8, 10, 3, 1, 0, 4, 6}
	bucket := map[int][]int{}

	max := 0
	res := []int{}

	//O(N)遍历最大最小
	for _, v := range arr {
		if v > max {
			max = v
		}
		//并放入对应的桶内
		bucket[v] = append(bucket[v], v)
	}

	//确定最大最小之后，遍历桶
	for i := 0; i <= max; i++ {
		if len(bucket[i]) > 0 {
			res = append(res, bucket[i]...)
		}
	}

	fmt.Println(res)
}
