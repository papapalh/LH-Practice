package main

func main() {
	guessNumber(5)
}

/**
 * 题目
 *    猜数字游戏的规则如下：
 *    每轮游戏，系统都会从 1 到 n 随机选择一个数字。 请你猜选出的是哪个数字。
 *    如果你猜错了，系统会告诉你，你猜测的数字比系统选出的数字是大了还是小了。
 *    你可以通过调用一个预先定义好的接口 guess(int num) 来获取猜测结果，返回值一共有 3 种可能的情况（-1，1 或 0）：
 *        -1 : 你猜测的数字比系统选出的数字大
 *         1 : 你猜测的数字比系统选出的数字小
 *         0 : 恭喜！你猜对了！
 * 示例 :
 *    输入: n = 10, pick = 6
 *    输出: 6
 * 思路
 *    二分法(同278题)
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func guessNumber(n int) int {
	min := 0

	hei := n

	for hei >= min {
		middle := (hei-min)/2 + min

		//验证结果
		res := guess(middle)
		if res == 0 {
			return middle
		} else if res == -1 {
			hei = middle - 1
		} else if res == 1 {
			min = middle + 1
		}
	}

	return 0
}

func guess(version int) int {
	return 0
}
