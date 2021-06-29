package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

//思路
//	哈希表 + 排序
//	如果单词是不重复的可以用位来做
//耗时
//	执行用时： 32 ms , 在所有 Go 提交中击败了 45.02% 的用户
//	内存消耗： 8.2 MB , 在所有 Go 提交中击败了 46.30% 的用户
func groupAnagrams(strs []string) [][]string {

	bitMap := map[string][]string{}

	if len(strs) == 0 {
		return [][]string{}
	}

	for _, str := range strs {

		k := strings.Split(string(str), "")

		sort.Strings(k)

		ks := strings.Join(k, "")

		bitMap[ks] = append(bitMap[ks], string(str))
	}

	res := make([][]string, 0)

	for _, m := range bitMap {
		res = append(res, m)
	}

	return res
}