<?php
/**
 * 题目
 *     写一个程序，输出从 1 到 n 数字的字符串表示。
 *     1. 如果 n 是3的倍数，输出“Fizz”；
 *     2. 如果 n 是5的倍数，输出“Buzz”；
 *     3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
 * 示例：
 *     n = 15,
 * 返回:
 *     [
 *         "1",
 *         "2",
 *         "Fizz",
 *         "4",
 *         "Buzz",
 *         "Fizz",
 *         "7",
 *         "8",
 *         "Fizz",
 *         "Buzz",
 *         "11",
 *         "Fizz",
 *         "13",
 *         "14",
 *         "FizzBuzz"
 *     ]
 */
class Solution
{
    /**
     * 思路
     *     没啥说的......
     * 耗时
     *     执行用时 : 24 ms, 在Fizz Buzz的PHP提交中击败了88.00% 的用户
     *     内存消耗 : 18.2 MB, 在Fizz Buzz的PHP提交中击败了42.86% 的用户
     * 时间复杂度
     *     O（n）
     */
    function fizzBuzz($n)
    {
        $arrOutput = [];

        $i = 1;

        while ($n >= $i) {

            if (($i % 3) == 0 && ($i % 5) == 0) {
                $arrOutput[] = 'FizzBuzz';
            }
            elseif(($i % 3) == 0) {
                $arrOutput[] = 'Fizz';
            }
            elseif(($i % 5) == 0) {
                $arrOutput[] = 'Buzz';
            }
            else {
                $arrOutput[] = (string)$i;
            }

            $i++;
        }

        return $arrOutput;
    }
}

$s = new Solution();
var_dump($s->fizzBuzz(15));