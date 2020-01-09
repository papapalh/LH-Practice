package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("    NaN()       取值极限 ", math.NaN())
	fmt.Println("    IsNaN(f)    f是否是一个NaN ", math.IsNaN(1.1))
	fmt.Println("    Signbit(x)  x是一个负数或者负零,返回真。", math.Signbit(-1.1))
	fmt.Println("    Ceil(x)     x向上取整", math.Ceil(1.1))
	fmt.Println("    Floor(x)    x向下取整 ", math.Floor(1.1))
	fmt.Println("    除法        默认向下取整 ", 10/3)
	fmt.Println("    Trunc(x)    x的整数部分 ", math.Trunc(1.1))


	fmt.Print("    Modf(f)     f的整数部分和小数部分，结果的正负号和都f相同 ")
		fmt.Println(math.Modf(1.1))
	fmt.Println("    Abs(x)      x的绝对值 ", math.Abs(1.1))
	fmt.Println("    Max(x,y)    返回x和y中最大值 ", math.Max(1.1, 2.2))
	fmt.Println("    Min(x,y)    返回x和y中最小值 ", math.Min(1.1, 2.2))
	fmt.Println("    Dim(x,y)    函数返回x-y和0中的最大值 ", math.Dim(1.1, 2.2))
	fmt.Println("    Mod(x,y)    取余运算 ", math.Mod(10, 3))
	fmt.Println("    Sqrt(x)     x的二次方根 ", math.Sqrt(9))
	fmt.Println("    Cbrt(x)     x的三次方根 ", math.Cbrt(27))
	fmt.Println("    Pow(x,y)    x的y次方 ", math.Pow(2,3))
	fmt.Println("    Pow10(x)    x的10次方 ", math.Pow10(2))
}