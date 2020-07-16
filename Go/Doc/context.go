package main

import (
	"context"
	"fmt"
	"time"
)

//使用原则
//	不要把Context存在一个结构体当中，显式地传入函数。Context变量需要作为第一个参数使用，一般命名为ctx；
//  即使方法允许，也不要传入一个nil的Context，如果你不确定你要用什么Context的时候传一个context.TODO；
//  使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数；
//  同样的Context可以用来传递到不同的goroutine中，Context在多个goroutine中是安全的；
//  在子Context被传递到的goroutine中，应该对该子Context的Done信道（channel）进行监控，一旦该信道被关闭（即上层运行环境撤销了本goroutine的执行），应主动终止对当前请求信息的处理，释放资源并返回。
func main () {

	/*
	 * Context 内容
	 *     Value 共享数据
	 *     Done  信道（channel），当Context被撤销或过期时，该信道是关闭的，即它是一个表示Context是否已关闭的信号。
	 */

	//创建根节点
	ctx := context.Background()

	/*
	 * 共享数据 Value
	 */
	//创建一个共享数据的子节点(会继承父节点的所有属性)
	//数据的获得是协程安全
	//使用这些数据的时候要注意同步(map,[]slice)等，要上锁
	valueSon1 := context.WithValue(ctx, "son1", "son2")
	valueSon2 := context.WithValue(valueSon1, "son2", "son2")

	//打印调用链->context.Background.WithValue(type string, val son2).WithValue(type string, val son2)
	fmt.Println(fmt.Sprintf("%+v", valueSon2))

	//这里Son2 继承了 son1
	fmt.Println(valueSon2.Value("son1"), valueSon2.Value("son2"))

	/*
	* 撤销 Cancel
	* 每个都会返回一个撤销函数，执行撤销函数，该结点的Context被撤销
	*/
	cancel1, func1 := context.WithCancel(ctx)
	_, func2 := context.WithCancel(cancel1)

	//撤销-当上层关闭，那么下层也应该关闭
	go func(f context.CancelFunc) {
		//10s后关闭上层
		time.Sleep(time.Second * 10)
		f()
	}(func1)

	if _, ok := <-cancel1.Done(); ok {
		func2()
	}

	fmt.Println("二层关闭")

	/*
	 * 撤销 TimeOut
	 * 每个都会返回一个撤销函数，执行撤销函数，该结点的Context被撤销
	 * WithTimeout函数与WithDeadline类似，只不过它传入的是从现在开始Context剩余的生命时长。
	 * 他们都同样也都返回了所创建的子Context的控制权，一个CancelFunc类型的函数变量
	 */
	//ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	//defer cancel()
	//
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	//}

	/*
	 * 过期撤销 deadline
	 * WithDeadline函数的作用也差不多，它返回的Context类型值同样是parent的副本，但其过期时间由deadline和parent的过期时间共同决定。
	 * 当parent的过期时间早于传入的deadline时间时，返回的过期时间应与parent相同。
	 * 父节点过期时，其所有的子孙节点必须同时关闭；反之，返回的父节点的过期时间则为deadline。
	 */
	//d := time.Now().Add(50 * time.Millisecond)
	//ctx, cancel := context.WithDeadline(context.Background(), d)



}

