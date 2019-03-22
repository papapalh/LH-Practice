<?php

/**
 * 二叉树遍历
 *     注意：已知 前序/后序 遍历方式，是不能确定一颗二叉树的
 */

class Node
{
    public $left;  // 左结点
    public $right; // 右结点
    public $val;   // 值
}

class Tree
{
    /**
     * 建立二叉树
     *        1
     *       / \
     *      2  3
     *     /\  /\
     *    4 5 6 7
     */
    public function createTree()
    {
        $head = new Node();
        $head->val = 1;

        $head->left = new Node();
        $head->left->val = 2;
        $head->left->left = new Node();
        $head->left->left->val = 4;
        $head->left->right = new Node();
        $head->left->right->val = 5;


        $head->right = new Node();
        $head->right->val = 3;
        $head->right->left = new Node();
        $head->right->left->val = 6;
        $head->right->right = new Node();
        $head->right->right->val = 7;

        return $head;
    }

    /**
     * 前序遍历
     * 遍历顺序为 1-2-3-4-5-6
     */
    public function proOrder($tree)
    {
        if (!$tree->val) return false;

        echo $tree->val, "\n";

        $this->proOrder($tree->left);
        $this->proOrder($tree->right);
    }

    /**
     * 中序遍历
     * 遍历顺序为 4-2-5-1-6-3-7
     */
    public function middleOrder($tree)
    {
        if (!$tree->val) return false;

        $this->middleOrder($tree->left);

        echo $tree->val, "\n";

        $this->middleOrder($tree->right);
    }

    /**
     * 后序遍历
     * 遍历顺序为 4-5-2-6-7-3-1
     */
    public function backOrder($tree)
    {
        if (!$tree->val) return false;

        $this->backOrder($tree->left);
        $this->backOrder($tree->right);

        echo $tree->val, "\n";
    }

    /**
     * 层级遍历
     * 遍历顺序为 1-2-3-4-5-6-7
     */
    public function hierarchy($tree)
    {
        if (!$tree->val) {
            return false;
        }

        // 根节点推入队列
        $list[] = $tree;

        // 循环队列
        while ($list) {
            $tmpList = [];
            foreach ($list as $v) {
                echo $v->val, "\n";

                if ($v->left->val) {
                    $tmpList[] = $v->left;
                }

                if ($v->right->val) {
                    $tmpList[] = $v->right;
                }
            }
            $list = $tmpList;
        }
    }

    /**
     * 递归层级遍历
     * 遍历顺序为 1-2-3-4-5-6-7
     */
    public static $depth = 0;
    function hierarchy2($root)
    {
        self::$depth++;
        if(!$root) {
            goto out;
        }
        else{
            echo $root->val,'-',self::$depth,"\n";
            $this->hierarchy2($root->left);
            $this->hierarchy2($root->right);
        }
        out:
        self::$depth--;
    }
}

$t = new Tree();
$tree = $t->createTree();

//$t->proOrder($tree);

//$t->middleOrder($tree);

//$t->backOrder($tree);

//$t->hierarchy($tree);

$t->hierarchy2($tree);