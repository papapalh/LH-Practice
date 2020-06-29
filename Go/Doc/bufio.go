package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio包实现了有缓冲的I/O。
//它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
func main() {

	/*
	 * Reader
	 */
	fileReader, _ := os.Open("/Users/lihong/Desktop/lua.txt")
	defer func() {
		fmt.Println("defer1")
		fileReader.Close()
	}()

	//默认大小缓冲、从r读取的*Reader。
	bufioReader := bufio.NewReader(fileReader)

	//读指定字节
	r := make([]byte, 2)
	bufioReader.Read(r)
	fmt.Println(r)

	//的丢弃所有缓冲数据, 重设为其下层从r读取数据。
	bufioReader.Reset(fileReader)

	//读一个字节
	fmt.Println(bufioReader.ReadByte())

	//撤回上一个字节读取(不可以连续调用，会报错)
	fmt.Println(bufioReader.UnreadByte())

	//读一个UTF-8
	fmt.Println(bufioReader.ReadRune())

	//撤销一个UFT-8
	fmt.Println(bufioReader.UnreadRune())

	//读取行(应该使用 ReadBytes("\n") / ReadString("\n"))
	fmt.Println(bufioReader.ReadLine())

	//读取直到第一次遇到delim字节
	fmt.Println(bufioReader.ReadBytes(116))

	//读取直到第一次遇到delim字节
	fmt.Println(bufioReader.ReadString(101))

	/*
	 * Write
	 */
	fileWrite, _ := os.Create("/Users/lihong/Desktop/luaWrite.txt")
	defer func() {
		fmt.Println("defer2")
		fileWrite.Close()
	}()

	//默认大小缓(4096)、从w写的*Reader。
	bufioWrite := bufio.NewWriter(fileWrite)

	//指定缓冲区大小的Buffer
	// bufioWrite := bufio.NewWriterSize(fileWrite, 8192)

	//Reset丢弃缓冲中的数据，清除任何错误，将b重设为将其输出写入w。
	bufioWrite.Reset(fileWrite)

	//写指定字节
	bufioWrite.Write([]byte{116, 116})

	//写单字节
	bufioWrite.WriteByte(116)

	//写入UTF8
	bufioWrite.WriteRune(116)

	//缓冲区已使用字节
	fmt.Println(bufioWrite.Buffered())

	//缓冲区未使用字节
	fmt.Println(bufioWrite.Available())

	//Flush方法将缓冲中的数据写入下层的io.Writer接口。
	//超过缓冲区大小的，不需要调用Flush会自动写入
	bufioWrite.Flush()
}
