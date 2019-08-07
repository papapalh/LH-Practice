<?php
/**
 * 题目
 *     给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 * 例如:
 *     给定二叉树: [3,9,20,null,null,15,7],
 *         3
 *        / \
 *       9  20
 *         /  \
 *        15   7
 * 返回其层次遍历结果：
 * [
 *     [3],
 *     [9,20],
 *     [15,7]
 * ]
 */

class TreeNode
{
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value)
    {
        $this->val = $value;
    }
}
/**
 * 思路
 *     前序遍历，获取深度
 * 耗时
 *     执行用时 :12 ms, 在所有 PHP 提交中击败了63.64%的用户
 *     内存消耗 :15.6 MB, 在所有 PHP 提交中击败了29.41%的用户
 */
class Solution
{
    public $outPut = [];

    function levelOrder($root)
    {
        if (!$root) {
            return $this->outPut;
        }

        $this->x($root, 0);

        return $this->outPut;
    }

    function x($root, $depth)
    {
        if (!$root) {
            return false;
        }

        $this->outPut[$depth][] = $root->val;

        $depth++;

        $this->x($root->left, $depth);
        $this->x($root->right, $depth);
    }
}

$tree = new TreeNode(3);
$tree->left = new TreeNode(9);
$tree->right = new TreeNode(20);
$tree->right->left = new TreeNode(15);
$tree->right->right = new TreeNode(7);

$s = new Solution();
var_dump($s->levelOrder($tree));


