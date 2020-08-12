package main

import "fmt"

/**
 * 题目
 *    三合一。描述如何只用一个数组来实现三个栈。
 *    你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。
 *    构造函数会传入一个stackSize参数，代表每个栈的大小。
 * 示例1:
 *    输入：
 *        ["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
 *        [[1], [0, 1], [0, 2], [0], [0], [0], [0]]
 *    输出：
 *        [null, null, null, 1, -1, -1, true]
 *        说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
 * 示例2:
 *    输入：
 *        [ *"TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
 *        [[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
 *    输出：
 *        [null, null, null, null, 2, 1, -1, -1]
 * 题解
 *    二维数据，管理栈
 * 耗时
 *    执行用时：100 ms, 在所有 Go 提交中击败了50.67%的用户
 *    内存消耗：17.4 MB, 在所有 Go 提交中击败了33.33%的用户
 */

func main() {
	s := Constructor(2)

	s1 := &s

	s1.Push(0, 1)
	s1.Push(0, 2)
	s1.Push(0, 3)
	fmt.Println(s1.Pop(0))
	fmt.Println(s1.Pop(0))
	fmt.Println(s1.Pop(0))
	fmt.Println(s1.Peek(0))

	fmt.Println(s1)
}

type TripleInOne struct {
	stackSize int
	stack     map[int][]int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stackSize: stackSize,
		stack:     map[int][]int{},
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {

	//初始化
	if this.stack[stackNum] == nil {
		this.stack[stackNum] = make([]int, 0)
	}

	//栈满不推
	if len(this.stack[stackNum]) >= this.stackSize {
		return
	}

	//推数据
	this.stack[stackNum] = append(this.stack[stackNum], value)
}

func (this *TripleInOne) Pop(stackNum int) int {
	//栈空-1
	if this.IsEmpty(stackNum) {
		return -1
	}

	tmp := this.stack[stackNum][len(this.stack[stackNum])-1]

	//处理数据
	this.stack[stackNum] = this.stack[stackNum][0 : len(this.stack[stackNum])-1]

	return tmp
}

func (this *TripleInOne) Peek(stackNum int) int {

	//栈空-1
	if this.IsEmpty(stackNum) {
		return -1
	}

	return this.stack[stackNum][len(this.stack[stackNum])-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return len(this.stack[stackNum]) == 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
