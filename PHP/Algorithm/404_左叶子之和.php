<?php
/**
 * 题目
 *     计算给定二叉树的所有左叶子之和。
 * 示例：
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 */
class Solution
{
    /**
     * 思路
     *     深度优先遍历，在多一层判断是否是单个左叶子节点
     *     如何使，则计算和
     * 耗时
     *     执行用时 : 20 ms, 在Sum of Left Leaves的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 15.4 MB, 在Sum of Left Leaves的PHP提交中击败了100.00% 的用户
     */
    public static $sum = 0;
    function sumOfLeftLeaves($root) {
        self::$sum = 0; // 初始化
        $this->recursion($root);
        return self::$sum;
    }

    function recursion($root) {
        if (!$root) {
            return;
        }

        if ($root->left && !$root->left->left && !$root->left->right) {
            self::$sum += $root->left->val;
        }

        $this->recursion($root->left);
        $this->recursion($root->right);
    }
}

class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}
$tree = new TreeNode(0);
$tree->left = new TreeNode(2);
$tree->left->left = new TreeNode(1);
$tree->left->left->left = new TreeNode(5);
$tree->left->left->right = new TreeNode(1);
$tree->left->right = null;

$tree->right = new TreeNode(4);
$tree->right->left = new TreeNode(3);
$tree->right->left->left = null;
$tree->right->left->right = new TreeNode(6);
$tree->right->right = new TreeNode(-1);
$tree->right->right->left = null;
$tree->right->right = new TreeNode(8);

$s = new Solution();
var_dump($s->sumOfLeftLeaves($tree));
