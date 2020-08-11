package main

/**
实现一个MyQueue类，该类用两个栈来实现一个队列。
示例：
	MyQueue queue = new MyQueue();
	queue.push(1);
	queue.push(2);
	queue.peek();  // 返回 1
	queue.pop();   // 返回 1
	queue.empty(); // 返回 false
说明：
	你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size 和 is empty 操作是合法的。
	你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
	假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
思路
	正常链表
耗时
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了15.38%的用户
*/
type MyQueue struct {
	normal []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.normal = append(this.normal, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	tmp := this.normal[0]

	this.normal = this.normal[1:]

	return tmp
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.normal[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.normal) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
