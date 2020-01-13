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
	var a []byte = []byte{1,2,3,4,5}
	var b []byte = []byte{2,3,4}

	fmt.Println("Equal(a,b)           判断两个切片的内容是否完全相同", "                   ", bytes.Equal(a, b))

	fmt.Println("Equal(s,t)           判断两个utf-8编码切片是否相同", "                    ", bytes.EqualFold(a, b))

	fmt.Println("Runes(s)             Runes函数返回和s等价的[]rune切片", "                 ", bytes.Runes(a))

	fmt.Println("HasPrefix(s,prefix)  判断s是否有前缀切片prefix","                        ", bytes.HasPrefix(a, b))
	fmt.Println("HasPrefix(s,suffix)  判断s是否有后缀切片prefix","                        ", bytes.HasSuffix(a, b))

	fmt.Println("Contains(a,b)        判断切片a是否包含子切片b", "                         ", bytes.Contains(a, b))

	fmt.Println("Count(s,sep)         Count计算s中有多少个不重叠的sep子切片。", "          ", bytes.Count(a, b))

	fmt.Println("Index(s,sep)         子切片sep在s中第一次出现的位置,不存在则返回-1", "    ", bytes.Index(a, b))

	fmt.Println("IndexByte(s,c)       字符c在s中第一次出现的位置,不存在则返回-1", "        ", bytes.IndexByte(a, 2))

	fmt.Println("IndexRune(s,r)       字符r在s中第一次出现的位置,不存在则返回-1", "        ", bytes.IndexRune(a, 2))

	fmt.Println("IndexAny(s,char)     字符c在s第一次出现位置,不存在/空字符串返回-1", "     ", bytes.IndexAny(a, "2"))

	fmt.Println("LastIndex(s,sep)     切片sep在字符串s最后一次出现位置,不存在返回-1", "    ", bytes.LastIndex(a, b))

	fmt.Println("Contains(a,b)        字符c在s最后一次出现的位置,不存在/空字符串返回-1"," ", bytes.LastIndexAny(a, "2"))

	fmt.Println("Repeat(a,count)      count个b串联形成的新的切片", "                       ", bytes.Repeat(a, 2))

	fmt.Println("Replace(s,old,new,n) s中前n个不重叠old都替换为new的切片拷贝,n<0全部替换", bytes.Replace(a, b, a, -1))
}