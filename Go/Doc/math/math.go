package main

import (
	"fmt"
	"math"
)

func main() {

	/*
	 * 数学操作
	 */
	//向上取整
	fmt.Println(math.Ceil(1.1))

	//向下取整(除法默认也是向下取整)
	fmt.Println(math.Floor(1.1))

	//绝对值
	fmt.Println(math.Abs(1.1))

	/*
	 * 数字判断
	 */
	//判断一个数字是否是NaN(取值极限)
	fmt.Println(math.IsNaN(1.1))

	//判断数字是否是负数/零
	fmt.Println(math.Signbit(-1.1))

	/*
	 * 浮点数
	 */
	//获取整数部分
	fmt.Println(math.Trunc(1.1))

	//获取整数+小数部分
	fmt.Println(math.Modf(1.1))

	/*
	 * 其他
	 */
	//最大
	fmt.Println(math.Max(1.1, 2.2))

	//最小
	fmt.Println(math.Min(1.1, 2.2))

	//返回x-y和0中的最大值
	fmt.Println(math.Dim(1.1, 2.2))

	//取余
	fmt.Println(math.Mod(10, 3))

	//x的二次方根
	fmt.Println(math.Sqrt(9))

	//x的三次方根
	fmt.Println(math.Cbrt(27))

	//x的y次方
	fmt.Println(math.Pow(2, 3))

	//x的10次方
	fmt.Println(math.Pow10(2))
}
