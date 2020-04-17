package main

import (
	"fmt"
)

func main() {
	//
	////runtime 相关
	//
	////GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置(P的值)
	//// 若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。
	//runtime.GOMAXPROCS(1)
	//
	////Goexit终止调用它的go程。其它go程不会受影响。
	////Goexit会在终止该go程前执行所有defer的函数。
	////在程序的main go程调用本函数，会终结该go程，而不会让main返回(程序会崩溃)
	////runtime.Goexit()
	//
	////Gosched使当前go程放弃处理，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。
	//runtime.Gosched()
	//
	////NumGoroutine返回当前存在的Go程数(可调度运行的)
	//runtime.NumGoroutine()
	//
	////SetMaxStack设置该以被单个go程调用栈可使用的内存最大值。
	////如果任何go程在增加其调用栈时超出了该限制，程序就会崩溃。SetMaxStack返回之前的设置。默认设置在32位系统是250MB，在64位系统是1GB。
	////SetMaxStack主要用于限制无限递归的go程带来的灾难。 它只会限制未来增长的调用栈。
	//debug.SetMaxStack(100)
	//
	////SetMaxThreads设置go程序可以使用的最大操作系统线程数。如果程序试图使用超过该限制的线程数，就会导致程序崩溃。SetMaxThreads返回之前的设置，初始设置为10000个线程。
	////该限制控制操作系统线程数，而非go程数。go程序只有在一个go程准备要执行，但现有的线程都阻塞在系统调用、cgo调用或被runtime.LockOSThread函数阻塞在其他go程时，才会创建一个新的线程。
	////SetMaxThreads主要用于限制程序无限制的创造线程导致的灾难。目的是让程序在干掉操作系统之前，先干掉它自己。
	//debug.SetMaxThreads(1000)

	//垃圾回收相关


	a := make(chan int)
	b := make(chan int)

	fmt.Println(cap(a))

	go func() {
		a <- 2
		close(a)
	}()

	go func() {
		b <- 3
		close(b)
	}()

	select {
	case x := <-a:
		fmt.Println(x)
	case z := <-b:
		fmt.Println(z)
	//default:
	//	fmt.Println("default")
	}
}