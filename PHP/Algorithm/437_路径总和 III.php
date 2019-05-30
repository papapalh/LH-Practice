<?php
/**
 * 题目
 *     给定一个二叉树，它的每个结点都存放着一个整数值。
 *     找出路径和等于给定数值的路径总数。
 *     路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
 *     二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
 * 示例：
 *     root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *          10
 *         /  \
 *        5   -3
 *       / \    \
 *      3   2   11
 *     / \   \
 *    3  -2   1
 *    返回 3。和等于 8 的路径有:
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3.  -3 -> 11
 */
class Solution
{
    /**
     * 思路
     *     递归回溯
     *     递归树的同时，将每个结点作为新结点计算路径长度
     *     等于双重递归这个数结构
     * 耗时
     *     执行用时 : 300 ms, 在Path Sum III的PHP提交中击败了83.33% 的用户
     *     内存消耗 : 15.5 MB, 在Path Sum III的PHP提交中击败了50.00% 的用户
     */
    public function pathSum($root, $sum)
    {
        if ($root == null) {
            return 0;
        }
        return $this->pathSum($root->left, $sum) + $this->pathSum($root->right, $sum) + $this->oneRootPathSum($root, $sum);
    }

    private function oneRootPathSum($node, $sum)
    {
        if ($node == null) {
            return 0;
        }
        $sum -= $node->val;
        return ($sum == 0 ? 1 : 0) + $this->oneRootPathSum($node->left, $sum) + $this->oneRootPathSum($node->right, $sum);
    }
}