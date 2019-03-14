<?php
/**
 * 题目
 * 		实现 strStr() 函数。
 * 		给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
 * 示例
 *    输入: haystack = "hello", needle = "ll"
 *    输出: 2
 *    输入: haystack = "aaaaa", needle = "bba"
 *    输出: -1
 */
class Solution
{
    /**
     * 思路1
     *    暴力破解，这个没啥说的,挨个循环吧
     * 耗时
     *    执行用时 : 56 ms, 在Implement strStr()的PHP提交中击败了26.32% 的用户
     *    内存消耗 : 16.5 MB, 在Implement strStr()的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *    O(n)
     */
    function strStr($haystack, $needle)
    {
        if (!$needle) {
            return 0;
        }
        $length = strlen($needle);

        $i = 0; // 短指针，指向匹配字符串头
        
        while ($str = substr($haystack, $i, $length)) {

            if ($str == $needle) {
                return $i;
            }
            $i++;
        }

        return -1;
    }

    /**
     * 思路2
     *    KMP 算法
     *    利用已经部分匹配这个有效信息，保持i指针不回溯，通过修改j指针，让模式串尽量地移动到有效的位置。
     * 具体算法示例
     *    输入 haystack = "hello world", needle = "orld"
     *    	1:循环 haystack 字符串,既然要找匹配 'or' 的字符串,那么开头肯定是为 'o' 的。
     *         根据这个，我们找到了下标为 5 的以 'o' 开头的。
     *      2:
     * 耗时
     *    执行用时 : 28 ms, 在Implement strStr()的PHP提交中击败了42.11% 的用户
     *    内存消耗 : 16.5 MB, 在Implement strStr()的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *    O(n)
     * 详解
     *     https://www.cnblogs.com/yjiyjige/p/3263858.html
     */
    function strStr($haystack, $needle)
    {
        if (!$needle) {
            return 0;
        }

        $target_length = strlen($needle);
        $length = strlen($haystack);

        $i = 0;
        while ($i <= $length) {

            if ($length - $i < $target_length) {
        		return -1;
        	}
            $str = '';
            
            if ($haystack[$i] == $needle[0]) {
            	$kmp = 0;
                for ($j = 0; $j < $target_length; $j++) {
                    $str .= $haystack[$i + $j];
                    if (($haystack[$i + $j] == $needle[0]) && !$kmp) {
                        $kmp = $i + $j + 1;
                    }
                }
                
                if ($str == $needle) {
                    return $i;
                } else {
                    $i = $kmp ? $kmp : $i;
                }
            }
            else {
                $i++;
            }
        }

        return -1;
    }
}


$s = ["h","e","l","l","o"];
$Object = new Solution();
$Object->reverseString($s);
var_dump($s);
