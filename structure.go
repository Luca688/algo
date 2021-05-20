package main

import "errors"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func GeneralNode(val int) *Node {

	return &Node{Val: val}
}

func GeneralLinkList(slice []int) *Node {
	var head *Node
	var currentNode *Node

	for _, val := range slice {
		node := GeneralNode(val)
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
func (head *Node) InsertNode(val int, index int) (headNode *Node, insertRes bool, err error) {

	if index  < 0{
		return nil, false, errors.New("error_index_less_then_0")
	}
	node := GeneralNode(val)
	if index == 0 {
		node.Next = head
		return node, true, nil
	}

	i := 1
	currentNode := head
	for  {
		if i == index {
			break
		}else {
			i++
			if currentNode.Next == nil {
				break
			}else {
				currentNode = currentNode.Next
			}
		}
	}

	node.Next = currentNode.Next
	currentNode.Next = node

	return head, true, nil
}

//delete test
func (head *Node) DeleteByVal(val int) (h *Node, res bool) {
	var prev *Node
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

	if hasVal == false {
		return head, false
	}

	//delete head
	if prev == nil && hasVal == true {
		return currentNode.Next, true
	}
	//删除非头节点
	prev.Next = currentNode.Next
	return head, true

}

//delete test
func (head *Node) DeleteByIndex(index int) (h *Node, res bool, err error) {
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
func (head *Node) FindNodeByVal(val int) bool {
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

func (head *Node) FindNodeByIndex(index int) (node *Node, err error) {
	if index <= 0 {
		return nil, errors.New("index less len 0")
	}
	current := head
	i := 1
	for current != nil {
		if i == index {
			return current, nil
			break
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

func GeneralStack(slice []int) *Node {

	var head *Node

	for _, val := range slice {
		node := GeneralNode(val)
		node.Next = head
		head = node
	}

	return head
}

//push

func PushStack(head *Node, val int) *Node {
	node := GeneralNode(val)
	node.Next = head
	return  node
}

//pop
func PopStack(heaad *Node) (newHead *Node, val int){

	if heaad.IsEmpty() {
		return heaad, -999
	}

	return heaad.Next, heaad.Val
}

//isEmpty

func (head *Node)IsEmpty() bool {
	if head == nil {
		return  true
	}else {
		return  false
	}
}
