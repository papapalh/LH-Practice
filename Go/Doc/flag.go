package main

import (
	"flag"
	"fmt"
)

func main() {
	var iInt          = flag.Int("int", 1234, "int")
	var iIntSixFour   = flag.Int64("int64", 1234, "int64")
	var iUint         = flag.Uint("uint", 1234, "uint")
	var iUintSixFour  = flag.Uint64("uint64", 1234, "uint64")
	var iFloatSixFour = flag.Float64("float64", 1234, "float64")
	var iString       = flag.String("string", "string", "string")
	var iBool         = flag.Bool("bool", false, "bool")

	flag.Parse()

	fmt.Println("基本语法 【go run flag.go -name value】")
	fmt.Println("    Parse()                   解析传入命令参数(必须解析)")
	fmt.Println("    Int(name,value,usage)     解析获取一个int    ", *iInt)
	fmt.Println("    Int64(name,value,usage)   解析获取一个int64  ", *iIntSixFour)
	fmt.Println("    Uint(name,value,usage)    解析获取一个Uint   ", *iUint)
	fmt.Println("    Uint64(name,value,usage)  解析获取一个Uint64 ", *iUintSixFour)
	fmt.Println("    Float64(name,value,usage) 解析获取一个int64  ", *iFloatSixFour)
	fmt.Println("    String(name,value,usage)  解析获取一个int64  ", *iString)
	fmt.Println("    Bool(name,value,usage)    解析获取一个int64  ", *iBool)
}