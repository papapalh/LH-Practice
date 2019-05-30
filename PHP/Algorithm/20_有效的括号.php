<?php
/**
 * 题目
 *     给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *     有效字符串需满足：
 *     左括号必须用相同类型的右括号闭合。
 *     左括号必须以正确的顺序闭合。
 *     注意空字符串可被认为是有效字符串。
 * 用例
 *     输入: "()"
 *     输出: true
 *     输入: "()[]{}"
 *     输出: true
 *     输入: "(]"
 *     输出: false
 *     输入: "{[]}"
 *     输出: true
 */
class Solution
{
    /**
     * 思路
     *     因为左右分开的符号是一定成对出现的。所以这里使用一个 栈 。保存符号
     *     如果发现是 关闭 符号，则 出栈 获取。
     *     如果不同，则退出算法
     * 详解
     *     1：输入 '[{}]}'
     *     2：因为 '[{' 不是关闭符，所以入栈，当前栈结构 ['[', '{']
     *     3：当 '}' 进入，发现是关闭符，所以出栈获取字典 '{',相同，所以是符合规则
     * 耗时
     *     执行用时 : 56 ms, 在Valid Parentheses的PHP提交中击败了26.42% 的用户
     *     内存消耗 : 14.9 MB, 在Valid Parentheses的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O(n)
     */
    function isValid($s)
    {
        $length = strlen($s);  // 循环次数
        $stack  = [];          // 栈
        $hash   = [            // hash表
            ')' => '(',
            ']' => '[',
            '}' => '{',
        ];

        for ($i = 0; $i < $length; $i++) {
            if ($hash[$s[$i]]) {
                if ($hash[$s[$i]] != array_pop($stack)) {
                    return false;
                }
                continue;
            }
            $stack[] = $s[$i];
        }

        if ($stack) {
            return false;
        }

        return true;
    }
}

$s = new Solution();
var_dump($s->isValid('[{}]'));