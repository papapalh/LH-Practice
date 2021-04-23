package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	/*
	 * Conn接口代表通用的面向流的网络连接。多个线程可能会同时调用同一个Conn的方法
	 */
	 // net.Conn{}

	/*
	 * Dial 连接地址 address
	 * 网络类型可以选择: "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
	 */
	net.Dial("tcp","127.0.0.1")

	// 客户端连接(增加超时)
	net.DialTimeout("tcp",":8080", time.Second * 10)

	//构建Dialer结构体连接
	c := net.Dialer{}
	c.Dial("tcp",":8080")

	/*
	 * Listener 监听地址 address
	 * 网络类型可以选择: "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
	 * type Listener interface {
     *     Addr() Addr //Addr返回该接口的网络地址
     *     Accept() (c Conn, err error) //Accept等待并返回下一个连接到该接口的连接
     *     Close() error //Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
     * }
	 */
	l,_ := net.Listen("tcp",":8080")

	//监听请求
	lc,_ := l.Accept()

	//读(结束为EOF)
	lc.Read([]byte{})

	//写
	lc.Write([]byte{})

	//关闭
	lc.Close()

}

//客户端(多次传输)
func client() {
	conn, _ := net.Dial("tcp", ":8080")

	//返回一个拥有 默认size 的reader，接收客户端输入
	reader := bufio.NewReader(os.Stdin)
	//缓存 conn 中的数据
	buf := make([]byte, 1024)

	fmt.Println("请输入客户端请求数据...")

	for {
		//客户端输入
		input, _ := reader.ReadString('\n')
		//去除输入两端空格
		input = strings.TrimSpace(input)
		//客户端请求数据写入 conn，并传输
		conn.Write([]byte(input+"\n"))
		//conn.Write([]byte(input))
		//服务器端返回的数据写入空buf
		//服务器端返回的数据写入空buf
		cnt, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("客户端读取数据失败 %s\n", err)
			continue
		}

		//回显服务器端回传的信息
		fmt.Print("服务器端回复" + string(buf[0:cnt]))
	}
}

//服务端(在传输的时候要注意粘包问题)
func server() {


	coon, _ := net.Listen("tcp", ":8080")

	for {

		l, err := coon.Accept()
		if err != nil {
			panic("建立连接失败:" + err.Error())
		}

		go handle(l)
	}
}

func handle(l net.Conn) {

	//defer l.Close()

	buf := make([]byte, 1, 4096)
	res := ""
	len := 0

	for {
		for {
			n, err := l.Read(buf[len:])
			if len > 0 {
				len++
			}

			fmt.Println(n,err,buf)
			if err != nil {
				l.Close() //关闭连接
				return
			}

			//EOF读取完毕
			if n == 0 {
				break
			}

			if string(buf[len:]) == "\n" {
				//读到/n就停止
				break
			}

			res += string(buf[len:])
		}
		l.Write([]byte("收到！" + res + "\n"))
	}
}