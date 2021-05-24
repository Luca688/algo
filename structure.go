package main

import "errors"

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func GeneralListNode(val int) *ListNode {

	return &ListNode{Val: val}
}

func GeneralLinkList(slice []int) *ListNode {
	var head *ListNode
	var currentNode *ListNode

	for _, val := range slice {
		node := GeneralListNode(val)
		if head == nil {
			head = node
		} else {
			currentNode.Next = node
		}
		currentNode = node
	}

	return head
}

//insert test ok
func (head *ListNode) InsertNode(val int, index int) (headNode *ListNode, insertRes bool, err error) {

	if index < 0 {
		return nil, false, errors.New("error_index_less_then_0")
	}
	node := GeneralListNode(val)
	if index == 0 {
		node.Next = head
		return node, true, nil
	}

	i := 1
	currentNode := head
	for {
		if i == index {
			break
		} else {
			i++
			if currentNode.Next == nil {
				break
			} else {
				currentNode = currentNode.Next
			}
		}
	}

	node.Next = currentNode.Next
	currentNode.Next = node

	return head, true, nil
}

//delete test
func (head *ListNode) DeleteByVal(val int) (h *ListNode, res bool) {
	var prev *ListNode
	currentNode := head
	hasVal := false

	for currentNode != nil {
		if currentNode.Val == val {
			hasVal = true
			break
		} else {
			prev = currentNode
			currentNode = currentNode.Next
		}

	}

	if !hasVal {
		return head, false
	}

	//delete head
	if prev == nil && hasVal {
		return currentNode.Next, true
	}
	//删除非头节点
	prev.Next = currentNode.Next
	return head, true

}

//delete test
func (head *ListNode) DeleteByIndex(index int) (h *ListNode, res bool, err error) {
	if index < 1 {
		return head, false, errors.New("delete index less than 0")
	}
	if index == 1 {
		if head.Next == nil {
			return head, false, errors.New("node length equal 1")
		} else {
			return head.Next, true, nil
		}
	}
	//删除非头节点
	currentNode := head.Next
	prev := head
	i := 2
	for currentNode.Next != nil {
		if i == index {
			break
		} else {
			i++
			prev = currentNode
			currentNode = currentNode.Next
		}
	}
	if i == index {
		prev.Next = currentNode.Next
		return head, true, nil
	} else {
		return head, false, nil
	}
}

//find test
func (head *ListNode) FindNodeByVal(val int) bool {
	currentNode := head

	for currentNode != nil {
		if currentNode.Val == val {
			return true
		} else {
			currentNode = currentNode.Next
		}
	}
	return false
}

func (head *ListNode) FindNodeByIndex(index int) (node *ListNode, err error) {
	if index <= 0 {
		return nil, errors.New("index less len 0")
	}
	current := head
	i := 1
	for current != nil {
		if i == index {
			return current, nil
		} else {
			i++
			current = current.Next
		}
	}
	return nil, nil
}



/**
stack
*/

type StackNode struct {
	Val  int
	Next *StackNode
	Prev *StackNode
}

func GeneralStackNode(val int) *StackNode {

	return &StackNode{Val: val}
}

func GeneralStack(slice []int) *StackNode {

	var head *StackNode

	for _, val := range slice {
		node := GeneralStackNode(val)
		node.Next = head
		head = node
	}

	return head
}

//push

func PushStack(head *StackNode, val int) *StackNode {
	node := GeneralStackNode(val)
	node.Next = head
	return node
}

//pop
func PopStack(heaad *StackNode) (newHead *StackNode, val int) {

	if heaad.IsEmpty() {
		return heaad, -999
	}

	return heaad.Next, heaad.Val
}

//isEmpty

func (head *StackNode) IsEmpty() bool {
	if head == nil {
		return true
	} else {
		return false
	}
}
