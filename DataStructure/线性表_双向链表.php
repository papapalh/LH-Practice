<?php

/**
 *   线性表，双向链表
 *   特点：
 *      1:所以在双向链表结点都有两个指针域，一个指向前驱位置，一个指向后驱位置。
 *      2:双向链表也叫双面链表。
 *      3:每个结点有两个链接：
 *          一个指向前一个结点，当此节点为第一个节点时，指向空值；
 *          一个指向下一个节点，当此节点为最后一个节点时，指向空值。       
 */

class DoubleLinkList
{
    public $headNode; //头结点
   
    public function __construct()
    {
        $this->initLink();
    }
   
    public function initLink() {
        $this->headNode = new Node();
    }

    /**
     * 检查链表是否为空
     * 检查头结点的后驱是否为空就可以
     */
    public function isEmpty()
    {
    	if (!$this->headNode->getNext()) {
    		return FALSE;
    	}

    	return TRUE;
    }

    /**
     * 获取链表长度
     * 和单链表相同
     */
    public function getLength()
    {
    	$count = 0;
    	$node  = $this->headNode;

    	while ($node = $node->getNext()) {
    		$count++;
    	}

    	return $count;
    }

    /**
     * 指定位置添加元素
     * 和单链表相比，需要更换下不但要更换后驱指针，前驱指针也需要更换
     */
    public function insert($i, $data)
    {
    	$j = 1; //从第一个元素开始遍历
        $node = $this->headNode; // 指向头结点
       
        while($node && ($j > $i)) {
            $node = $node->getNext();   
            ++ $j;
        }
       
        if(!$node) { // $i个结点不存在时
            return false;
        }
       
        $newNode = new Node();
        $newNode->setData($data);             // 设置新结点值
        $newNode->setNext($node->getNext());  // 设置新节点后驱
        $newNode->setBefore($node);           // 设置新节点前驱

        $node->setNext($newNode);  // 设置下一个结点
        return true;
    }
}

/**
 * 节点实现
 */
class Node
{
    private $data;   //存储数据
    private $next;   //存储下一个结点的指针
    private $before; //存储上一个结点的指针

    public function getData()
    {
        return $this->data;
    }

    public function getNext()
    {
        return $this->next;
    }

    public function getBefore()
    {
        return $this->before;
    }

    public function setData($data)
    {
        $this->data = $data;
    }

    public function setNext($next)
    {
        $this->next = $next;
    }

    public function setBefore($befor)
    {
        $this->before = $befor;
    }
}

$list = new DoubleLinkList();
$list->insert(1,2);
$list->insert(1,3);
$list->insert(1,4);
$list->insert(2,3);

print_r($list);
