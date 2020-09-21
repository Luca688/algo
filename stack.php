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


class stack
{
    private $maxSize = 0;

    private $useSize = 0;

    private $top = null;

    public function __construct($maxSize)
    {
        $this->maxSize = $maxSize;
    }

    public function isFull()
    {
        return $this->maxSize === $this->useSize;
    }

    public function push($val)
    {
        if ($this->isFull()) return $this;


        $next = null;
        if ($this->top != null){
            $next = $this->top;
        }

        ++$this->useSize;
        $this->top = new node($val, $next);

        return $this;
    }

    public function pop()
    {
        if ($this->isEmpty()) return $this;

        echo 'pop' . $this->top->val;

        $this->top = $this->top->next;
        --$this->useSize;

        return $this;
    }

    public function isEmpty()
    {
        return $this->useSize === 0;
    }
}

$aa = (new stack(5))->push(1)->push(2)->push(3)->push(4)->push(5)
    ->pop()->pop()->pop()->pop()->pop();
var_dump($aa);