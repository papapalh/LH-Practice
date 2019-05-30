<?php
/**
 * 题目
 *     给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
 * 例如，给定一个 3叉树 :
 *           1
 *         / \ \
 *        3  2  4
 *       / \
 *      5  6
 *     返回其层序遍历:
 *     [
 *         [1],
 *         [3,2,4],
 *         [5,6]
 *     ]
 * 说明:
 *     树的深度不会超过 1000。
 *     树的节点总数不会超过 5000。
 */
class Solution
{
    /**
     * 思路
     *     广度优先遍历
     *     Leet Code 这个编辑器报错真是很迷离啊~
     * 耗时
     *     执行用时 : 760 ms, 在N-ary Tree Level Order Traversal的PHP提交中击败了96.15% 的用户
     *     内存消耗 : 144.2 MB, 在N-ary Tree Level Order Traversal的PHP提交中击败了100.00% 的用户
     */
    function levelOrder($root)
    {
        $dict  = [];
        $depth = 1;
        $aOutput = [];

        if (!$root) {
            return [];
        }

        $aOutput[$depth][] = $root->val;
        $depth += 1;
        $dict[$depth] = $root->children;

        while ($dict[$depth]) {

            foreach ($dict[$depth] as $d) {

                $aOutput[$depth][] = $d->val;

                if ($dict[$depth+1]) {
                    $dict[$depth+1] = array_merge($dict[$depth+1], $d->children);
                }
                else {
                    $dict[$depth+1] = $d->children;
                }
            }

            $depth += 1;
        }

        return $aOutput;
    }
}
//class Node
//{
//    public $val;
//    public $children;
//
//    function __construct($val, $children)
//    {
//        $this->val = $val;
//        $this->children = $children;
//    }
//}
//$tree = new Node(1,[new Node(3, []), new Node(2, []), new Node(4, [])]);

$s = new Solution();
var_dump($s->levelOrder($tree));