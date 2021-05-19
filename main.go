package main

import (
	"fmt"
)

/**
在一个二维数组中（每个一维数组的长度相同），
每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
  [1,2,8,9],
  [2,4,9,12],
  [4,7,10,13],
  [6,8,11,15]
]
给定 target = 7，返回 true。

给定 target = 3，返回 false。：
*/
func main() {
	fmt.Println("vim-go")
	//
	//data := [][]int{
	//	{1, 2, 8, 9},
	//	{2, 4, 9, 12},
	//	{4, 7, 10, 13},
	//	{6, 8, 11, 15},
	//}
	//
	//target := 5
	//res := Find2DimensionArray(target, data)

	str := "We Are Happy"
	res := RepliceSpace(str)
	fmt.Printf("res:%v", res)

}

type Node struct {
	Val int
	Next *Node
}

/**
输入一个链表，按链表从尾到头的顺序返回一个ArrayList。
{67,0,24,58} [58,24,0,67]
 */
//func ReverseLinkList(slice []int) []int {
//
//	var link *Node
//	var current *Node
//	for _, value := range slice {
//		cur := Node{Val:value}
//		if link == nil {
//			link = &cur
//		}else {
//			current.Next = &cur
//		}
//		current = &cur
//	}
//
//	//三指针
//	var left, mid, right *Node
//
//}

/**
请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
 */
func RepliceSpace(str string) string {
	res := ""
	//todo
	return res
}

func Find2DimensionArray(target int, array [][]int) bool {

	res := false

	maxRowIndex := len(array) - 1
	maxColumnIndex := len(array[0]) - 1

	//大于或者小于最大最小元素
	if target < array[0][0] || target > array[maxRowIndex][maxColumnIndex] {

		return false
	}
	//等于最大或最小元素
	if target == array[0][0] || target == array[maxRowIndex][maxColumnIndex] {

		return true
	}

	currentRow := maxRowIndex
	for {
		if currentRow < 0 {
			break;
		}
		//等于当前最小元素
		if target == array[currentRow][0] {
			res = true
			break
		}
		//大于当前列最小元素
		if target > array[currentRow][0] {
			//当前组内搜索
			for _, val := range array[currentRow] {
				//找到元素
				if val == target {
					res = true
					break
				}
			}
			currentRow -=1
		} else {
			currentRow -= 1
		}

	}

	return res
}
