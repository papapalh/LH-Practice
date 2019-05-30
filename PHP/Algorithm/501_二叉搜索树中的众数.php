<?php
/**
 * 题目
 *     给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 *     假定 BST 有如下定义：
 *         结点左子树中所含结点的值小于等于当前结点的值
 *         结点右子树中所含结点的值大于等于当前结点的值
 *         左子树和右子树都是二叉搜索树
 * 例如：
 *     给定 BST [1,null,2,2],
 *         1
 *          \
 *           2
 *          /
 *         2
 *     返回[2].
 * 提示：如果众数超过1个，不需考虑输出顺序
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
 */
class Solution
{
    /**
     * 思路
     *     遍历整个树，获取节点值，计入哈希表
     *     垃圾答案
     * 耗时
     *     执行用时 : 40 ms, 在Find Mode in Binary Search Tree的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 18.4 MB, 在Find Mode in Binary Search Tree的PHP提交中击败了100.00% 的用户
     */
    public $dict = [];
    function findMode($root) {
        $arrOutput = [];
        $this->x($root);

        $max = max($this->dict);
        foreach ($this->dict as $root => $count) {
            if ($count == $max) {
                $arrOutput[] = $root;
            }
        }

        return $arrOutput;
    }

    function x($root)
    {
        if (!$root) {
            return false;
        }

        $this->dict[$root->val]++;

        $this->x($root->left);
        $this->x($root->right);
    }
}