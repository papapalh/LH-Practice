<?php

/**
 * LRU 算法【Least recently used（最近最少使用）】
 * 根据数据的历史访问记录来进行淘汰数据，其核心思想是 "如果数据最近被访问过，那么将来被访问的几率也更高"。
 *
 * 算法污染
 *     缓存了不常用的数据集合，即为缓存污染。
 */

class Factory
{
    public $new;

    public function __construct($class)
    {
        $this->new = new $class;
    }

    public function getInstance()
    {
        return $this->new;
    }
}

/**
 * 思路
 *     单链表实现
 * 原理
 *     单链表
 * 流程
 *     1. 新数据插入到链表头部；
 *     2. 每当缓存命中（即缓存数据被访问），则将数据移到链表头部；
 *     3. 当链表满的时候，将链表尾部的数据丢弃。
 * 优点
 *     实现简单。
 *     当存在热点数据时，LRU的效率很好。
 * 缺点
 *     偶发性的、周期性的批量操作会导致LRU命中率急剧下降，缓存污染情况比较严重。
 *     命中时需要遍历链表，找到命中的数据块索引，然后需要将数据移到头部。
 */
class Lru
{
    public static $lruList = []; // 顺序存储单链表结构
    public static $maxLen  = 5;  // 链表允许最大长度
    public static $nowLen  = 0;  // 链表当前长度

    /**
     * LRU_1 constructor.
     * 由于 PHP 不是常驻进程程序，所以链表初始化可以通过 Mysql/Redis 实现
     */
    public function __construct()
    {
        self::$lruList = [];
        self::$nowLen  = count(self::$lruList);
    }

    /**
     * 获取 key => value
     * @param $key
     * @return null
     */
    public function get($key)
    {
        $value = null;

        // lru 队列为空，直接返回
        if (!self::$lruList) {
            self::$lruList[] = [$key => $this->getData($key)]; // 根据实际项目情况获取数据
            self::$nowLen++;
            return $value;
        }

        // 查找 lru 缓存
        for ($i = 0; $i < self::$nowLen; $i++) {

            // 如果存在缓存，则直接返回，并将数据重新插入链表头部
            if (isset(self::$lruList[$i][$key])) {

                $value = self::$lruList[$i][$key];

                unset(self::$lruList[$i]);

                array_unshift(self::$lruList, [$key => $value]);

                break;
            }
        }


        // 如果没有找到 lru 缓存
        if (!isset($value)) {

            // 插入头部
            array_unshift(self::$lruList, [$key => $this->getData($key)]); // 根据实际项目情况获取数据
            self::$nowLen++;

            if (self::$nowLen > self::$maxLen) {
                self::$nowLen--;
                array_pop(self::$lruList);
            }
        }

        return $value;

    }

    /**
     * 输出 Lru 队列
     */
    public function echoLruList()
    {
        var_dump(self::$lruList);
    }

    /**
     * 根据真实环境获取数据
     * @param $key
     * @return string
     */
    public function getData($key)
    {
        return 'data';
    }
}

/**
 * 思路
 *     为了避免 LRU 的 '缓存污染' 问题
 *     增加一个队列来维护缓存出现的次数。其核心思想是将“最近使用过1次”的判断标准扩展为“最近使用过K次”。
 * 原理
 *     相比LRU，LRU-K需要多维护一个队列，用于记录所有缓存数据被访问的历史。
 *     只有当数据的访问次数达到K次的时候，才将数据放入缓存。
 *     当需要淘汰数据时，LRU-K会淘汰第K次访问时间距当前时间最大的数据
 * 流程
 *     1.数据第一次被访问，加入到访问历史列表；
 *     2.如果数据在访问历史列表里后没有达到K次访问，则按照一定规则 LRU淘汰；
 *     3.当访问历史队列中的数据访问次数达到K次后，将数据索引从历史队列删除，将数据移到缓存队列中，并缓存此数据，缓存队列重新按照时间排序；
 *     4.缓存数据队列中被再次访问后，重新排序；
 *     5.需要淘汰数据时，淘汰缓存队列中排在末尾的数据，即：淘汰“倒数第K次访问离现在最久”的数据。
 * 优点
 *     LRU-K降低了“缓存污染”带来的问题，命中率比LRU要高。
 * 缺点
 *     LRU-K队列是一个优先级队列，算法复杂度和代价比较高。
 *     由于LRU-K还需要维护历史队列，所以消耗的内存会更多。
 */
class Lru_K
{
    public static $historyList = []; // 访问历史队列
    public static $lruList     = []; // 顺序存储单链表结构
    public static $maxLen      = 5;  // 链表允许最大长度
    public static $nowLen      = 0;  // 链表当前长度

    /**
     * LRU_K constructor.
     * 由于 PHP 不是常驻进程程序，所以链表初始化可以通过 Mysql/Redis 实现
     */
    public function __construct()
    {
        self::$lruList     = [];
        self::$historyList = [];
        self::$nowLen      = count(self::$lruList);
    }

    /**
     * 获取 key => value
     * @param $key
     * @return null
     */
    public function get($key)
    {
        $value = null;

        // 查找 lru 缓存
        for ($i = 0; $i < self::$nowLen; $i++) {

            // 如果存在缓存，则直接返回，并将数据重新插入链表头部
            if (isset(self::$lruList[$i][$key])) {

                $value = self::$lruList[$i][$key];

                unset(self::$lruList[$i]);

                array_unshift(self::$lruList, [$key => $value]);

                break;
            }
        }

        // 如果没有找到 lru 缓存, 则进入历史队列进行计数，当次数大于等于5时候，进入缓存队列
        if (!isset($value)) {
            self::$historyList[$key]++;

            $value = $this->getData($key);

            // 进入缓存队列
            if (self::$historyList[$key] >= 5) {
                array_unshift(self::$lruList, [$key => $value]); // 根据实际项目情况获取数据
                self::$nowLen++;

                if (self::$nowLen > self::$maxLen) {
                    self::$nowLen--;
                    array_pop(self::$lruList);
                }

                unset(self::$historyList[$key]); // 历史队列推出
            }
        }

        return $value;

    }

    /**
     * 输出 Lru 队列
     */
    public function echoLruList()
    {
        var_dump(self::$lruList);
        var_dump(self::$historyList);
    }

    /**
     * 根据真实环境获取数据
     * @param $key
     * @return string
     */
    public function getData($key)
    {
        return 'data';
    }
}

//$Lru = new Factory('Lru_K');
$Lru = new Factory('Lru_K');
$Lru = $Lru->getInstance();

$Lru->get('five');
$Lru->get('two');  // 由于队列长度不够被淘汰
$Lru->get('one');
$Lru->get('three');
$Lru->get('five'); // 优先级高
$Lru->get('six');
$Lru->get('seven');
$Lru->echoLruList();
