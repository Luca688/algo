<?php


class node
{
    public function __construct($val, $next)
    {
        $this->val = $val;
        $this->next = $next;
    }

    public $val;
    public $next;
}



class linkedList
{
    private $length = 0;

    private $header = null;

    public function insert($val, $key = null)
    {
        if ($this->header === null) return $this->insertHeader($val);

        if ($key >= $this->length || $key === null) return $this->insertFoot($val);

        return $this->insertMiddle($val, $key);
    }

    private function insertMiddle($val, $key)
    {
        $node = &$this->header;
        $i = 1;
        while (1){
            if ($key === $i){
                $insertNode = new node($val, $node->next);
                $node->next = $insertNode;
                break;
            }else{
                $i++;
                $node = &$node->next;
            }
        }
        ++$this->length;
        return $this;
    }

    private function insertHeader($val)
    {
        $this->header = new node($val, null);
        $this->length += 1;
        return $this;
    }

    private function insertFoot($val)
    {
        $node = $this->header;
        while ($node->next){

            $node = $node->next;
        }

        $node->next = new node($val, null);
        ++$this->length;
        return $this;
    }

    public function findByKey($key)
    {

        $i = 1;
        $val = null;
        $node = $this->header;
        while (true){
            if ($i === $key){
                $val = $node->val;
                break;
            }
            $i ++;
            $node = $node->next;
            if ($node == null) break;
        }

        return $val;
    }

    public function find($val)
    {
        $i = 1;
        $valKey = null;
        $node = $this->header;
        while (true){
            if ($node->val === $val){
                $valKey = $i;
                break;
            }
            $i++;
            $node = $node->next;
            if ($node == null) break;
        }

        return $valKey;
    }

    public function getLength()
    {
        return $this->length;
    }

    public function delete($key)
    {
        $node = &$this->header;

        if ($key === 1){
            $node = $node->next;

            return $this;
        }

        $i = 1;
        while (1){

            if ($i === $key -1){
                if ($node->next){

                    if ($node->next->next) $next = &$node->next->next;
                    else $next = null;

                    $node->next = $next;
                }
                break;
            }
            $i ++;
            $node = &$node->next;
        }

        return $this;
    }

}


$aa = (new linkedList())->insert(1)->insert(2)->insert(3)->insert(4)
->delete(3);

var_dump($aa);