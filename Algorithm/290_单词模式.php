<?php
/**
 * 题目
 *     给定一种 pattern(模式) 和一个字符串 str ，判断 str 是否遵循相同的模式。
 *     这里的遵循指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应模式。
 * 示例1:
 *     输入: pattern = "abba", str = "dog cat cat dog"
 *     输出: true
 * 示例 2:
 *     输入:pattern = "abba", str = "dog cat cat fish"
 *     输出: false
 * 示例 3:
 *     输入: pattern = "aaaa", str = "dog cat cat dog"
 *     输出: false
 * 示例 4:
 *     输入: pattern = "abba", str = "dog dog dog dog"
 *     输出: false
 * 说明:
 *     你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。    
 */
class Solution
{
    /**
     * 思路
     *     双哈希表
     *     用于检查两个位置的存在情况
     * 耗时
     *     执行用时 : 20 ms, 在Word Pattern的PHP提交中击败了64.71% 的用户
     *     内存消耗 : 14.7 MB, 在Word Pattern的PHP提交中击败了33.33% 的用户
     * 时间复杂度
     *     O（n）
     */
    function wordPattern($pattern, $str)
    {
        $arr   = explode(' ', $str);
        $dict  = [];
        $dict2 = [];

        $pLen = strlen($pattern);
        $sLen = count($arr);

        // 长度错误，直接返回错误
        if ($pLen != $sLen) {
            return false;
        }

        // 特殊处理 长度为2的
        if ($pLen == 2 && ($pattern[0] != $pattern[1]) && $arr[0] == $arr[1]) {
            return false;
        }
    
        for ($i = 0; $i < $pLen; $i++) {

            // 双哈希表映射
            if (!$dict[$pattern[$i]]) {
                $dict[$pattern[$i]] = $arr[$i];
                $dict2[$arr[$i]] = $pattern[$i];
                continue; 
            }

            if ($dict2[$arr[$i]] != $pattern[$i]) {
                return false;
            }

            if ($dict[$pattern[$i]] != $arr[$i]) {
                return false;
            }
        }

        return true;
    }
}