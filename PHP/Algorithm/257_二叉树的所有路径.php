<?php
/**
 * 题目
 *     给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *     说明: 叶子节点是指没有子节点的节点。
 * 示例:
 *      输入:
 *          1
 *        /   \
 *       2     3
 *        \
 *        5
 * 输出: ["1->2->5", "1->3"]
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
 */
class Solution
{
    /**
     * 思路
     *     深度优先遍历，输出路径
     * 耗时
     *     执行用时 : 20 ms, 在Binary Tree Paths的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 15 MB, 在Binary Tree Paths的PHP提交中击败了14.29% 的用户
     */
    public static $path = [];
    function binaryTreePaths($root)
    {
        if (!$root) {
            return [];
        }

        // 因为 Leet Code 一次脚本跑用例，这个静态变量不会被销毁。这里主动销毁
        self::$path = [];

        $this->path($root, '');

        return self::$path;
    }

    function path($root, $path)
    {
        if (!$root->left && !$root->right) {
            self::$path[] = $path.$root->val;
        }

        $path .= $root->val.'->';

        if ($root->left) {
            $this->path($root->left, $path);
        }
        if ($root->right) {
            $this->path($root->right, $path);
        }
    }
}