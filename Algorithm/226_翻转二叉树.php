<?php
/**
 * 题目
 *     翻转一棵二叉树。
 * 示例：
 *     输入：
 *             4
 *           /   \
 *          2     7
 *         / \   / \
 *        1   3 6   9
 *     输出：
 *             4
 *           /   \
 *          7     2
 *         / \   / \
 *        9   6 3   1
 * 备注:
 *     这个问题是受到 Max Howell 的 原问题 启发的 ：
 *     谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
 */
class Solution
{
    /**
     * 思路
     *     这个小故事真的好玩，哈哈
     *     深度优先递归
     *     每个结点的左右节点互换，完成反转
     * 耗时
     *     执行用时 : 12 ms, 在Invert Binary Tree的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.7 MB, 在Invert Binary Tree的PHP提交中击败了71.43% 的用户
     */
    function invertTree($root)
    {
        // 返回树结构，处理引用
        $tree = $root;

        $this->x($root);

        return $tree;
    }

    function x($root)
    {
        if (!$root) {
            return false;
        }

        // 左右结点，交换顺序
        list($root->right, $root->left) = [$root->left, $root->right];

        $this->x($root->left);
        $this->x($root->right);
    }
}

class TreeNode
{
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}
$tree = new TreeNode(4);
$tree->left = new TreeNode(2);
$tree->left->left = new TreeNode(1);
$tree->left->right = new TreeNode(3);
$tree->right = new TreeNode(7);
$tree->right->left = new TreeNode(6);
$tree->right->right = new TreeNode(9);

$s = new Solution();
var_dump($s->invertTree($tree));