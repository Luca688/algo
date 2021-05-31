package algo

import (
	"errors"
	"math"
	"strings"
)

//输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
//[1,2,3,4,5,6,7],[3,2,4,1,6,5,7] {1,2,5,3,4,6,7}
// func reConstructBinaryTree(pre []int, vin []int) *BinaryTreeNode {
// 	// write code here
// }

//从尾到头打印链表
//输入一个链表，按链表从尾到头的顺序返回一个ArrayList。
//{67,0,24,58}
//[58,24,0,67]
func PrintListFromTailToHead(head *ListNode) []int {
	var res []int

	if head == nil {
		return res
	}
	currentNode := head
	for currentNode != nil {
		tmp := []int{currentNode.Val}
		res = append(tmp, res...)
		currentNode = currentNode.Next
	}

	return res
}

//替换空格
//请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
func ReplaceSpace(s string) string {

	res := strings.ReplaceAll(s, " ", "%20")
	return res
}

//二维数组中的查找
/**
在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
  [1,2,8,9],
  [2,4,9,12],
  [4,7,10,13],
  [6,8,11,15]
]
给定 target = 7，返回 true。
给定 target = 3，返回 false。
**/
func Find(target int, array [][]int) bool {

	//空数组
	if len(array) == 0 {
		return false
	}
	//只有一行
	if len(array) == 1 {
		res, _ := BinarySearch(array[0], target)
		return res
	}

	columnLength, columnMax := len(array), len(array)-1

	rowMax := len(array[0]) - 1

	//小于最小或大于最大元素
	if target < array[0][0] || target > array[columnMax][rowMax] {
		return false
	}
	//等于最大最小元素
	if target == array[0][0] || target == array[columnMax][rowMax] {
		return true
	}

	//第一列切片
	var column []int
	for i := 0; i < columnLength; i++ {
		column = append(column, array[i][0])
	}

	//行
	min := 0
	max := columnMax
	var targetRow []int

	for {
		if target == column[min] || target == column[max] {
			return true
		}

		//无法确定最后两行哪个大
		if max-min == 1 {
			for k, v := range array[min] {
				targetRow = append(targetRow, v)
				targetRow = append(targetRow, array[max][k])
			}

			break
		}

		mid := int(math.Floor(float64((max + min) / 2)))
		if target >= column[mid] {
			min = mid
		} else {
			max = mid
		}
	}

	//行查找
	res, _ := BinarySearch(targetRow, target)

	return res
}

//删除链表中重复的节点
//在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
// 例如，链表1->1->1->1->1->1->2->3->3->3->3->4->4->5 处理后为 1->2->5
func SortLinkListRepeat(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	currentNode := head
	var lastNode *ListNode

	for currentNode != nil {

		if currentNode.Next == nil {
			break
		} else {
			//当前元素值与下一个相等
			if currentNode.Val == currentNode.Next.Val {
				//遍历丢弃中间相同的元素
				for {
					if currentNode.Next != nil && currentNode.Val == currentNode.Next.Val {
						currentNode.Next = currentNode.Next.Next
					} else {
						break
					}
				}
				//头节点就有相同元素
				if lastNode == nil {
					lastNode = currentNode.Next
					head = lastNode
				} else {
					lastNode.Next = currentNode.Next
				}
				currentNode = currentNode.Next

			} else {
				lastNode = currentNode
				currentNode = currentNode.Next
			}
		}
	}

	return head
}

//输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
//{1,3,5},{2,4,6} 返回值 {1,2,3,4,5,6}
//{2,11,77,79,80}{3,4,5,82,83,84} 返回  {2,3,4,5,11,77,79,80,83,84}

func IncreasingLinkMerge(head1 *ListNode, head2 *ListNode) *ListNode {

	var res *ListNode
	var resTail *ListNode
	leftCurrent := head1
	rightCurrent := head2

	for {
		if leftCurrent == nil {
			resTail.Next = rightCurrent
			break
		}
		if rightCurrent == nil {
			resTail.Next = leftCurrent
			break
		}

		if leftCurrent.Val <= rightCurrent.Val {

			if res == nil {
				res = leftCurrent
			} else {
				resTail.Next = leftCurrent
			}
			resTail = leftCurrent
			leftCurrent = leftCurrent.Next
		} else {

			if res == nil {
				res = rightCurrent
			} else {
				resTail.Next = rightCurrent
			}
			resTail = rightCurrent
			rightCurrent = rightCurrent.Next
		}
	}

	return res

}

