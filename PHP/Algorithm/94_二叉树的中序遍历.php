<?php
/**
 * 题目
 *     给定一个二叉树，返回它的中序 遍历。
 * 示例:
 *     输入: [1,null,2,3]
 *         1
 *          \
 *          2
 *         /
 *        3
 *     输出: [1,3,2]
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */
class Solution
{
    /**
     * 思路
     *     中序遍历
     * 耗时
     *     执行用时 :8 ms, 在所有 PHP 提交中击败了83.87%的用户
     *     内存消耗 :14.7 MB, 在所有 PHP 提交中击败了74.51%的用户
     */
    public $output;

    function inorderTraversal($root)
    {
        if (!$root) {
            return [];
        }

        $this->inorderTraversal($root->left);

        $this->output[] = $root->val;

        $this->inorderTraversal($root->right);

        return $this->output;
    }
}