<?php
/**
 * 题目
 *     给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 * 用例
 *     输入: s = "leetcode"
 *     返回 0.
 *     输入: s = "loveleetcode",
 *     返回 2.
 */
class Solution
{
    /**
     * 思路
     *     HASH 表方法
     *     如果为首次出现，则放入 Hash/结果集中
     *     如果再次出现，则删除结果集，保留Hash表数据
     *     这样，结果集中的第一个元素就是首次出现不重复的字符了
     * 耗时
     *     执行用时: 80 ms, 在First Unique Character in a String的PHP提交中击败了72.73% 的用户
     *     内存消耗: 15 MB, 在First Unique Character in a String的PHP提交中击败了100.00% 的用户
     * 时间复杂度 
     *     O(n)
     */
    function firstUniqChar($s)
    {
        $length  = strlen($s) - 1;
        $hash    = []; // 定义 Hash 表
        $result  = []; // 获取结果

        for ($i = 0; $i <= $length; $i++) {

            // 如果存在，则在结果中删除，在 HASH 表中保留
            if (isset($hash[$s[$i]])) {
                unset($result[$s[$i]]);
                continue;
            }

            $hash[$s[$i]] = $result[$s[$i]] = $i;
        }

        return $result ? current($result) : -1;
    }
}

$str = 'aadadaad';
$s = new Solution();
echo $s->firstUniqChar($str);