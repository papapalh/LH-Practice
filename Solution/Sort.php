<?php
/**
 *   https://www.cnblogs.com/fivestudy/p/10212306.html
 *   排序算法   时间复杂度   最优          最差            稳定性
 *   冒泡排序   O(n^2)      O(n)         O(n^2)          稳定
 *   选择排序   O(n^2)      O(n^2)       O(n^2)          不稳定
 *   插入排序   O(n^2)      O(n)         O(n^2)          稳定
 *   归并排序   O(n log n)  O(n log n)   O(n log n)      稳定
 *   快速排序   O(n log n)  O(n log n)   O(n^2)          不稳定
 *   堆排序     O(n log n)  O(n log n)   O(n log n)      不稳定
 *   计数排序   O(n+k)      O(n+k)       O(n+k)          稳定
 *   桶排序     O(n^2)      O(n+k)       O(n^2)          稳定
 */

/**
 * 思路
 *     冒泡排序
 *     暴力破解的方法，在需要排序数很多的情况下，时间会很长。
 *     不适用实际的排序
 * 步骤
 *     比较相邻的元素。如果第一个比第二个大，就交换他们两个。
 *     对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
 *     针对所有的元素重复以上的步骤，除了最后一个。
 *     持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
 */
function BubbleSort($arr)
{
    $count = count($arr);

    // 第一次循环只是为了确定需要排序多少次
    for ($i = 1; $i < $count; $i++) {

        // 里面的这次循环才是真正的在排序
        for ($j = 0; $j < ($count-$i); $j++) {
            if ($arr[$j] > $arr[$j+1]) {
                list($arr[$j], $arr[$j+1]) = [$arr[$j+1], $arr[$j]];
            }
        }
    }

    return $arr;
}

/**
 * 思路
 *     选择排序
 *     和冒泡排序类似，不过冒泡是在每个元素中交互，是稳定的。
 *     但是选择排序是根据遍历这个数组获取最大最小来重新操作这个数组。数组中本身的位置都已经改变了
 *     在自己做的时候用了系统的内置函数，其实和自己写差不多
 * 详解
 *     首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
 *     再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
 *     重复第二步，直到所有元素均排序完毕。
 */
function SelectionSort($arr)
{
    $arrOutput = [];
    $count = count($arr);

    for ($i = 0; $i < $count; $i++) {

        $minKey = null;

        foreach ($arr as $k => $a) {
            if (!isset($minKey)) {
                $minKey = $k;
                continue;
            }

            if ($a < $arr[$minKey]) {
                $minKey = $k;
            }
        }

        $arrOutput[] = $arr[$minKey];
        unset($arr[$minKey]);
    }

    return $arrOutput;
}

/**
 * 思路
 *     插入排序
 *     和冒泡排序类似，不过冒泡排序是不管之前是否已经排序完成的
 *     而插入排序，会认为前面的排序是已经排好序的序列
 *     是直接对排好序的序列进行操作
 * 详解
 *     将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
 *     从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
 *     如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。
 */
function InsertSort($arr)
{
    $count = count($arr);

    for($i = 1; $i < $count; $i++) {

        $tmp = $arr[$i];

        //内层循环控制
        for($j = $i-1; $j >= 0;$j--) {
            if($tmp < $arr[$j]) {
                //发现插入的元素要小，交换位置，将后边的元素与前面的元素互换
                $arr[$j+1] = $arr[$j];
                $arr[$j] = $tmp;
            } else {
                //如果碰到不需要移动的元素，由于是已经排序好是数组，则前面的就不需要再次比较了。
                break;
            }
        }
    }
    return $arr;
}

/**
 * 思路
 *     归并排序
 *     对已经排序好的两个数组进行归并，效果最好
 * 详解
 *     比较两个数组的头元素，将比较小的数组头元素出栈，放入新数组
 *     一直循环这个数组，完成数组的排序
 */
