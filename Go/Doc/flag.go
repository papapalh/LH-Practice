package main

import (
	"flag"
	"fmt"
)

func main() {

	//第一种写法
	var iInt          = flag.Int("int", 1, "int")
	var iIntSixFour   = flag.Int64("int64", 2, "int64")
	var iUint         = flag.Uint("uint", 3, "uint")
	var iUintSixFour  = flag.Uint64("uint64", 4, "uint64")
	var iFloatSixFour = flag.Float64("float64", 5, "float64")
	var iString       = flag.String("string", "string", "string")
	var iBool         = flag.Bool("bool", false, "bool")

	//第二种写法
	var tInt int
	var tInt64 int64
	var tUInt uint
	var tUInt64 uint64
	var tFloat64 float64
	var tString string
	var tBool bool

	flag.IntVar(&tInt, "tInt", 1, "tInt")
	flag.Int64Var(&tInt64, "tInt64", 1, "tInt64")
	flag.UintVar(&tUInt, "tUInt", 1, "tUInt")
	flag.Uint64Var(&tUInt64, "tUInt64", 1, "tUInt64")
	flag.Float64Var(&tFloat64, "tFloat64", 1, "tFloat64")
	flag.StringVar(&tString, "tString", "tString", "tString")
	flag.BoolVar(&tBool, "tBool", true, "tBool")

	flag.Parse()

	fmt.Println("基本语法 【go run flag.go -name value】")
	fmt.Println("    Parse()                   解析传入命令参数(必须解析)")
	fmt.Println("    Int(name,value,usage)     解析获取一个int     ", *iInt)
	fmt.Println("    Int64(name,value,usage)   解析获取一个int64   ", *iIntSixFour)
	fmt.Println("    Uint(name,value,usage)    解析获取一个uint    ", *iUint)
	fmt.Println("    Uint64(name,value,usage)  解析获取一个uint64  ", *iUintSixFour)
	fmt.Println("    Float64(name,value,usage) 解析获取一个float64 ", *iFloatSixFour)
	fmt.Println("    String(name,value,usage)  解析获取一个string  ", *iString)
	fmt.Println("    Bool(name,value,usage)    解析获取一个bool    ", *iBool)

	fmt.Println("")
	fmt.Println("    Int(name,value,usage)     解析获取一个int     ", tInt)
	fmt.Println("    Int64(name,value,usage)   解析获取一个int64   ", tInt64)
	fmt.Println("    Uint(name,value,usage)    解析获取一个uint    ", tUInt)
	fmt.Println("    Uint64(name,value,usage)  解析获取一个uint64  ", tUInt64)
	fmt.Println("    Float64(name,value,usage) 解析获取一个float64 ", tFloat64)
	fmt.Println("    String(name,value,usage)  解析获取一个string  ", tString)
	fmt.Println("    Bool(name,value,usage)    解析获取一个bool    ", tBool)
}