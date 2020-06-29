package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//Once 只执行一次，由 sync.atomic 原子性控制
	//OnceFunc()

	//Mutex 互斥锁
	//MutexFunc()

	//RwMutex 读写锁
	//RwMutexFunc()

	//WaitGroup 处理一组线程的结束
	WaitGroupFunc()
}

func OnceFunc() {
	//Once 只执行一次，由 sync.atomic 原子性控制
	var once sync.Once
	once.Do(func() {
		fmt.Println("Only once")
	})
}

func MutexFunc() {
	var mu sync.Mutex
	for i := 0; i < 5; i++ {

		go func() {
			mu.Lock()
			fmt.Println("Lock")

			time.Sleep(time.Second)

			mu.Unlock()
			fmt.Println("UnLock")
		}()

	}

	time.Sleep(time.Second * 20)
}

func RwMutexFunc() {
	var rwMu sync.RWMutex

	for i := 0; i < 5; i++ {

		go func() {
			//读锁-允许多个读锁共存
			rwMu.RLock()

			fmt.Println("RLock")

			time.Sleep(time.Second)

			//释放读锁
			rwMu.RUnlock()
			fmt.Println("RUnlock")
		}()
	}

	time.Sleep(time.Second * 20)
}

func WaitGroupFunc() {

	//结构体为值传递，传入函数需要引用，请注意
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {

		wg.Add(1)

		go func() {
			time.Sleep(time.Second * 20)
			wg.Done()
			fmt.Println("WaitGroup")
		}()
	}

	//这个方法是堵塞的
	wg.Wait()
}
