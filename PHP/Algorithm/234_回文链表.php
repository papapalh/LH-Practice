<?php


class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}
class Solution {

    /**
     * @param ListNode $head
     * @return Boolean
     */
    function isPalindrome($head) {

        $new = [];
        $old = [];

        while ($head) {
            $old[] = $head->val;
            array_unshift($new, $head->val);
            $head = $head->next;
        }

        if ($old == $new) {
            return true;
        }

        return false;
    }
}

$list = new ListNode(1);
$list->next = new ListNode(2);
$list->next->next = new ListNode(1);

$s = new Solution();
var_dump($s->isPalindrome($list));