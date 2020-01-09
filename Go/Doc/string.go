package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("EqualFold(s,t)       判断两个字符串(utf-8)是否相同(忽略大小写) ", strings.EqualFold("x","X"))
	fmt.Println("HasPrefix(s,prefix)  判断s是否有前缀字符串prefix               ", strings.HasPrefix("xxxX","xxx"))
	fmt.Println("HasSuffix(s,prefix)  判断s是否有后缀字符串prefix               ", strings.HasSuffix("Xxxx","xxx"))
	fmt.Println("Contains(s,substr)   判断字符串s是否包含子串substr             ", strings.Contains("xxxx","x"))
	fmt.Println("Count(s,substr)      返回字符串s中有几个不重复的sep子串        ", strings.Count("xxxX","x"))
	fmt.Println("Index(s,substr)      sep在s第一次出现的位置,不存在返回-1       ", strings.Index("xxxX","xxx"))
	fmt.Println("LastIndex(s,substr)  sep在s最后一次出现的位置,不存在返回-1     ", strings.LastIndex("xxxX","xxx"))
	fmt.Println("ToLower(s)           字母都转为对应的小写版本的拷贝            ", strings.ToLower("xxxX"))
	fmt.Println("ToUpper(s)           字母都转为对应的大写版本的拷贝            ", strings.ToUpper("xxxX"))
	fmt.Println("Repeat(s,count)      返回count个s串联的字符串                  ", strings.Repeat("abcd", 3))
	fmt.Println("Replace(s,old,new,n) 将s中前n个old替换为new,n<0会替换所有old   ", strings.Replace("xxxX", "x", "a", -1))
	fmt.Println("Trim(s,cutset)       将s前后端所有cutset都去掉                 ", strings.Trim(" xxxX ", " "))
	fmt.Println("TrimSpace(s)         将s前后端所有空白去掉                     ", strings.TrimSpace(" xxxX "))
	fmt.Println("TrimLeft(s,cutset)   将s左端所有cutset都去掉                   ", strings.TrimLeft(" xxxX", " "))
	fmt.Println("TrimRight(s,cutset)  将s右端所有cutset都去掉                   ", strings.TrimRight("xxxX ", " "))
	fmt.Println("Split(s,sep)         将s以sep方式切割                          ", strings.Split("x-x-x-x", "-"))
	fmt.Println("Split(a,sep)         将a以sep方式组合                          ", strings.Join([]string{"a","b","a","b"}, "-"))
}