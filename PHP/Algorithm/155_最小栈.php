<?php

/**
 * 题目
 *      设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 *      push(x) -- 将元素 x 推入栈中。
 *      pop() -- 删除栈顶的元素。
 *      top() -- 获取栈顶元素。
 *      getMin() -- 检索栈中的最小元素。
 * 示例:
 *      MinStack minStack = new MinStack();
 *      minStack.push(-2);
 *      minStack.push(0);
 *      minStack.push(-3);
 *      minStack.getMin();   --> 返回 -3.
 *      minStack.pop();
 *      minStack.top();      --> 返回 0.
 *      minStack.getMin();   --> 返回 -2.
 */
class MinStack
{
    /**
     * 思路
     *     在入栈时候推入两个元素，一个是当前入栈的值，一个是当入栈时的最小值
     *     这样增加了空间1，但是却可以在常数阶的复杂度下获取最小值
     * 耗时
     *     执行用时 : 312 ms, 在Min Stack的PHP提交中击败了50.00% 的用户
     *     内存消耗 : 21.7 MB, 在Min Stack的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O（n）
     */
    public static $stack = [];
    public static $min;
    function __construct() {
        self::$stack = [];
    }

    /**
     * @param Integer $x
     * @return NULL
     */
    function push($x)
    {
        if (!isset(self::$min)) {
            self::$min = $x;
        }
        elseif(self::$min > $x) {
            self::$min = $x;
        }

        self::$stack[] = $x;
        self::$stack[] = self::$min;
    }

    /**
     * @return NULL
     */
    function pop()
    {
        array_pop(self::$stack);
        return array_pop(self::$stack);
    }

    /**
     * @return Integer
     */
    function top()
    {
        return self::$stack[count(self::$stack)-2];
    }

    /**
     * @return Integer
     */
    function getMin()
    {
        return end(self::$stack);
    }
}
/**
 * Your MinStack object will be instantiated and called as such:
 * $obj = MinStack();
 * $obj->push($x);
 * $obj->pop();
 * $ret_3 = $obj->top();
 * $ret_4 = $obj->getMin();
 */