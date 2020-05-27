package main

import (
	"fmt"
	"net"
)

func main() {
	tcpServer, _ := net.ResolveTCPAddr("tcp4", "10.17.59.240:8080")
	listener, _ := net.ListenTCP("tcp", tcpServer)

	for {
		//当有新的客户端请求来的时候，拿到与客户端的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//接受客户端消息
		buf := make([]byte, 100)
		conn.Read(buf)
		fmt.Println(string(buf))

		//向客户端发送消息
		conn.Write([]byte("server tcp connect success."))

		conn.Close()
	}
}