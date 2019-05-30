<?php
/**  
 * 
 * 栈的 顺序结构 与 链结构，他们在时间复杂度上都一样，都为O(1).
 * 如果栈的数量可预知，则使用顺序栈，否则，则使用链栈
 * 链栈 要求每个元素都有指针域，增加了内存开销，但对于长度无限制。
 */
class LinkStack
{
    public $top;   // 栈顶位置
    public $count; // 栈长度
        
    public function push($elem)
    {
        $node       = new StackNode();
        $node->data = $elem;      // 数据
        $node->next = $this->top; // 下一个结点指针
        $this->top  = $node;      // 当前栈顶
        $this->count++;           // 栈长度变化
        return true;
    }

    public function pop()
    {
        if($this->count == 0) return false; // 空栈没有数据，无法出栈
        $elem = $this->top->data; // 获取栈顶数据，出栈
        $top  = $this->top;
        $this->top = $this->top->next;
        $this->count--;
        unset($top);
        return $elem;
    }
}

class StackNode {
    public $data;
    public $next;
}

$stack = new LinkStack();
$stack->push(1);
$stack->push(2);
// $stack->push(3);
// $stack->push(4);
// $stack->push(5);

print_r($stack->pop());
print_r($stack->pop());
print_r($stack->pop());
// var_dump($stack);
// var_dump($stack->top->next);
