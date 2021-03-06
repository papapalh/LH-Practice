<?php

/**
 * 一致性哈希实现接口
 * Interface ConsistentHash
 */
interface ConsistentHash
{
    //将字符串转为hash值
    public function cHash($str);
    //添加一台服务器到服务器列表中
    public function addServer($server);
    //从服务器删除一台服务器
    public function removeServer($server);
    //在当前的服务器列表中找到合适的服务器存放数据
    public function lookup($key);
}

/**
 * 具体一致性哈希实现
 * Class MyConsistentHash
 */
class MyConsistentHash implements ConsistentHash
{
    public $serverList    = []; //服务器列列表
    public $virtualPos    = []; //虚拟节点的位置
    public $virtualPosNum = 5;  //每个节点对应5个虚节点

    /**
     * 将字符串转换成32位无符号整数hash值
     * @param $str
     * @return int
     */
    public function cHash($str)
    {
        return sprintf('%u', crc32(md5($str)));
    }

    /**
     * 在当前的服务器列表中找到合适的服务器存放数据
     * @param $key 键名
     * @return mixed 返回服务器IP地址
     */
    public function lookup($key)
    {
        $point       = $this->cHash($key); //落点的hash值
        $finalServer = current($this->virtualPos); //先取圆环上最小的一个节点当成结果

        foreach($this->virtualPos as $pos=>$server)
        {
            if($point <= $pos)
            {
                $finalServer = $server;
                break;
            }
        }

        reset($this->virtualPos);//重置圆环的指针为第一个
        return $finalServer;
    }

    /**
     * 添加一台服务器到服务器列表中
     * @param $server string
     * @return bool
     */
    public function addServer($server)
    {
        if (!$server) {
            return FALSE;
        }

        if(!isset($this->serverList[$server]))
        {
            // 挂载指定虚拟结节点
            for($i = 0; $i < $this->virtualPosNum; $i++)
            {
                $pos = $this->cHash($server . '-' . $i);
                $this->virtualPos[$pos] = $server;
                $this->serverList[$server][] = $pos;
            }

            ksort($this->virtualPos,SORT_NUMERIC);
        }
        return TRUE;
    }
    /**
     * 移除一台服务器（循环所有的虚节点，删除值为该服务器地址的虚节点）
     * @param $key
     * @return bool
     */
    public function removeServer($key)
    {
        if(isset($this->serverList[$key]))
        {
            //删除对应虚节点
            foreach($this->serverList[$key] as $pos)
            {
                unset($this->virtualPos[$pos]);
            }
            //删除对应服务器
            unset($this->serverList[$key]);
        }
        return TRUE;
    }
}

$hash = new MyConsistentHash();

$hash->addServer('127.0.0.1');

var_dump($hash->virtualPos);