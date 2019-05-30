<?php

/**
 * 题目
 *    给定一个二叉树，找出其最小深度。
 *    最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *    说明: 叶子节点是指没有子节点的节点。
 * 示例:
 *    给定二叉树 [3,9,20,null,null,15,7],
 *        3
 *       / \
 *      9  20
 *        /  \
 *       15   7
 *    返回它的最小深度  2.
 */

class Solution
{
    /**
     * 思路1
     *     层级递归，获取每个节点的深度，如果没有左右节点，则为叶子节点，放入字典表
     *     最后取得字典表最小值
     */
    public static $depth = 0;
    public static $min = [];
    function minDepth($root)
    {
        if (!$root) {
            return 0;
        }
        $this->min($root);

        return min(self::$min);;
    }
    function min($root)
    {
        self::$depth++;
        if(!$root) {
            goto out;
        }
        else{
            if (!$root->left && !$root->right) {
                self::$min[] = self::$depth;
            }
            $this->minDepth($root->left);
            $this->minDepth($root->right);
        }
        out:
        self::$depth--;
    }

    /**
     * 思路2
     *     大神递归
     * 耗时
     *     执行用时 : 88 ms, 在Minimum Depth of Binary Tree的PHP提交中击败了33.33% 的用户
     *     内存消耗 : 16.6 MB, 在Minimum Depth of Binary Tree的PHP提交中击败了100.00% 的用户
     */
    function minDepth1($root)
    {
        if ($root == null) {
            return 0;
        }

        if ($root->left == null) {
            return $this->minDepth($root->right) + 1;
        }

        if ($root->right == null) {
            return $this->minDepth($root->left) + 1;
        }

        return min($this->minDepth($root->left), $this->minDepth($root->right)) + 1;
    }
}