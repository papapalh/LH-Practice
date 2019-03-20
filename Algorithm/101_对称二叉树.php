<?php

/**
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *         1
 *        / \
 *       2   2
 *      / \ / \
 *     3  4 4  3
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *         1
 *        / \
 *       2   2
 *       \   \
 *       3    3
 *  说明:
 *  如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
 */
class Solution
{
    public static $list = [];

    /**
     * 思路
     *    层级遍历二叉树
     *    通过左右层级来对比是否为对称
     * 详解
     *         1
     *        / \
     *       2   2
     *      / \ / \
     *     3  4 4  3
     *     从二层开始遍历
     *         得到数组 [2] [2] 对比是否对称
     *     三层遍历
     *         得到数组 [3，4] [4，3] 对比是否对称
     * 耗时
     *    执行用时 : 24 ms, 在Symmetric Tree的PHP提交中击败了58.82% 的用户
     *    内存消耗 : 15 MB, 在Symmetric Tree的PHP提交中击败了100.00% 的用户
     */
    function isSymmetric($root)
    {
        if (!$root || (!$root->left && !$root->right)) {
            return true;
        }

        // 根节点推入队列
        self::$list[] = $root;

        while (self::$list) {
            $list = [];
            $count = 0;
            foreach (self::$list as $v) {

                if ($v->left) {
                    $list[1][] = $v->left;
                }

                if ($v->right) {
                    $list[2][] = $v->right;
                }
                $count++;
            }

            for ($i = 0;$i < $count; $i++) {

                if ($list[1][$i]->val != ($list[2][$count-$i]->val)) {
                    return false;
                }
            }

            self::$list = array_merge($list[1], $list[2]);
        }

        return true;
    }
}