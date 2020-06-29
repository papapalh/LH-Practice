package main

import (
	"bytes"
	"fmt"
)

//bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似。
//byte = uint8
//rune = int32
func main() {
	/*
	 * Byte基本方法
	 */
	//切片是否完全相同
	fmt.Println(bytes.Equal([]byte{1}, []byte{1}))

	//切片UTF8是够相同(将unicode大写、小写、标题三种格式字符视为相同)
	fmt.Println(bytes.EqualFold([]byte{1}, []byte{1}))

	//Byte -> Rune
	fmt.Println(bytes.Runes([]byte{1}))

	/*
	 * Byte查找
	 */
	s := []byte{1, 2, 3, 4, 5}
	sep := []byte{1, 2}
	//判断前缀切片sep
	fmt.Println(bytes.HasPrefix(s, sep))

	//判断后缀切片sep
	fmt.Println(bytes.HasSuffix(s, sep))

	//判断子切片sep
	fmt.Println(bytes.Contains(s, sep))

	//查找子切片个数sep
	fmt.Println(bytes.Count(s, sep))

	//子切片sep在s中第一次出现的位置,不存在则返回-1
	bytes.Index(s, sep)

	//子切片sep最后一次出现位置, 不存在返回-1
	fmt.Println(bytes.LastIndex(s, sep))

	//字符99在s中第一次出现的位置,不存在则返回-1
	fmt.Println(bytes.IndexByte(s, 99))

	/*
	 * Byte 操作
	 */
	//串联count个s
	fmt.Println(bytes.Repeat(s, 2))

	/*
	 * Byte替换
	 */
	//Byte替换 Replace(s,old,new,n) s中前n个不重叠old都替换为new的切片拷贝,n<0全部替换
	fmt.Println(bytes.Replace(s, []byte{1}, []byte{2}, -1))

	/*
	 * Byte Reader
	 */
	n := make([]byte, 1)
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	//构建Reader
	byteRead := bytes.NewReader(data)

	//未被读取的部分
	fmt.Println(byteRead.Len())

	//读取 n Byte
	byteRead.Read(n)
	fmt.Println(n)

	//读取 byte
	fmt.Println(byteRead.ReadByte())

	//撤销读取的上一个Byte数据(也就是说下一个读取的还是这个)
	fmt.Println(byteRead.UnreadByte())

	//读取 Rune(4字节)(主要读取是utf8字符串)
	fmt.Println(byteRead.ReadRune())

	/*
	 * Byte Buffer Reader
	 * 一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲
	 */
	//构建Buffer
	byteBuffer := bytes.NewBuffer(data)

	//重设缓冲，因此会丢弃全部内容，等价于b.Truncate(0)
	//byteBuffer.Reset()

	//重设缓冲，丢弃缓冲中除前n字节数据外的其它数据，如果n小于零或者大于缓冲容量将panic。
	// b.Truncate(0)

	//缓冲未被读取的部分
	fmt.Println(byteBuffer.Len())

	//返回未读取部分字节数据的切片
	//	注意：如果中间没有调用其他方法，修改返回的切片的内容会直接改变Buffer的内容。
	fmt.Println(byteBuffer.Bytes())

	//将未读取部分的字节数据作为字符串返回
	fmt.Println(byteBuffer.String())

	//读取 n Byte
	byteBuffer.Read(n)
	fmt.Println(n)

	//读取 byte
	fmt.Println(byteBuffer.ReadByte())

	//撤销读取的上一个Byte数据(也就是说下一个读取的还是这个)
	// fmt.Println(byteBuffer.UnreadByte())

	//读取 n byte 数据
	fmt.Println(byteBuffer.Next(2))

	//读取 Rune(4字节)(主要读取是utf8字符串)
	fmt.Println(byteBuffer.ReadRune())

	//撤销读取的上一个Rune数据(也就是说下一个读取的还是这个)
	// fmt.Println(byteBuffer.UnreadRune())

	//连续读取字节至delim停止
	fmt.Println(byteBuffer.ReadBytes(10))

	//连续读取字节至delim停止(返回字符串)
	fmt.Println(byteBuffer.ReadString(12))

	/*
	 * Byte Buffer Write
	 * 一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲
	 */

	//写[]Byte缓冲，返回值n为len(p),缓冲过大会引发painc.
	fmt.Println(byteBuffer.Write([]byte{21}))

	//写[]Rune缓冲，返回值n为len(p),缓冲过大会引发painc.
	fmt.Println(byteBuffer.WriteString("哈"))

	//写Byte缓冲，返回值n为len(p),缓冲过大会引发painc.
	fmt.Println(byteBuffer.WriteByte(22))

	//写Rune缓冲，返回值n为len(p),缓冲过大会引发painc.
	fmt.Println(byteBuffer.WriteRune(23))

	//r中读取数据直到结束并将读取的数据写入缓冲中，如必要会增加缓冲容量。
	// byteBuffer.ReadFrom(r)

	//WriteTo从缓冲中读取数据直到缓冲内没有数据或遇到错误，并将这些数据写入w
	//fmt.Println(bf.WriteTo(w)
}
