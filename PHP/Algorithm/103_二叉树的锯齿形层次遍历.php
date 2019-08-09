<?php
/**
 * 题目
 *     给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）
 * 例如:
 *     给定二叉树: [3,9,20,null,null,15,7],
 *         3
 *        / \
 *       9  20
 *         /  \
 *        15   7
 * 返回其层次遍历结果：
 * [
 *     [3],
 *     [20,9],
 *     [15,7]
 * ]
 */

class TreeNode
{
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value)
    {
        $this->val = $value;
    }
}
/**
 * 思路
 *     迭代树遍历，同时控制方向
 * 耗时
 *     执行用时 :8 ms, 在所有 PHP 提交中击败了93.75%的用户
 *     内存消耗 :15.3 MB, 在所有 PHP 提交中击败了14.29%的用户
 */
class Solution
{
    function zigzagLevelOrder($root)
    {
        if (!$root) {
            return [];
        }

        $direction            = 1; // 方向默认为正向
        $depth = 1;
        $nodeDict[$depth][]   = $root;
        $outPutDict[$depth][] = $root->val;
        $direction            = $direction ^ 1; // 初始结点之后即为反向

        while ($nodeDict[$depth]) {

            foreach ($nodeDict[$depth] as $node) {
                if ($node->left) {
                    $nodeDict[$depth+1][]   = $node->left;
                    $outPutDict[$depth+1][] = $node->left->val;

                }
                if ($node->right) {
                    $nodeDict[$depth+1][]   = $node->right;
                    $outPutDict[$depth+1][] = $node->right->val;

                }
            }

            if ($outPutDict[$depth+1] && $direction ^ 1) {
                $outPutDict[$depth+1] = array_reverse($outPutDict[$depth+1]);
            }

            $direction = $direction ^ 1;

            $depth++;
        }

        return $outPutDict;
    }
}

$tree = new TreeNode(3);
$tree->left = new TreeNode(9);
$tree->right = new TreeNode(20);
$tree->right->left = new TreeNode(15);
$tree->right->right = new TreeNode(7);

$s = new Solution();
var_dump($s->zigzagLevelOrder($tree));