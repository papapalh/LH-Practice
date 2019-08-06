<?php
/**
 * 雪花算法
 *    其核心思想就是：
 *    使用一个 64 bit 的 long 型的数字作为全局唯一 id。
 *    在分布式系统中的应用十分广泛，且ID 引入了时间戳，基本上保持自增的。
 * 产生公式
 *    | 0(最高位预留) | 时间戳(41位) | 机器ID(10位) | 随机序列(12位) |
 */
class IdCreate
{
    const max12bit = 4095;

    public static function createOnlyId()
    {
        // 获取微秒时间戳(42位)，截取并转化为 41位二进制
        $microtime = decbin(floor(microtime(true) * 1000));

        // 10bit 的机器号，由集群产出
        $machineId = self::machine();

        // 12bit 的随机数
        $random = str_pad(decbin(mt_rand(0, self::max12bit)), 12, "0", STR_PAD_LEFT);

        // 拼接
        $base = '0' . $microtime . $machineId . $random;

        // 十进制 返回
        return bindec($base);
    }

    /**
     * 集群
     * @param int $machineId
     * @return string
     */
    public static function machine($machineId = 0)
    {
        return str_pad($machineId, 10, "0", STR_PAD_LEFT);
    }
}

$cast_id = IdCreate::createOnlyId();
var_dump($cast_id);