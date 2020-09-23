<?php

class node
{

    public $val;
    public $l;
    public $r;

    /**
     * node constructor.
     *
     * @param $l
     * @param $r
     */
    public function __construct($val, $l = null, $r = null)
    {
        $this->val        = $val;
        $this->l  = $l;
        $this->r = $r;
    }

}

class bst
{
    public $_root = null;

    public function find($val)
    {
        $res = false;

        $node = $this->_root;

        while ($node){
            if ($node->val == $val){
                $res = true;
                break;
            }else{

                if ($val <= $node->val){
                    $node = $node->l;
                }else{
                    $node = $node->r;
                }
            }
        }

        if ($res){
            echo 'has-' . $val . "\r\n";
        }else{
            echo 'non-' . $val . "\r\n";
        }

        return $this;
    }

    public function findRecursive($bst, $val)
    {
        if ($bst){

            if ($bst->val == $val){

                return true;
            }elseif ($val > $bst->val){
                return $this->findRecursive($bst->r, $val);
            }elseif ($val < $bst->val){
                return $this->findRecursive($bst->l, $val);
            }
        }else{

            return false;
        }
    }


    public function insert($val)
    {
        if ($this->_root === null) {

            $this->_root =new node($val);
            return $this;
        }

        $node = &$this->_root;

        while (true){
            if ($val <= $node->val){
                if ($node->l === null) {
                    $node->l = new node($val);
                    break;
                }
                $node = &$node->l;
            }else{
                if ($node->r === null) {
                    $node->r = new node($val);
                    break;
                }
                $node = &$node->r;
            }
        }

        return $this;
    }

    public function delete($val)
    {
        //todo
        $node = &$this->_root;
        $parent = null;

        while ($node){
            if ($val == $node->val){

                if ($node->l == null && $node->r == null){
                    if ($parent->l->val == $val){
                        $parent->l = null;
                    }elseif ($parent->r->val == $val){
                        $parent->r = null;
                    }
                }
                elseif($node->l == null){

                }
                elseif($node->r == null){

                }

                break;
            }elseif ($val < $node->val){
                $parent = $node;
                $node = $node->l;
            }else{
                $parent = $node;
                $node = $node->r;
            }
        }

        return $this;
    }

    /***
     * 中左右
     */
    public function preOrder($node = null)
    {
        if ($node == 'root') $node = $this->_root;

        if ($node === null) return;

        echo $node->val . "\r\n";

        $this->preOrder($node->l);
        $this->preOrder($node->r);
    }


    /***
     * 左中右
     */
    public function inOrder($node = null)
    {
        if ($node == 'root') $node = $this->_root;

        if ($node === null) return;

        $this->inOrder($node->l);
        echo $node->val . "\r\n";
        $this->inOrder($node->r);
    }

    /***
     * 左右中
     */
    public function postOrder($node = null)
    {
        if ($node == 'root') $node = $this->_root;

        if ($node === null) return;

        $this->postOrder($node->l);
        $this->postOrder($node->r);
        echo $node->val . "\r\n";
    }

    public function levelOrder()
    {
        $current = [];
        $current[] = $this->_root;

        while (count($current) > 0){

            $next = [];
            foreach ($current as $node){
                echo $node->val . '--';
                if ($node->l) $next[] = $node->l;
                if ($node->r) $next[] = $node->r;
            }
            $current = $next;
            echo "\r\n" . '--------' . "\r\n";
        }

        return $this;
    }

    /**
     * 中左右
     * @return $this
     */
    public function preOrderIteration()
    {
        $node = $this->_root;
        $res = [];
        $res[] = $node->val;
        $stack = [];
        array_push($stack,$node->r);
        array_push($stack,$node->l);
        while ($stack){
            $current = array_pop($stack);
            $res[] = $current->val;
            if ($current->r) array_push($stack, $current->r);
            if ($current->l) array_push($stack, $current->l);
        }

        echo implode(',', $res);

        return $this;
    }

    /***
     * 左中右
     */
    public function inOrderIteration($tree)
    {
        $res = [];

        $stack = [];
        $current = $tree;

        while ($current || $stack){
            while ($current) {
                array_push($stack, $current);
                $current = $current->l;
            }

            while (!$current){
                $current = array_pop($stack);
                if (!$current) break;
                $res[] = $current->val;
                $current = $current ? $current->r : null;
            }

        }

        return $res;
    }

    public function getTree()
    {
        return $this->_root;
    }

    public function getHeight($tree)
    {
    }

    public function getHeightByLevel($tree)
    {
        if (!$tree) return 0;

        $next = [];
        array_push($next, $tree);


        $height = 0;
        while (!empty($next)){

            $tmpTemp = [];
            while (1) {
                $current = array_pop($next);
                if ($current == null) {
                    $height ++;
                    $next = $tmpTemp;

                    break;
                }else{

                    if ($current->l){
                        $tmpTemp[] = $current->l;
                    }
                    if($current->r){
                        $tmpTemp[] = $current->r;
                    }
                }
            }
        }

        return $height;
    }

    public function postOrderIteration($tree)
    {
        $current = $tree;
        $res = [];
        $stack = [];

        while ($current){
            array_push($stack, $current);
            $current = $current->l;
        }


        while ($stack){
            $current = array_pop($stack);
            //叶子节点
            if ($current->l == null && $current->r == null){

                $res[] = $current->val;
            }
            //左右都有
            elseif ($current->l && $current->r){
                //左右已输出
                if (in_array($current->r->val, $res)){
                    $res[] = $current->val;
                }else{

                    array_push($stack, $current);
                    array_push($stack, $current->r);
                    if (!in_array($current->l->val, $res)){
                        array_push($stack, $current->l);
                    }
                }
            }
            //只有左
            elseif ($current->l){

                if (in_array($current->l->val, $res)){
                    $res[] = $current->val;
                }else{
                    array_push($stack, $current);
                    array_push($stack, $current->l);
                }
            }
            //只有右
            else{
                if (in_array($current->r->val, $res)){
                    $res[] = $current->val;
                }else{
                    array_push($stack, $current);
                    array_push($stack, $current->r);
                }
            }
        }


        return $res;
    }


    public function getHeightByRecursive($root)
    {
        if ($root === null) return 0;

        $l = $this->getHeightByRecursive($root->l);
        $r = $this->getHeightByRecursive($root->r);

        return 1+($l > $r ? $l : $r);

    }

}

$aa = (new bst())->insert(6)->insert(4)->insert(8)->insert(2)->insert(5)->insert(7)->insert(9)->insert(1)->insert(3)->insert(10);

$tree = $aa->getTree();
//echo json_encode($tree);

 print_r($aa->getHeightByRecursive($tree));





/**
 *                      6
 *                4             8
 *              2   5         7  9
 *            1   3                 10
 */
//->insert(6)->insert(4)->insert(8)->insert(2)->insert(5)->insert(7)->insert(9)->insert(1)->insert(3)

//var_dump(json_encode($aa->_root));