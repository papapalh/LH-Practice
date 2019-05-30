<?php

class Node {
    public $val; //存储数据
    public $next; //存储下一个结点的指针
}

$head = new Node();
$head->val = 1;
$node = $head;

$next = new Node();
$next->val = 2;
$node->next = $next;
$node = $next;

$next = new Node();
$next->val = 3;
$node->next = $next;
$node = $next;


$next = new Node();
$next->val = 4;
$node->next = $next;
$node = $next;

$next = new Node();
$next->val = 5;
$node->next = $next;
$node = $next;
