<?php

/**
 * 题目
 *     给定两个字符串 s 和 t，它们只包含小写字母。
 *     字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
 *     请找出在 t 中被添加的字母。
 * 示例:
 *     输入：
 *     s = "abcd"
 *     t = "abcde"
 * 输出：
 *     e
 * 解释：
 *     'e' 是那个被添加的字母。
 */
class Solution
{
    /**
     * 思路
     *    哈希表
     *    两次循环，第一次创建哈希表，第二次查找
     *    这都能超过100%，应该是大家都觉得这个太简单就不做了吧......
     *    这两天实在是太忙了，就做个这个简单的吧
     * 耗时
     *     执行用时 : 24 ms, 在Find the Difference的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 15 MB, 在Find the Difference的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    function findTheDifference($s, $t)
    {
        $sLen = strlen($s);
        $tLen = strlen($t);

        $dict = [];

        for ($i = 0; $i < $sLen; $i++) {
            $dict[$s[$i]] += 1;
        }

        for ($i = 0; $i < $tLen; $i++) {
            if (!$dict[$t[$i]]) {
                return $t[$i];
            }

            $dict[$t[$i]]--;
        }
    }
}

$s = new Solution();
var_dump($s->findTheDifference('a', 'aa'));