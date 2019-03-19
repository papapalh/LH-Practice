<?php

/**
 * 题目
 *     给定两个二叉树，编写一个函数来检验它们是否相同。
 *     如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 * 示例
 *     输入:       1         1
 *               / \       / \
 *              2   3     2   3
 *             [1,2,3],   [1,2,3]
 *     输出: true
 *
 *     输入:      1          1
 *              /           \
 *             2             2
 *           [1,2],     [1,null,2]
 *     输出: false
 */
class Solution
{
    /**
     * 思路
     *     深度优先 - 前序遍历二叉树
     *     递归处理 - 对比结点
     * 耗时
     *     执行用时 : 16 ms, 在Same Tree的PHP提交中击败了80.00% 的用户
     *     内存消耗 : 14.9 MB, 在Same Tree的PHP提交中击败了100.00% 的用户
     */
    function isSameTree($p, $q)
    {
        if($p == null && $q == null) {
            return true;
        }

        if ($p != null && $q != null && $p->val == $q->val) {
            return $this->isSameTree($p->left, $q->left) && $this->isSameTree($p->right, $q->right);
        }
        return false;
    }
}