package main

import (
	"fmt"
	"time"
)

var poolNum int = 5
var poolChan chan struct{} = make(chan struct{}, poolNum)
var poolMap map[int]string = map[int]string{
	1: "未占用",
	2: "未占用",
	3: "未占用",
	4: "未占用",
	5: "未占用",
}

//chan 实现连接池
func main() {

	for i := 0; i < 10; i++ {
		//获取连接
		getConnct(i)

		time.Sleep(time.Second * 2)
	}
}

//获取连接
func getConnct(no int) {

	fmt.Println("===")

	fmt.Printf("等待:%d等待获取连接连接\n", no)

	//从连接池获取连接
	poolChan <- struct{}{}

	fmt.Printf("获取:%d获取连接连接\n", no)

	//循环连接池，找到一个能用的
	for n, p := range poolMap {
		if p == "未占用" {
			poolMap[n] = "已占用"
			fmt.Printf("占用:%d已占用%d号连接\n", no, n)
			break
		}
	}

	fmt.Println("===")

	//每隔20S释放一个链接
	go func() {
		time.Sleep(time.Second * 20)
		relaseConnect(no)
	}()
}

//释放连接
func relaseConnect(no int) {

	fmt.Println("===")

	//连接池释放一个空链接
	<-poolChan

	//循环连接池，释放一个链接(这里应该是传指针进来)
	for n, p := range poolMap {
		if p == "已占用" {
			poolMap[n] = "未占用"
			fmt.Printf("释放:%d已释放%d号连接\n", no, n)
			break
		}
	}

	fmt.Println("===")
}
