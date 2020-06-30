package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var t int

func init() {
	//1:服务端 2:客户端
	flag.IntVar(&t, "t", 0, "t 【1:服务端 2:客户端】")
	flag.Parse()
}

/*
 * 本质是编程设计的问题。
 * TCP本身就是流协议传输的。报文的切分与组装，更应该放到接收层处理。
 * 解决方案
 *	   关闭 Nagle 算法
 *     在报文传输增加结束符，比如{...}\t{...}
 *     报文增加信息，自己切分报文段，比如{"length":326, data:{...}}
 */
func main() {

	switch t {
	case 1:
		server()
	case 2:
		client()
	default:
		fmt.Println("input error: t 【1:服务端 2:客户端】")
	}
}

//go run Go/TCP/TCP粘包.go -t 1
func server() {

	//监听端口
	netListen, _ := net.Listen("tcp", ":9988")

	//端口关闭
	defer netListen.Close()

	for {

		//accept阻塞，等待请求
		conn, _ := netListen.Accept()

		go func() {
			buffer := make([]byte, 1024)
			for {
				n, err := conn.Read(buffer)
				if err != nil {
					return
				}

				conn.Write([]byte(time.Now().String()))

				fmt.Println(string(buffer[:n]))
			}
		}()
	}
}

//go run Go/TCP/TCP粘包.go -t 2
func client() {
	server := "127.0.0.1:9988"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", server)

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	defer conn.Close()

	fmt.Println("connect success")

	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write([]byte(words))
	}
}
