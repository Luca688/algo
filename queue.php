<?php

class node
{
    public $val;
    public $next;

    public function __construct($val, $next)
    {
        $this->val = $val;
        $this->next = $next;
    }
}

class queue
{
    private $maxSize = null;
    private $useSize = 0;

    private $header = null;

    public function __construct($maxSize)
    {
        $this->maxSize = $maxSize;
        $this->header = new node(null, null);
    }

    public function isFull()
    {
        return $this->maxSize === $this->useSize;
    }

    public function add($val)
    {
        $addNode = new node($val, null);

        $node = &$this->header;
        while ($node->next){
            $node = &$node->next;
        }

        $node->next = $addNode;
        ++$this->useSize;

        return $this;
    }

    public function isEmpty()
    {
        return 0 === $this->useSize;
    }

    public function delete()
    {
        $node = $this->header->next->next;
        echo $this->header->next->val . '--';
        $this->header->next = $node;

        --$this->useSize;
        return $this;

    }
}

$aa = (new queue(5))->add(1)->add(2)->add(3)->add(4)->add(5)
    ->delete()->delete();

var_dump($aa);