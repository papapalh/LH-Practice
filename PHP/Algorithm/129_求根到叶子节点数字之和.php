<?php

/**
 * 题目
 *     给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
 *     例如，从根到叶子节点路径 1->2->3 代表数字 123。
 *     计算从根到叶子节点生成的所有数字之和。
 * 说明: 叶子节点是指没有子节点的节点。
 *     示例 1:
 *     输入: [1,2,3]
 *         1
 *        / \
 *       2   3
 *     输出: 25
 *     解释:
 *         从根到叶子节点路径 1->2 代表数字 12.
 *         从根到叶子节点路径 1->3 代表数字 13.
 *     因此，数字总和 = 12 + 13 = 25.
 * 示例 2:
 *     输入: [4,9,0,5,1]
 *         4
 *        / \
 *       9   0
 *      / \
 *     5   1
 *     输出: 1026
 *     解释:
 *         从根到叶子节点路径 4->9->5 代表数字 495.
 *         从根到叶子节点路径 4->9->1 代表数字 491.
 *         从根到叶子节点路径 4->0 代表数字 40.
 *     因此，数字总和 = 495 + 491 + 40 = 1026.
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
 *     树的全路径
 * 耗时
 *     执行用时 :12 ms, 在所有 PHP 提交中击败了50.00%的用户
 * 内存消耗 :15.1 MB, 在所有 PHP 提交中击败了33.33%的用户
 */
class Solution
{
    public $sum = 0;

    function sumNumbers($root)
    {
        if (!$root) {
            return 0;
        }

        $this->x($root, 0);

        return $this->sum;
    }

    function x($root, $path)
    {
        if (!$root->left && !$root->right) {
            $this->sum += $path * 10 + $root->val ;
            return true;
        }

        $path = $path * 10 + $root->val;

        if ($root->left) {
            $this->x($root->left, $path);
        }

        if ($root->right) {
            $this->x($root->right, $path);
        }
    }
}

$tree = new TreeNode(1);
$tree->left = new TreeNode(2);
$tree->left->left = new TreeNode(4);
//$tree->left->right = new TreeNode(3);
$tree->right = new TreeNode(3);
//$tree->right->left = new TreeNode(6);
//$tree->right->right = new TreeNode(9);

$s = new Solution();

var_dump($s->sumNumbers($tree));