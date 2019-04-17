<?php

/**
 * 题目
 *     给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *     如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *     您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * 示例：
 *     输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 *     输出：7 -> 0 -> 8
 *     原因：342 + 465 = 807
 */
class Node {
    public $val; //存储数据
    public $next; //存储下一个结点的指针
}

// node1
$node1 = new Node();
$node1->val = 5;
$node = $node1;

$next = new Node();
$next->val = 4;
$node->next = $next;
$node = $next;

$next = new Node();
$next->val = 3;
$node->next = $next;
$node = $next;

// node2
$node2 = new Node();
$node2->val = 2;
$node = $node2;

$next = new Node();
$next->val = 6;
$node->next = $next;
$node = $next;

$next = new Node();
$next->val = 4;
$node->next = $next;
$node = $next;

class Solution
{
    /**
     * 思路
     *     先获取两个链表相加之和
     *     在去计算进位问题
     *     当然也可以在计算链表相加的时候计算
     * 耗时
     *     执行用时 : 120 ms, 在Add Two Numbers的PHP提交中击败了40.82% 的用户
     *     内存消耗 : 45.2 MB, 在Add Two Numbers的PHP提交中击败了81.30% 的用户
     * 时间复杂度
     *     0（max(l1, l2) x 2）
     */
    function addTwoNumbers($l1, $l2)
    {
    	$arr = [];
    	$res = $l1;

    	// 先根据链表计算数值
    	while ($l1 || $l2) {

    		$arr[] = $l1->val + $l2->val;
    		
    		$l1 = $l1->next;
    		$l2 = $l2->next;
    	}

    	$p = $res;
    	$bit = 0;
    	foreach ($arr as $a) {

    		if ($a >= 10) {

    			$p->val = $a - 10;

    			if ($p->next) {
    				$bit = 1;
    			}
    			else {
    				$node = new Node();
    				$node->val = 1;
    				$p->next = $node;
    			}
    		}
    		else {
    			$p->val = $a + $bit;
    			$bit = 0;
    		}

    		$p = $p->next;
    	}

        return $res;
    }
}

$s = new Solution();
var_dump($s->addTwoNumbers($node1, $node2));