//输入一个链表，输出该链表中倒数第k个结点。
//如果该链表长度小于k，请返回空。
//{1,2,3,4,5},1   {5}

func ReciprocalKNode(head *ListNode, k int) (res int, err error) {

	fast := 1
	fastNode := head
	slow := 1
	slowNode := head

	if head == nil {
		return 0, errors.New("empty link")
	}
	if k <= 0 {
		return 0, errors.New("k less 0")
	}
	//note fast <= k
	for fast = 1; fast < k; fast++ {
		if fastNode == nil {
			break
		} else {
			fastNode = fastNode.Next
		}

	}

	if fast < k {

		return 0, errors.New("length less then k")
	}

	for {
		if fastNode.Next == nil {
			break
		} else {
			slow++
			fast++
			slowNode = slowNode.Next
			fastNode = fastNode.Next
		}
	}

	if k == 1 {
		return fastNode.Val, nil
	} else {
		return slowNode.Val, nil
	}

}

//给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。
//https://my.oschina.net/songjilong/blog/3159816
func EntryNodeOfLoop(head *ListNode) *ListNode {

	//1.遍历一遍有重复的就是环
	allNode := make(map[*ListNode]*ListNode)
	currentNode := head
	var res *ListNode
	for currentNode != nil {
		_, ok := allNode[currentNode]
		if ok == false {
			allNode[currentNode] = currentNode
			currentNode = currentNode.Next
		} else {
			res = currentNode
			break
		}
	}
	return res

	//2. 双指针
}

//输入两个无环的单链表，找出它们的第一个公共结点。
func FindFirstCommonNode(left *ListNode, right *ListNode) *ListNode {

	var res *ListNode
	allNode := make(map[*ListNode]*ListNode)

	leftNode := left
	rightNode := right

	for leftNode != nil && rightNode != nil {

		_, hasLeft := allNode[leftNode]
		if hasLeft {
			res = leftNode
			break
		} else {
			allNode[leftNode] = leftNode
			leftNode = leftNode.Next
		}
		_, hasRight := allNode[rightNode]
		if hasRight {
			res = rightNode
			break
		} else {
			allNode[rightNode] = rightNode
			rightNode = rightNode.Next
		}
	}

	return res

}

//输入一个链表，反转链表后，输出新链表的表头。
//{1,2,3} {3,2,1}
//{1,2,3,4,5}
//转成栈再转链表
//todo

//输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针random指向一个随机节点），
//请对此链表进行深拷贝，并返回拷贝后的头结点。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）
//todo

//todo 用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

//todo 定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））。

//todo 输入一棵二叉树，判断该二叉树是否是平衡二叉树。
//具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。

//在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符
//,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）

//从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。
//{8,6,10,5,7,9,11}  [[8],[6,10],[5,7,9,11]]
func (root *BinaryTreeNode) Print() [][]int {

	// var res [][]int
	res := make([][]int, 0, 100)

	currentNodeList := []*BinaryTreeNode{root}
	for len(currentNodeList) > 0 {
		var nextNodeList []*BinaryTreeNode
		var currentList []int

		for _, node := range currentNodeList {
			currentList = append(currentList, node.Val)
			if node.LeftChild != nil {
				nextNodeList = append(nextNodeList, node.LeftChild)
			}
			if node.RightChild != nil {
				nextNodeList = append(nextNodeList, node.RightChild)
			}
		}
		res = append(res, currentList)
		currentList = currentList[0:0]

		currentNodeList = nextNodeList
		nextNodeList = nextNodeList[0:0]

	}
	return res
}

//给定一棵二叉搜索树，请找出其中的第k小的TreeNode结点。
//输入：{5,3,7,2,4,6,8},3 返回值：{4}
func KthNode(root *BinaryTreeNode, k int) *BinaryTreeNode {

	_, nodeList := root.InOrderByRecursive()

	if k > len(nodeList) {
		return nil
	}

	return nodeList[k-1]
}
