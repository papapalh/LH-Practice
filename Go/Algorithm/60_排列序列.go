package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getPermutation(9, 116907))
}

//思路
//	全排列思想
//耗时
//	超时
func getPermutation(n int, k int) string {

	res := []string{}

	used :=  map[int]bool{}

	xxx(&res, "", used, n, k)

	return res[k-1]
}

func xxx(res *[]string, node string, used map[int]bool, n int, k int) {

	if len(node) == n {
		*res = append(*res, node)
		return
	}

	if len(*res) > k {
		return
	}

	for i := 1; i <= n; i++ {

		if r, ok := used[i]; r == true && ok {
			continue
		}

		node = node + strconv.Itoa(i)
		used[i] = true

		xxx(res, node, used, n, k)

		node = node[:len(node)-1]
		used[i] = false
	}
}