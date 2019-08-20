<?php
/**
 * 创建 Redis 单例
 * @return mixed
 */
class redisInstance
{
    static $redis;

    private function __construct()
    {
    }

    public static function getInstance()
    {
        if (self::$redis) {
            return self::$redis;
        }

        self::$redis = new \Redis();

        return self::$redis;
    }

}

class redisLock
{
    public $reids;

    public function __construct()
    {
        $this->reids = redisInstance::getInstance();
    }

    /**
     * 键加锁，默认加锁时间为 1分钟 = 60 * 1000(毫秒)
     * @param $key
     * @param int $ttl
     * @return bool
     */
    public function lock($key, $ttl = 60000)
    {
        $lockKey = $key . '_lock';

        // 通过setNx命令拿到锁
        $lock = null;

        // 没有拿到锁，则一直循环等待锁资源释放
        while (!$lock) {
            $lock = $this->reids->set($lockKey, 1, ['NX', 'PX' => $ttl]);
        }

        return true;
    }

    /**
     * 释放锁
     * @param $key
     * @return bool
     */
    public function unLock($key)
    {
        $lockKey = $key . '_lock';

        $lock = false;

        // 释放锁
        while (!$lock) {
            $lock = $this->reids->del($lockKey);
        }

        return true;
    }
}

$redis = redisInstance::getInstance();
$redisLock = new redisLock();

/**
 * 例如，买商品，进行库存递减
 *     1：对库存加锁
 *     2：递减
 *     3：释放锁
 * 影响
 *     加锁导致的并发度降低
 */
$key = 'stock';
$redisLock->lock($key);
$stock = $redis->get($key);
if ($redis->get($key) <= 0) {
    $redisLock->unLock($key);
    return false;
}
$redis->decr($key);
$redisLock->unLock($key);