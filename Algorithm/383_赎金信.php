<?php

/**
 * 题目
 *     给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。
 *     (题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)
 * 注意：
 *     你可以假设两个字符串均只含有小写字母。
 * 示例
 *     canConstruct("a", "b") -> false
 *     canConstruct("aa", "ab") -> false
 *     canConstruct("aa", "aab") -> true
 */
class Solution
{
    /**
     * 思路
     *     哈希表
     *     记录次数，在进行匹配的时候，删除对应次数，看看是否能拼装完成
     * 耗时
     *     执行用时 : 40 ms, 在Ransom Note的PHP提交中击败了66.67% 的用户
     *     内存消耗 : 15 MB, 在Ransom Note的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O（n）
     */
    function canConstruct($ransomNote, $magazine)
    {
        $sLen  = strlen($ransomNote); // 赎金信长度
        $mLen  = strlen($magazine);   // 杂志长度
        $mDict = []; // 杂志字典

        // 创建杂志字典
        for ($i = 0; $i < $mLen; $i++) {
            $mDict[$magazine[$i]]++;
        }

        for ($i = 0; $i < $sLen; $i++) {

            if (!$mDict[$ransomNote[$i]]) {
                return false;
            }

            $mDict[$ransomNote[$i]]--;
        }

        return true;
    }
}

$s = new Solution();
var_dump($s->canConstruct('aa', 'ab'));