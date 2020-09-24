<?php


var_dump(reverseLinkedList());

/**
 * K 个一组反转链表
 * @param array $arr
 * @return array
 */
function reverseLinkedList(array $arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], $k = 3)
{
    if ($k == 0 || empty($arr)) return [];

    $chunk = array_chunk($arr, $k);

    $res = [];

    foreach ($chunk as $v){

        $res[] = rll($v);
    }

    $data = [];
    foreach ($res as $v){
        foreach ($v as $vi){

            $data[] = $vi;
        }
    }
    
    return $data;

}

function rll($arr)
{
    $count = count($arr);

    if (count($arr) <= 1){

        return $arr;
    }

    $head = 0;
    $rear = $count - 1;

    while (1){
        if ($rear - $head <= 1) return $arr;

        $tmp = $arr[$rear];
        $arr[$rear] = $arr[$head];
        $arr[$head] = $tmp;
        $head ++;
        $rear --;
    }

    return $arr;
}

