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
	SelectSort()
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
	// 常规的冒泡排序，每个元素都与他的前一个元素作比较，时间复杂度O(N2)/最好O(N2)
	arr := [...]int{4, 1, 6, 2, 9, 6, 7, 8}
	for i := 1; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	fmt.Println(arr)

	// 改进，增加flag参数，使在有序的情况下，获得更好的执行效率，时间复杂度为O(N2)/最好为O(N)
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
func SelectSort() {
	arr := [...]int{1, 11, 12, 4, 2, 6, 9, 0, 3, 7, 8, 2}
	newArr := []int{}            //新处理切片
	repeat := map[int]struct{}{} //避免重复处理

	for i := 0; i < len(arr)-1; i++ {

		minkey := -999

		for j := 0; j < len(arr)-1; j++ {

			if _, ok := repeat[j]; ok {
				continue
			}

			if minkey == -999 {
				minkey = j
				continue
			}

			if arr[minkey] > arr[j] {
				minkey = j
			}
		}

		repeat[minkey] = struct{}{}

		newArr = append(newArr, arr[minkey])
	}

	fmt.Println(newArr)
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
func InsertSort() {
	arr := [...]int{1, 11, 12, 4, 2, 6, 9, 0, 3, 7, 8, 2}

	for i := 1; i < len(arr); i++ {

		tmp := arr[i]

		for j := i - 1; j >= 0; j-- {

			if arr[j] > tmp {
				arr[j+1] = arr[j] //遍历前序数据，需要变更的数据向后推进一个
				arr[j] = tmp      //元素向前找地方插入
			} else {
				break //如果碰到不需要移动的元素，由于是已经排序好是数组，则前面的就不需要再次比较了。
			}
		}
	}

	fmt.Println(arr)
}

/**
 * 思路
 *     归并排序
 *     对已经排序好的两个数组进行归并，效果最好
 * 详解
 *     比较两个数组的头元素，将比较小的数组头元素出栈，放入新数组
 *     一直循环这个数组，完成数组的排序
 */
func MergeSort() {
	arr1 := [...]int{1, 2, 4, 6, 8, 10}
	arr2 := [...]int{1, 3, 5, 7, 9}
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

	fmt.Println(arr)
}
