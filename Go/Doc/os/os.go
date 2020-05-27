package main

import (
	"fmt"
	"os"
)

func main () {
	fmt.Print("Hostname()       主机名", ); fmt.Println(os.Hostname())
	fmt.Println("Environ()        系统环境变量", ); // fmt.Println(os.Environ())
	fmt.Print("Getenv()         检索系统环境变量", ); fmt.Println(os.Getenv("GOPATH"))
	fmt.Print("Setenv()         设置系统环境变量", ); fmt.Println(os.Setenv("aa", "bb"))


	file,_ := os.Open("/Users/lihong/Desktop/Data/360/lihong/Doc/io.go")
	fmt.Print("Open(n)          Open打开一个文件用于读取(只读模式)"); fmt.Println(file)

	openFile,_ := os.OpenFile("/Users/lihong/Desktop/Data/360/lihong/Doc/main.go", os.O_RDWR, 0)
	fmt.Print("OpenFile(n,f,p)  一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。"); fmt.Println(openFile)

	fmt.Print("Name()           文件名称"); fmt.Println(file.Name())

	fmt.Println("Create(n)        采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件"); // fmt.Println(os.Create("/Users/lihong/Desktop/Data/360/lihong/Doc/file.go"))

	r := make([]byte, 4)
	fmt.Print("Read(b)          读取Byte字节数据"); file.Read(r); fmt.Println(string(r))

	ra := make([]byte, 4)
	fmt.Print("ReadAt(b, off)   从偏移值off读取Byte字节数据"); file.ReadAt(ra, 0); fmt.Println(string(ra))

	w := []byte{97, 103, 101, 32}
	fmt.Print("Write(b)         写入Byte字节数据(从头写)"); fmt.Println(openFile.Write(w))

	fmt.Print("WriteString(b)   写入字符串数据(从头写)"); fmt.Println(openFile.WriteString("xxxx"))

	fmt.Print("WriteAt(b, off)  从偏移值off写入字节数据"); fmt.Println(openFile.WriteAt([]byte{97, 103, 101, 32}, 1))

	defer func() {
		file.Close()
		openFile.Close()
	}()

}