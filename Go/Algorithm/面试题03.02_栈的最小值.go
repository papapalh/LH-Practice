package main

import "fmt"

/**
 * 请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。
 * 示例：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 */

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {

	stack := Constructor()

	stack1 := &stack

	stack1.Push(-2)
	stack1.Push(0)
	stack1.Push(-3)
	fmt.Println(stack1.GetMin())
	stack1.Pop()
	fmt.Println(stack1.Top())
	fmt.Println(stack1.GetMin())

	fmt.Println(stack1)

}

/**
 * 第一种方法 (双栈)
 *    方法
 *        一个栈(栈1)保存正常数据，另一个栈(栈2)保存最小的数据。
 *        两个栈入栈和出栈都是同时的。
 *    示例
 *        push 3,6,9,4,1
 *            3 栈1为3         栈2为3
 *            6 栈1为3,6       栈2为3,3
 *            9 栈1为3,6,9     栈2为3,3,3
 *            4 栈1为3,6,9,4   栈2为3,3,3,3
 *            1 栈1为3,6,9,5,1 栈2为3,3,3,3,1
 * 第二种方法 (单栈)
 *    方法
 *        一个栈保存最小和正常数据，也就是说，一次出栈/入栈都是两个元素。
 *        push 3,6,9,4,1
 *            3 栈1为3,3
 *            6 栈1为3,3,6,3
 *            9 栈1为3,6,6,3,9,3
 *            4 栈1为3,6,6,3,9,3,4,3
 *            4 栈1为3,6,6,3,9,3,4,3,1,1
 * 题解使用双栈(操作切片)
 *     执行用时：24 ms, 在所有 Go 提交中击败了15.09%的用户
 *     内存消耗：7.2 MB, 在所有 Go 提交中击败了100.00%的用户
 */

type MinStack struct {
	normal []int
	min    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {

	//入最小栈
	if len(this.min) != 0 && x > this.min[len(this.min)-1] {
		this.min = append(this.min, this.min[len(this.min)-1])
	} else {
		this.min = append(this.min, x)
	}

	this.normal = append(this.normal, x)
}

//出栈
func (this *MinStack) Pop() {
	//处理栈(删除数组最后一个元素)
	this.normal = this.normal[:len(this.normal)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.normal[len(this.normal)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) != 0 {
		return this.min[len(this.min)-1]
	}

	return 0
}
