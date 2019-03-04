<?php
/**
 *   线性表，链式存储结构
 *   特点：
 *   1. 用一组任意的存储单元存储线性表的数据元素，可以存储在内存中未被占用的任何位置
 *   2. 顺序存储结构只需要存储数据元素信息，而链式还需要存储后继元素的存储地地址
 *   3. 链表中第一个结点的存储位置叫 头指针（必要元素），头结点是在单链表的第一个结点之前附设一个结点（不必要）
 *   4. 插入和删除，都需要遍历结点找到第$i个结点，然后才是删除和插入，时间复杂度为O(n),可以看出与顺序存储结构的
 *      线性表是没有什么优势的              
 */
class LinkList
{
    public $headNode; //头结点，不是数据的一部分，只是个虚拟的节点
   
    public function __construct()
    {
        $this->initLink();
    }
   
    public function initLink() {
        $this->headNode = new Node(); //头结点
    }
   
    /**
     * @desc
     *         得到第$i个元素 时间复杂度O(n)
     * @param int $i
     */
    public function getElem($i) {
        $j = 1;//从第一个元素开始遍历
        $node = $this->headNode->getNext();
        while( $node && $j < $i) {
             $node = $node->getNext();
             ++$j;
        }
       
        if(!$node || $j > $i) { //$i 个元素不存在时
            return false;
        }
       
        $data = $node->getData(); //取第$i个元素的数据
        return $data;
    }
   
    /**
     * $desc
     *         在$i个位置之前插入数据$data         时间复杂度 O(n)
     *         在此插入多个结点时（$data多个时），    时间复杂度 第一次为O(n),后面为O(1)
     * @param int $i
     * @param $data
     */
    public function insert($i, $data) {
        $j = 1; //从第一个元素开始遍历
        $node = $this->headNode; // 指向头结点
       
        while($node && $j < $i) {
            $node = $node->getNext();   
            ++ $j;       
        }
       
        if(!$node || $j > $i) { // $i个结点不存在时
            return false;
        }
       
        $newNode = new Node();
        $newNode->setData($data);
        $newNode->setNext($node->getNext());
        $node->setNext($newNode);
        return true;
    }
    
   /**
    * @desc  
    *         删除第$i 个数据结点
    *         从$headNode之后的第一个结点才是第1个结点，以此类推 
    *        时间复杂度O(n)
    *
    *        如果在$i位置开始，删除连续的,第一次为O(n),后续为O(1)
    *        备注：如果仅仅是该方法的实现，是没有体现出单链表有啥子好处，因为调用10次delete()都是O(n),除非在该方法里进行改造。
    *             但是，如果改造该方法，我们也可以改造顺序结构的delete()方法，也可以实现连接删除，而不用每次都遍历
    * @param int $i
    */
    public function delete($i) {
        $j = 1;
        $node = $this->headNode; // 指向头结点
        while($node && $j<$i) {
            $node = $node->getNext();
            ++ $j;
        }
        if(!$node || $j>$i) {
            return false;
        }        
        $delNode = $node->getNext();         // 需要删除的结点
        $node->setNext($delNode->getNext());// 删除结点的后继结点赋值给结点
        unset($delNode);                    // 清除不需要的结点
        return true;
    }
}

/**
 * 结点
 */
class Node {
    private $data; //存储数据
    private $next; //存储下一个结点的指针
   
    public function getData() {
        return $this->data;
    }

    public function getNext() {
        return $this->next;
    }

    public function setData($data) {
        $this->data = $data;
    }

    public function setNext($next) {
        $this->next = $next;
    }
}

$list = new LinkList();
$list->insert(1, 1);
$list->insert(2, 2);
$list->insert(3, 3);
$list->insert(4, 4);
print_r($list);
echo '<hr/>';
$list->delete(3);
print_r($list);