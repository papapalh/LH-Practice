package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (s intSlice) Len() int { return len(s) }
func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	sort.Sort(a)
	fmt.Println(a)
}
