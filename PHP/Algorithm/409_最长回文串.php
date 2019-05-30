<?php
/**
 * 题目
 *     给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
 *     在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
 * 注意:
 *     假设字符串的长度不会超过 1010。
 * 示例 1:
 *     输入: "abccccdd"
 *     输出: 7
 *     解释:
 *         我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 */
class Solution
{
    /**
     * 思路
     *     哈希表
     *     先计算字符在字符串中出现次数，因为出现偶数次的字符肯定是回复串的组成部分，计算它的长度
     *     之后再计算 奇数次-1 可以构成回文的次数
     * 耗时
     *     执行用时 : 16 ms, 在Longest Palindrome的PHP提交中击败了100.00% 的用户
     *     内存消耗 : 14.7 MB, 在Longest Palindrome的PHP提交中击败了50.00% 的用户
     * 时间复杂
     *     O（n x 3）
     */
    function longestPalindrome($s)
    {
        $len       = strlen($s);
        $dict      = [];
        $twoSum    = 0; // 偶数和
        $count     = [];
        $sum       = 0;
        $oneExists = false;

        for ($i = 0; $i < $len; $i++) {
            $dict[$s[$i]]++;
        }

        foreach ($dict as $str => $nums) {

            if (($nums & 1) == 1) {

                if ($nums == 1 && !$oneExists) {
                    $oneExists = true;
                    $sum = 1;
                }
                $count[] = $nums;
                continue;
            }

            $twoSum += $nums;
        }

        // 找到中位点
        if ($sum == 0) {
            $sum = $count[0];
            unset($count[0]);
        }
        $sum += $twoSum;

        // 计算奇数列 -1值
        foreach ($count as $k => $c) {
            $sum += ($c-1);
        }

        return $sum;
    }
}

$s = new Solution();
var_dump($s->longestPalindrome("ababababad"));