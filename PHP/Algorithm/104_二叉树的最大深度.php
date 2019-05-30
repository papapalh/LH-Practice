<?php

/**
 * 题目
 *    给定一个二叉树，找出其最大深度。
 *    二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *    说明: 叶子节点是指没有子节点的节点。
 * 示例：
 *    给定二叉树 [3,9,20,null,null,15,7]，
 *       3
 *      / \
 *    89  20
 *       /  \
 *      15   7
 *    返回它的最大深度 3 。
 */

class Solution
{
    /**
     * 思路
     *     深度优先 - 递归
     *     这道题有点不动脑子，看了大神的想法
     * 耗时
     *     执行用时 : 64 ms, 在Maximum Depth of Binary Tree的PHP提交中击败了16.67% 的用户
     *     内存消耗 : 16.9 MB, 在Maximum Depth of Binary Tree的PHP提交中击败了100.00% 的用户
     */
    function maxDepth($root)
    {
        return $root == null ? 0 : max($this->maxDepth($root->left), $this->maxDepth($root->right)) + 1;
    }
}