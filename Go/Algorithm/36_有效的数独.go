package main

import (
	"fmt"
)

func main() {

	a := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}

	fmt.Println(isValidSudoku(a))
}

//思路
//	HASH 表为主要解法, 通过判断 行|列|3*3格子的重复性来判断数据正确
//  最重要的格子算法 (row/3)*3+column/3
//耗时
//	执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//	内存消耗： 4.5 MB , 在所有 Go 提交中击败了 7.21% 的用户
func isValidSudoku(board [][]byte) bool {

	rowDict := map[int]map[string]struct{}{}
	columnDict := map[int]map[string]struct{}{}
	block33 := map[int]map[string]struct{}{}

	for row, byteItem := range board {

		for column, value := range byteItem {

			if string(value) == "." {
				continue
			}

			//判断行
			if _, ok := rowDict[row][string(value)]; ok {
				return false
			}
			if rowDict[row] == nil {
				rowDict[row] = make(map[string]struct{})
			}
			rowDict[row][string(value)] = struct{}{}

			//判断列
			if _, ok := columnDict[column][string(value)]; ok {
				return false
			}
			if columnDict[column] == nil {
				columnDict[column] = make(map[string]struct{})
			}
			columnDict[column][string(value)] = struct{}{}

			//3*3格子
			if _, ok := block33[(row/3)*3+column/3][string(value)]; ok {
				return false
			}
			if block33[(row/3)*3+column/3] == nil {
				block33[(row/3)*3+column/3] = make(map[string]struct{})
			}
			block33[(row/3)*3+column/3][string(value)] = struct{}{}

		}
	}

	return true
}