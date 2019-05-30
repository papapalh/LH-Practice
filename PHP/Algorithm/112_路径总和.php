<?php

/**
 * 题目
 *      给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 *      说明: 叶子节点是指没有子节点的节点。
 * 示例: 
 *      给定如下二叉树，以及目标和 sum = 22，
 *
 *             5
 *            / \
 *           4   8
 *          /   / \
 *         11  13  4
 *        /  \      \
 *       7    2      1
 *      返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
 */
class Solution
{
    /**
     * 思路
     *     深度优先-前序遍历记录上个结点的和
     *     之后向下遍历
     *     等于要获取二叉树全部的路径
     * 耗时
     *     执行用时 : 48 ms, 在Path Sum的PHP提交中击败了33.33% 的用户
     *     内存消耗 : 16.6 MB, 在Path Sum的PHP提交中击败了100.00% 的用户
     * 时间复杂度 
     *     O(n)
     */
    function hasPathSum($root, $sum)
    {
        if (!$root) {
            return false;
        }

        $this->path($root, 0);

        foreach (self::$path as $path) {
            if ($path == $sum) {
                return true;
            }
        }

        return false;
    }

    public static $path = [];
    function path($root, $num)
    {
        $num += $root->val;
        if (!$root->right && !$root->left) {
            self::$path[] = $num; 
        }

        if ($root->left) {
            $this->path($root->left,$num);
        }
        if ($root->right) {
            $this->path($root->right,$num);
        }
    }
}