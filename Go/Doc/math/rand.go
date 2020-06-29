package main

import (
	"fmt"
	"math/rand"
)

//伪随机数生成器
//每次都会产生确定的序列
func main() {

	//为避免每次产生同样的序列，可以用Seed初始化
	rand.Seed(2)

	//非负的伪随机int值
	fmt.Println(rand.Int())

	//返回一个int32类型的非负的31位伪随机数。
	fmt.Println(rand.Int31())

	//返回一个int64类型的非负的63位伪随机数。
	fmt.Println(rand.Int63())

	//返回一个uint32类型的非负的32位伪随机数。
	fmt.Println(rand.Uint32())

	//返回一个取值范围在[0.0, 1.0)的伪随机float32值。
	fmt.Println(rand.Float32())

	//返回一个取值范围在[0.0, 1.0)的伪随机float64值。
	fmt.Println(rand.Float64())

	//返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
	fmt.Println(rand.Intn(10))
}
