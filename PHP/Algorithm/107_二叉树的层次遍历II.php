<?php

/**
 * 题目
 *    给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 * 示例
 *    给定二叉树 [3,9,20,null,null,15,7],
 *        3
 *       / \
 *      9  20
 *        /  \
 *       15   7
 *    返回其自底向上的层次遍历为：
 *    [
 *        [15,7],
 *        [9,20],
 *        [3]
 *    ]
 */

class Solution
{
    /**
     * 思路
     *    广度优先 - 层级遍历
     *    还是层级遍历，从上至下，不过是在输出的时候颠倒下顺序，属于取巧做法，并不符合题意
     * 耗时
     *     执行用时 : 36 ms, 在Binary Tree Level Order Traversal II的PHP提交中击败了0.00% 的用户
     *     内存消耗 : 15.5 MB, 在Binary Tree Level Order Traversal II的PHP提交中击败了100.00% 的用户
     */
    function levelOrderBottom($root)
    {
        $tree = [];

        if (!$root) {
            return $tree;
        }

        // 推入根节点
        $tree[] = [$root->val];

        // 层级循环字典
        $dict = [$root];

        while ($dict) {
            $depth = []; // 层节点
            $t     = []; // 层节点值
            foreach ($dict as $node) {
                if ($node->left) {
                    $depth[] = $node->left;
                    $t[] = $node->left->val;
                }
                if ($node->right) {
                    $depth[] = $node->right;
                    $t[] = $node->right->val;
                }
            }

            if ($t) {
                array_unshift($tree, $t);
            }

            $dict = $depth;
        }

        return $tree;
    }
}
