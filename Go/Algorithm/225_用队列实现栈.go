package main

func main() {}

/**
 * 题目
 *    使用队列实现栈的下列操作：
 *    push(x) -- 元素 x 入栈
 *    pop() -- 移除栈顶元素
 *    top() -- 获取栈顶元素
 *    empty() -- 返回栈是否为空
 * 注意:
 *    你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
 *    你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
 *    你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
 * 思路
 *    记住栈的特性就行，栈的问题没什么好说的了
 * 耗时
 *    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 *    内存消耗：2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
type MyStack struct {
	normal []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.normal = append(this.normal, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	tmp := this.normal[len(this.normal)-1]

	this.normal = this.normal[0 : len(this.normal)-1]

	return tmp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.normal[len(this.normal)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.normal) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
