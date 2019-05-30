<?php

/**
 * 题目
 *     给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
 *     对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
 *     你可以返回任何满足上述条件的数组作为答案。
 * 示例：
 *     输入：[4,2,5,7]
 *     输出：[4,5,2,7]
 *     解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
 */
class Solution
{
    /**
     * 思路
     *     桶排序类似
     *     第一次循环数组，确定奇数偶数，推入对应桶内。
     *     第二次循环数组个数，将对应奇偶从桶内拿出来，放入各个位置
     *     两次循环，效率不好
     * 耗时
     *     执行用时 : 2088 ms, 在Sort Array By Parity II的PHP提交中击败了0.00% 的用户
     *     内存消耗 : 19 MB, 在Sort Array By Parity II的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O（n）
     */
    function sortArrayByParityII_1($A)
    {
        $dict = [];

        // 判断奇偶 - 推入字典
        foreach ($A as $a) {
            if (($a&1) == 0) {
                $dict['2'][] = $a;
                continue;
            }
            $dict['1'][] = $a;
        }

        $aOutput = [];
        $o = true;

        foreach ($A as $a) {

            if ($o) {
                $aOutput[] = array_shift($dict['2']);
                $o = false;
            }
            else {
                $aOutput[] = array_shift($dict['1']);
                $o = true;
            }
        }

        return $aOutput;
    }

    /**
     * 思路
     *     双指针
     *     桶两个指针代表(奇数，偶数)位置
     *     在循环数组每个元素时，根据数值的奇偶性，放入对应指针位置
     *     最后对键排序
     * 耗时
     *     执行用时 : 164 ms, 在Sort Array By Parity II的PHP提交中击败了28.57% 的用户
     *     内存消耗 : 19.4 MB, 在Sort Array By Parity II的PHP提交中击败了100.00% 的用户
     * 时间复杂度
     *     O（n）
     */
    function sortArrayByParityII($A)
    {
        $j = 1; // 奇数指针
        $o = 0; // 偶数指针

        foreach ($A as $a) {
            // 偶数
            if (($a&1) == 0) {
                $aOutput[$o] = $a;
                $o += 2;
                continue;
            }

            $aOutput[$j] = $a;
            $j += 2;
        }

        ksort($aOutput);

        return $aOutput;
    }
}

$s = new Solution();
var_dump($s->sortArrayByParityII([4,2,5,7]));