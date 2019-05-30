<?php

/**
 * 《数据结构 - 线性表》顺序存储结构
 */

class ArrayList 
{
    public $Arr;
    public $Length;

    /**
     * 构建线性表
     */
    public function __construct($Arr)
    {
        $this->Arr    = $Arr;
        $this->Length = count($SqArr);
    }
    
    /**
     * 销毁顺序线性表
     */
    public function Destroy()
    {
        $this->Arr    = NULL;
        $this->Length = 0;
    }

    /**
     * 将线性表重置为空
     */
    public function Init()
    {
        $this->Arr    = [];
        $this->Length = 0;
    }
    
    /**
     * 判断线性表是否为空
     */
    public function Empty()
    {
        if($this->Length == 0) {
            return TRUE;
        }

        return FALSE;
    }

    /**
     * 返回线性表的长度
     */
    public function ListLength(){
        return $this->Length;
    }

    /**
     * 返回线性表中第$index个数据元素
     */
    public function GetIndex($index)
    {
        if($this->Length == 0 || $index < 1 || $index > $this->Length) {
            return 'ERROR';
        }
        return $this->Arr[$index-1];
    }

    /**
     * 在第index的位置插入元素elem
     * 添加之后，其他位置元素也需要随之改变
     */
    public function Insert($index, $elem)
    {
        if($index < 1 || $index > ($this->Length + 1)) {
            return 'ERROR';
        }

        if($index <= $this->Length) {
            for ($i = $this->Length - 1; $i >= $index - 1; $i--) {
                $this->Arr[$i + 1] = $this->Arr[$i];
            }
        }
        $this->Arr[$index - 1] = $elem;
        $this->Length++;
        return 'ok';
    }

    /**
     * 删除第index位置的元素elem
     * 删除之后，其他位置元素也需要随之改变
     */
    public function Delete($index){
        if ($index < 1 || $index > $this->Length + 1) {
            return 'ERROR';
        }
        if ($index < $this->Length) {
            for($i = $index; $i < $this->Length; $i++) {
                $this->Arr[$i - 1] = $this->Arr[$i];
            }
        }
        $this->Length--;
        return $this->Arr[$index - 1];
    }
}