<?php

/**
 *
 * 堆的定义
堆其实就是一棵完全二叉树（若设二叉树的深度为h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边），

定义为：具有n个元素的序列（h1,h2,...hn），当且仅当满足（hi>=h2i,hi>=h2i+1）或（hi<=h2i,hi<=2i+1） (i=1,2,...,n/2)时称之为堆

大顶堆
堆顶元素（即第一个元素）为最大项，并且（hi>=h2i,hi>=h2i+1）

小顶堆
堆顶元素为最小项，并且（hi<=h2i,hi<=2i+1）
 *
 */


class node
{
    public $l;
    public $r;
    public $v;

    /**
     * node constructor.
     *
     * @param $v
     */
    public function __construct($v)
    {
        $this->v = $v;
    }

}

class heap
{
}

$before = [9, 3, 6, 7, 5, 8, 1, 0, 4, 2];