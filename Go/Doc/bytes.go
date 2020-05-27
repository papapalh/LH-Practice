package main

import (
	"bytes"
	"fmt"
)

const MinRead = 512

func main() {
	//byte = uint8
	//rune = int32
	//bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似。

	// Compare 返回一个整数表示两个[]byte切片按字典序比较的结果（类同C的strcmp）。如果a==b返回0；如果a<b返回-1；否则返回+1。nil参数视为空切片。
	var a []byte = []byte{1, 2, 3, 4, 5}
	var b []byte = []byte{2, 3, 4}

	fmt.Println("Equal(a,b)           判断两个切片的内容是否完全相同", "                   ", bytes.Equal(a, b))

	fmt.Println("Equal(s,t)           判断两个utf-8编码切片是否相同", "                    ", bytes.EqualFold(a, b))

	fmt.Println("Runes(s)             Runes函数返回和s等价的[]rune切片", "                 ", bytes.Runes(a))

	fmt.Println("HasPrefix(s,prefix)  判断s是否有前缀切片prefix", "                        ", bytes.HasPrefix(a, b))
	fmt.Println("HasPrefix(s,suffix)  判断s是否有后缀切片prefix", "                        ", bytes.HasSuffix(a, b))

	fmt.Println("Contains(a,b)        判断切片a是否包含子切片b", "                         ", bytes.Contains(a, b))

	fmt.Println("Count(s,sep)         Count计算s中有多少个不重叠的sep子切片。", "          ", bytes.Count(a, b))

	fmt.Println("Index(s,sep)         子切片sep在s中第一次出现的位置,不存在则返回-1", "    ", bytes.Index(a, b))

	fmt.Println("IndexByte(s,c)       字符c在s中第一次出现的位置,不存在则返回-1", "        ", bytes.IndexByte(a, 2))

	fmt.Println("IndexRune(s,r)       字符r在s中第一次出现的位置,不存在则返回-1", "        ", bytes.IndexRune(a, 2))

	fmt.Println("IndexAny(s,char)     字符c在s第一次出现位置,不存在/空字符串返回-1", "     ", bytes.IndexAny(a, "2"))

	fmt.Println("LastIndex(s,sep)     切片sep在字符串s最后一次出现位置,不存在返回-1", "    ", bytes.LastIndex(a, b))

	fmt.Println("Contains(a,b)        字符c在s最后一次出现的位置,不存在/空字符串返回-1", " ", bytes.LastIndexAny(a, "2"))

	fmt.Println("Repeat(a,count)      count个b串联形成的新的切片", "                       ", bytes.Repeat(a, 2))

	fmt.Println("Replace(s,old,new,n) s中前n个不重叠old都替换为new的切片拷贝,n<0全部替换", bytes.Replace(a, b, a, -1))

	fmt.Println("Reader")
	fmt.Println("    介绍：通过从一个[]byte读取数据，实现了io.Reader、io.Seeker、io.ReaderAt、io.WriterTo、io.ByteScanner、io.RuneScanner接口。")

	data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	read := make([]byte, 2)

	re := bytes.NewReader(data)
	fmt.Print("    NewReader(b) 通过[]byte创建 Reader                 ")
	fmt.Println(re)

	fmt.Print("    Len()        返回r包含的切片中还没有被读取的部分   ")
	fmt.Println(re.Len())

	fmt.Print("    Read(b)      读b字节数据                           ")
	fmt.Println(re.Read(read))

	fmt.Print("    Read(b)      读byte字节数据                        ")
	fmt.Println(re.ReadByte())

	fmt.Print("    ReadRune(b)  读4byte字节数据                       ")
	fmt.Println(re.ReadRune())

	fmt.Print("    UnreadByte() 撤销上一个读取的数据                  ")
	fmt.Println(re.UnreadByte())

	fmt.Println("Buffer")
	fmt.Println("    一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲")

	bf := bytes.NewBuffer(data)
	fmt.Print("    NewBuffer(b)   通过[]byte创建 Buffer               ")
	fmt.Println(bf)

	fmt.Println("    Reset()        重设缓冲，因此会丢弃全部内容，等价于b.Truncate(0)          ")
	//bf.Reset()

	fmt.Print("    Len()          返回缓冲中未读取部分的字节长度；b.Len() == len(b.Bytes())   ")
	fmt.Println(bf.Len())

	fmt.Print("    Bytes()        返回未读取部分字节数据的切片；len(b.Bytes()) == b.Len()     ")
	fmt.Println(bf.Bytes())
	//fmt.Println("    	注意：如果中间没有调用其他方法，修改返回的切片的内容会直接改变Buffer的内容。")

	fmt.Print("    String()       将未读取部分的字节数据作为字符串返回        ")
	fmt.Println(bf.String())

	fmt.Print("    Read(b)        读取b字节数据      ")
	fmt.Println(bf.Read(read))
	//fmt.Println("    	注意：返回值n是读取的字节数，除非缓冲中完全没有数据可以读取并写入p，此时返回值err为io.EOF；否则err总是nil。")

	fmt.Print("    ReadByte()     读取byte字节数据   ")
	fmt.Println(bf.ReadByte())

	fmt.Print("    Next()         返回未读取部分前n字节数据的切片，并且移动读取位置  ")
	fmt.Println(bf.Next(1))

	fmt.Print("    UnreadByte()   撤销上一个读取byte的数据               ")
	fmt.Println(bf.UnreadByte())

	fmt.Print("    ReadRune()     读取并返回缓冲中的下一个utf-8码值。    ")
	fmt.Println(bf.ReadRune())

	fmt.Print("    UnreadRune()   撤销上一个读取Rune的数据               ")
	fmt.Println(bf.UnreadRune())

	fmt.Print("    ReadBytes()    连续读取至 第一次遇到delime字节 结束   ")
	fmt.Println(bf.ReadBytes(5))

	fmt.Print("    ReadString()   连续读取至 第一次遇到delime字节 结束   ")
	fmt.Println(bf.ReadString(7))

	fmt.Print("    Write(p)       将p的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p),缓冲过大会引发painc  ")
	fmt.Println(bf.Write([]byte{10, 11, 12}))

	fmt.Print("    WriteString(s) 将s的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p),缓冲过大会引发painc  ")
	fmt.Println(bf.WriteString("1213141516"))

	fmt.Print("    WriteByte(b)   将b的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p),缓冲过大会引发painc  ")
	fmt.Println(bf.WriteByte(1))

	fmt.Print("    WriteRune(r)   将r的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p),缓冲过大会引发painc  ")
	fmt.Println(bf.WriteRune(1))

	fmt.Println("    ReadFrom(io) ")
	//fmt.Println(bf.ReadFrom(1))

	fmt.Println("    WriteTo(io) ")
	//fmt.Println(bf.WriteTo(1))

	fmt.Println(bf.Bytes())
}
