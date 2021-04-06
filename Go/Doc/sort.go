package main

import (
	"fmt"
	"sort"
)

////自定类型
//type Interface interface {
//	// Len方法返回集合中的元素个数
//	Len() int
//	// Less方法报告索引i的元素是否比索引j的元素小
//	Less(i, j int) bool
//	// Swap方法交换索引i和j的两个元素
//	Swap(i, j int)
//}

//类型自定义(可以是[]int/[]string/[]struct...)
type intSlice []int
func (s intSlice) Len() int { return len(s) }
func (s intSlice) Swap(i, j int) {s[i], s[j] = s[j], s[i] }
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {

	/*
	 *	[]int 类型
	 */

	//直接对 []int 类型做排序
	intSort := []int{4, 3, 2, 1, 9, 8, 7, 6}
	sort.Ints(intSort)
	fmt.Println(intSort)

	//检查切片是否已排序为递增顺序。
	fmt.Println(sort.IntsAreSorted(intSort))

	//在递增顺序的a中搜索x，返回x的索引。
	//如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	fmt.Println(sort.SearchInts(intSort, 5))


	/*
	 *	[]float64 类型
	 */
	//直接对 []float64 类型做排序
	float64Sort := []float64{1.1, 3.3, 2.2, 1.1, 9.9, 8.8, 7.7, 6,6}
	sort.Float64s(float64Sort)
	fmt.Println(float64Sort)

	//检查切片是否已排序为递增顺序。
	fmt.Println(sort.Float64sAreSorted(float64Sort))

	//在递增顺序的a中搜索x，返回x的索引。
	//如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	fmt.Println(sort.SearchFloat64s(float64Sort, 5.5))

	/*
	 *	[]strings 类型
	 */

	//直接对 []string 类型做排序
	stringSort := []string{"d", "a", "b", "c", "f"}
	sort.Strings(stringSort)
	fmt.Println(stringSort)

	//检查切片是否已排序为递增顺序。
	fmt.Println(sort.StringsAreSorted(stringSort))

	//在递增顺序的a中搜索x，返回x的索引。
	//如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	fmt.Println(sort.SearchStrings(stringSort, "e"))

	/*
	 *	类型自定义
	 */

	intSortMy := intSlice{4, 3, 2, 1, 9, 8, 7, 6}

	//直接对自定义类型做排序
	sort.Sort(intSortMy)
	fmt.Println(intSortMy)

	//检查切片是否已排序为递增顺序。
	fmt.Println(sort.IsSorted(intSortMy))
}