function MergeSort($arr)
{
    $a     = [];
    $len   = count($arr);
    $left  = array_slice($arr, 0, floor($len / 2));
    $right = array_slice($arr, floor($len / 2));
    sort($left);
    sort($right);

    $leftPoint  = 0;
    $rightPoint = 0;

    while (true) {

        if (!$left[$leftPoint] && !$right[$rightPoint]) {
            break;
        }
        elseif ($left[$leftPoint] && !$right[$rightPoint]) {
            $a[] = $left[$leftPoint];
            $leftPoint++;
            continue;
        }
        elseif (!$left[$leftPoint] && $right[$rightPoint]) {
            $a[] = $right[$rightPoint];
            $rightPoint++;
            continue;
        }

        if ($left[$leftPoint] > $right[$rightPoint]) {
            $a[] = $right[$rightPoint];
            $rightPoint++;
        }
        elseif($left[$leftPoint] < $right[$rightPoint]) {
            $a[] = $left[$leftPoint];
            $leftPoint++;
        }
        else {
            $a[] = $left[$leftPoint];
            $a[] = $right[$rightPoint];
            $leftPoint++;
            $rightPoint++;
        }
    }

    return $a;
}

/**
 * 思路
 *     快速排序
 *     本质是通过 二分法不断的对数列进行排序
 *     快速排序基本上被认为在相同数量级中，平均性能最好的。
 * 详解
 *     在对数列进行排序时候，挑选一个基本数[flag][一般是第一个]把大于这个数的放右边，小于的放左边。
 *     之后再不停的递归对数列排序，直到最后一个。
 */
function QuickSort($arr)
{
    $len = count($arr);

    // 避免空数组造成无限死循环
    if($len <= 1) {
        return $arr;
    }

    $left  = [];
    $right = [];

    $flag = $arr[0];

    // 开始二分
    for($i = 1; $i < $len; $i++) {
        if($arr[$i] <= $flag) {
            $left[] = $arr[$i];
        }
        else {
            $right[] = $arr[$i];
        }
    }

    // 递归
    $left  = quickSort($left);
    $right = quickSort($right);

    return array_merge($left, [$flag], $right);
}

/**
 * 思路
 *     堆排序
 * 详解
 *     学习堆这种数据结构之后回来补充
 */
function HeapSort($arr)
{
}

/**
 * 思路
 *     计数排序
 *     类似哈希表
 * 详解
 *     首先通过O（n）计算数组出先最大最小，以确定循环次数
 *     开辟额外空间存储每个元素出现位置和次数。
 *     根据最大最小区间循环输出，最后排序完成
 */
function CountingSort($arr)
{
    $min    = $arr[0];
    $max    = $arr[0];
    $bucket    = [];
    $arrOutput = [];

    // O(n) 遍历获取最大最小值,用于确定循环次数
    foreach ($arr as $a) {
        $bucket[$a]++;

        if ($max < $a) {
            $max = $a;
        }
        if ($min > $a) {
            $min = $a;
        }
    }

    // 根据次数取出并排序
    for ($i = $min; $i <= $max; $i++) {
        if ($bucket[$i]) {
            $arrOutput  = array_merge($arrOutput, array_fill(0, $bucket[$i], $i));
        }
    }

    return $arrOutput;
}

/**
 * 思路
 *     桶排序
 *     使用于排序区间不大的排序
 * 步骤
 *     把每个元素放入对应的桶内(类型哈希表)
 *     最后按照顺序把每个桶的元素拿出来，完成排序
 */
function BucketSort($arr)
{
    $max       = $arr[0];
    $bucket    = [];
    $arrOutput = [];

    // O(n) 遍历获取最大最小值,用于确定循环次数
    foreach ($arr as $a) {
        $bucket[$a][] = $a;

        if ($max < $a) {
            $max = $a;
        }
    }

    // 根据次数取出并排序
    for ($i = 0; $i <= $max; $i++) {
        if ($bucket[$i]) {
            $arrOutput  = array_merge($arrOutput, $bucket[$i]);
        }
    }

    return $arrOutput;
}

$arr = [5,6,4,2,8,7,9,1,3];

//var_dump(BubbleSort($arr));

//var_dump(SelectionSort($arr));

//var_dump(InsertSort($arr));

//var_dump(MergeSort($arr));

//var_dump(QuickSort($arr));

//var_dump(CountingSort($arr));

//var_dump(BucketSort($arr));