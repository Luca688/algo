<?php


$before = [5, 8, 9, 3, 10, 2, 1, 6, 7, 4, 0];

print_r(mergeArray([1,45,70], [2, 55, 80]));


/***
 * 合并有序序列
 * @param $left
 * @param $right
 *
 * @return mixed
 */
function mergeArray($left, $right)
{
    if (empty($left)) return $right;
    if (empty($right)) return $left;

    $after = [];
    $leftIndex = 0;
    $rightIndex = 0;

//    [1,45,70]
//    [2, 55, 80]

    while (true){

        if (count($after) == count($left) + count($right)) break;


        if ($leftIndex == count($left)) {
            $after[] = $right[$rightIndex];
            $rightIndex++;
            continue;
        }
        if ($leftIndex == count($left)) {
            $after[] = $left[$leftIndex];
            $leftIndex++;
            continue;
        }

        if ($left[$leftIndex] >= $right[$rightIndex]) {
            $after[] = $right[$rightIndex];
            $rightIndex ++;
        }elseif($left[$leftIndex] < $right[$rightIndex]){
            $after[] = $left[$leftIndex];
            $leftIndex ++;
        }
    }

    return $after;
}


function heap($before)
{
    //todo
    return $before;
}

function quick($before)
{
    if (count($before) <= 1) return $before;

    $flag = $before[0];
    $left = [];
    $right = [];

    for ($i = 1; $i< count($before); $i++){

        if ($before[$i] >= $flag){
            $right[] = $before[$i];
        } else {
            $left[] = $before[$i];
        }
    }

    $left = quick($left);
    $right = quick($right);
    return array_merge($left, [$flag],$right);
}

/**
 * 比较两个相邻的元素，将值大的元素交换到右边 大的往下沉
 *
 * @param $before
 *
 * @return mixed
 */
function bubble($before)
{
    $count = count($before);


    for ($i = 0; $i < $count; $i++){

        for ($j = 0; $j < $count-$i-1; $j++){

            if ($before[$j] > $before[$j+1]){

                $temp = $before[$j];
                $before[$j] = $before[$j+1];
                $before[$j+1] = $temp;
            }
        }
    }

    return $before;
}


function insert($before)
{
    for ($i = 1; $i < count($before); $i ++){

        $temp = $before[$i];
//        if ($temp >= $before[$i-1]) continue;
        for ($j = $i-1; $j >=0; $j --){
            if ($before[$j] > $temp){

                $before[$j+1] = $before[$j];
                $before[$j] = $temp;
            }else{
                break;
//                $before[$j] = $temp;
            }
        }
    }

    return $before;
}


/**
 * 一张牌插入到其他已经有序的牌中的适当位置。
 * @param $before
 */
function insertTemp($before)
{
    $after = [];

    for ($i = 0; $i < count($before); $i++){

        $val = $before[$i];

        if (count($after) === 0) {

            $after[] = $val;
        }elseif($val <= $after[0]){

            $after = array_merge([$val], $after);
        }elseif($val >= $after[count($after)-1]){

            $after = array_merge($after, [$val]);
        } else {

            for ($j = 0; $j < count($after) - 1; $j++) {

                if ($val > $after[$j] && $val < $after[$j + 1]) {

                    $less    = array_slice($after, 0, $j + 1);
                    $greater = array_slice($after, $j + 1);
                    $after   = array_merge($less, [$val], $greater);
                    break 1;
                }
            }
        }
    }

    return $after;
}

/**
 * 找到数组中最小的那个元素，其次，将它和数组的第一个元素交换位置(如果第一个元素就是最小元素那么它就和自己交换)。其次，在剩下的元素中找到最小的元素，将它与数组的第二个元素交换位置。如此往复，直到将整个数组排序。
 *
 * @param $before
 */
function select($before)
{
    for ($i = 0; $i < count($before); $i ++){
        $min = $before[$i];
        $minIndex = $i;

        for ($j = $i+1; $j < count($before); $j ++){

            if ($before[$j] < $min) {
                $min = $before[$j];
                $minIndex = $j;
            }
        }

        if ($minIndex !=$i){

            $temp = $before[$minIndex];
            $before[$minIndex] = $before[$i];
            $before[$i] = $temp;
        }
    }

    return $before;
}




function getMin($input)
{
    if (count($input) <= 0) return [];

    $data['min_val']   = $input[0];
    $data['min_index'] = 0;

    foreach ($input as $k => $v) {

        if ($v < $data['min_val']) {
            $data['min_val']   = $v;
            $data['min_index'] = $k;
        }
    }

    return $data;
}