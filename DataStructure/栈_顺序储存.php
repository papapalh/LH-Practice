<?php
/**
 *    栈的定义：栈是只允许在表尾进行操作的线性表
 *    栈的顺序存储结构的实现,只实现了最主要的进，出操作。
 *    顺序结构即由一段连续的空间存储，即数组实现
 */
class Stack
{
    const  MAXSIZE = 20; // 存储栈长度
    public $top;  // 栈位置
    public $data; // 栈
   
    // 入栈
    public function push($elem)
    {
        if($this->top == self::MAXSIZE - 1) {
            return false;
        }
        $this->top++;
        $this->data[$this->top] = $elem;
        return true;
    }

    // 出栈
    public function pop()
    {
        if($this->top == -1) {
            return false;
        }
        $elem = $this->data[$this->top];
        $this->top--;
        return $elem;
    }
}

$stack = new Stack();
$stack->push(1);
$stack->push(2);
$stack->push(3);
$stack->push(4);

print_r($stack->pop());
