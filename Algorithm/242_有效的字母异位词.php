<?php
/**
 * 题目
 *     给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
 * 用例
 *     输入: s = "anagram", t = "nagaram"
 *     输出: true
 *     输入: s = "rat", t = "car"
 *     输出: false
 */
class Solution
{
    /**
     * 什么是异位词
     *     两个单词如果包含相同的字母，次序不同，则称为字母易位词(anagram)
     * 异位词特点
     *     字符数相同/顺序不同/出现字符相同
     * 思路
     *     既然是找寻字符串的不同，那我们可以直接建立对应的 哈希表，来对两个 hash 表 进行操作
     *     在比对的时候可以达到O(1)的复杂度
     * 耗时
     *     执行用时 : 20 ms, 在Valid Anagram的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 15 MB, 在Valid Anagram的PHP提交中击败了100.00% 的用户
     * 时间复杂度 
     *     O(n)
     */
    function isAnagram($s, $t)
    {
        $hashS = [];
        $hashT = [];
        $lengthS = strlen($s);
        $lengthT = strlen($t);

        // 如果字符数不相同，必然不是
        if ($lengthS != $lengthT) {
            return false;
        }

        // S，T 字符串 Hash 表
        for ($i = 0; $i < $lengthT; $i++) { 
            $hashT[$t[$i]] += 1;
            $hashS[$s[$i]] += 1;
        }

        // 是否存在和出现次数比对
        foreach ($hashT as $key => $value) {
            if ($hashS[$key] != $value) {
                return false;
            }
        }

        return true;
    }
}

$s = new Solution();
var_dump($s->isAnagram('anagram', 'nagaram'));
