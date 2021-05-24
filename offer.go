package main

import (
	"errors"
)

//删除链表中重复的节点
//在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
// 例如，链表1->1->1->1->1->1->2->3->3->3->3->4->4->5 处理后为 1->2->5
func SortLinkListRepeat(head *Node) *Node {
	if head == nil {
		return head
	}

	currentNode := head
	var lastNode *Node

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

func IncreasingLinkMerge(head1 *Node, head2 *Node) *Node {

	var res *Node
	var resTail *Node
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

func ReciprocalKNode(head *Node, k int) (res int, err error) {

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
func EntryNodeOfLoop(head *Node) *Node {

	//1.遍历一遍有重复的就是环
	allNode := make(map[*Node]*Node)
	currentNode := head
	var res *Node
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
func FindFirstCommonNode(left *Node, right *Node) *Node {

	var res *Node
	allNode := make(map[*Node]*Node)

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
