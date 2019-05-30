<?php
/**
 * 题目
 *     将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
 *     本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 * 示例:
 *     给定有序数组: [-10,-3,0,5,9],
 *     一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
 *         0
 *        / \
 *      -3   9
 *      /   /
 *    -10  5
 */
class Solution
{
    /**
     * 思路
     *     有序数列构建二叉搜索数
     *     取数组的中间元素作为根结点，将数组分成左右两部分，对数组的两部分用递归的方法分别构建左右子树。
     * 耗时
     *     执行用时 : 36 ms, 在Convert Sorted Array to Binary Search Tree的PHP提交中击败了86.36% 的用户
     *     内存消耗 : 17.9 MB, 在Convert Sorted Array to Binary Search Tree的PHP提交中击败了100.00% 的用户
     */
    function sortedArrayToBST($nums) {
        if (!$nums) {
            return;
        }
        $mid = floor(count($nums) / 2);

        $root = new TreeNode($nums[$mid]);

        $root->left = $this->sortedArrayToBST(array_slice($nums, 0, $mid));
        $root->right = $this->sortedArrayToBST(array_slice($nums, $mid+1));

        return $root;
    }
}