package main

import (
	"fmt"
	"os"
)

func main() {

	f := "/Users/lihong/Desktop/Data/360/lihong/Doc/os/os.go"

	/*
	 * os 基本
	 */
	//主机名
	fmt.Println(os.Hostname())

	//当前工作目录
	fmt.Println(os.Getwd())

	//更改文件权限
	fmt.Println(os.Chmod(f, 420))

	//更改组
	// Chown(name string, uid, gid int) error

	//创建目录
	// os.Mkdir("name", os.FileMode)

	//文件移动
	// func Rename(oldpath, newpath string) error

	/*
	 * os 环境变量
	 */
	//系统全量环境变量
	// fmt.Println(os.Environ())

	//设置系统环境变量
	fmt.Println(os.Setenv("aa", "bb"))

	//获取指定系统环境变量
	fmt.Println(os.Getenv("aa"))

	/*
	 * FileInfo
	 */
	fileInfo, _ := os.Stat(f)
	fmt.Println(fmt.Sprintf("文件名称: %s \n 文件大小: %d \n 文件的模式位: %v \n 文件的修改时间: %v \n 是否是文件夹: %t",
		fileInfo.Name(),      // 文件的名字（不含扩展名）
		fileInfo.Size(),      // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		int(fileInfo.Mode()), // 文件的模式位
		fileInfo.ModTime(),   // 文件的修改时间
		fileInfo.IsDir(),     // 是否是文件夹
	))

	/*
	 * os 获取文件句柄
	 *	    描述符
	 * 		O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	 *		O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	 *		O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	 *		O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	 *		O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	 *		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	 *		O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	 *		O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	 */
	//创建文件(0666)。文件描述符: O_RDWR
	fileByCreate, _ := os.Create("/Users/lihong/Desktop/on-line/Practice/tmp.go")

	//打开一个文件用于读取。文件描述符: O_RDONLY
	fileByOpen, _ := os.Open(f)

	//指定模式打开文件。大多数调用者都应用Open或Create代替本函数
	fileByOpenFile, _ := os.OpenFile("/Users/lihong/Desktop/Data/360/lihong/Doc/main.go", os.O_RDWR, 0)

	//获取文件基本信息(FileInfo)
	fmt.Println(fileByOpenFile.Stat())

	//打开文件一定要释放
	defer func() {
		fileByCreate.Close()
		fileByOpen.Close()
		fileByOpenFile.Close()
	}()

	/*
	 * os 读
	 */
	//读取Byte字节数据
	r := make([]byte, 4)
	fileByOpen.Read(r)
	fmt.Println(string(r))

	//读取Byte字节数据(从偏移值off之后开始读)
	fileByOpen.ReadAt(r, 3)
	fmt.Println(string(r))

	/*
	 * os 写
	 */
	//写入Byte字节数据(从头写)
	fileByCreate.Write([]byte{97, 103, 101, 32})

	// //写入字符串数据(从头写)
	fileByCreate.WriteString("xxxx")

	//从偏移值off写入字节数据
	//fmt.Println(fileByOpenFile.WriteAt([]byte{97, 103, 101, 32}, 0))
}
