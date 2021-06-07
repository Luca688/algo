package algo

//https://www.cnblogs.com/flmei/p/10768617.html

import (
	"errors"
	"fmt"
)

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

/***
tree
**/

type BinaryTreeNode struct {
	LeftChild  *BinaryTreeNode
	RightChild *BinaryTreeNode
	Val        int
}

func NewBinaryTreeNode(val int) *BinaryTreeNode {
	return &BinaryTreeNode{Val: val}
}

func GeneralFullBinaryTree(slice []int) *BinaryTreeNode {

	var root *BinaryTreeNode

	if len(slice) <= 0 {
		return nil
	}

	root = NewBinaryTreeNode(slice[0])
	var currentNodeList []*BinaryTreeNode
	// var nextNodeList []*BinaryTreeNode
	currentNodeList = append(currentNodeList, root)

	sliceIndex := 1

	for len(currentNodeList) > 0 {
		var tmpNextList []*BinaryTreeNode
		for _, currentNode := range currentNodeList {
			currentNode.LeftChild = NewBinaryTreeNode(slice[sliceIndex])
			tmpNextList = append(tmpNextList, currentNode.LeftChild)
			sliceIndex++
			if sliceIndex >= len(slice) {
				return root
			}
			currentNode.RightChild = NewBinaryTreeNode(slice[sliceIndex])
			tmpNextList = append(tmpNextList, currentNode.RightChild)
			sliceIndex++
			if sliceIndex >= len(slice) {
				return root
			}
		}
		currentNodeList = tmpNextList
	}

	return root
}

func (root *BinaryTreeNode) LevelOrder() []int {
	var res []int
	var currentNodeList []*BinaryTreeNode

	currentNodeList = append(currentNodeList, root)

	for len(currentNodeList) > 0 {
		var nextNodeList []*BinaryTreeNode
		for _, currentNode := range currentNodeList {
			if currentNode != nil {
				res = append(res, currentNode.Val)
				if currentNode.LeftChild != nil {
					nextNodeList = append(nextNodeList, currentNode.LeftChild)
				}
				if currentNode.RightChild != nil {
					nextNodeList = append(nextNodeList, currentNode.RightChild)
				}
			}
		}
		currentNodeList = nextNodeList
	}
	return res
}

//根左右
func (root *BinaryTreeNode) PrevOrderByIteration() []int {
	var res []int
	var currentNodeList []*BinaryTreeNode
	currentNode := root
	for {
		if currentNode == nil {
			break
		} else {
			res = append(res, currentNode.Val)

			var newNodes []*BinaryTreeNode

			if currentNode.LeftChild != nil {
				newNodes = append(newNodes, currentNode.LeftChild)
			}
			if currentNode.RightChild != nil {
				newNodes = append(newNodes, currentNode.RightChild)
			}

			if currentNodeList == nil {
				currentNodeList = newNodes
			} else {
				currentNodeList = append(newNodes, currentNodeList[1:]...)
			}

			if len(currentNodeList) > 0 {
				currentNode = currentNodeList[0]
			} else {
				currentNode = nil
			}
		}
	}
	return res
}

func (root *BinaryTreeNode) PrevOrderByRecursive() []int {
	var res []int
	// var res = make([]int, 10)
	PrevOrderByRecursiveSub(root, &res)
	return res
}
func PrevOrderByRecursiveSub(node *BinaryTreeNode, res *[]int) {

	if node != nil {

		*res = append(*res, node.Val)
		fmt.Println(res)
		if node.LeftChild != nil {
			PrevOrderByRecursiveSub(node.LeftChild, res)
		}
		if node.RightChild != nil {
			PrevOrderByRecursiveSub(node.RightChild, res)
		}
	} else {
		return
	}
}

//左根右
// func (root *BinaryTreeNode) SequenceOrderByIteration() []int {
// 	var res []int
// 	var nodeList = make([]*BinaryTreeNode, 10)
// 	nodeList[0] = root
// 	var lastPopNode *BinaryTreeNode
// 	for {
// 		fmt.Println(res)
// 		if len(nodeList) <= 0 {
// 			break
// 		}
// 		var currentNodeList []*BinaryTreeNode
// 		currentNode := nodeList[0]
// 		nodeList = nodeList[1:]
// 		fmt.Println("nodeList")
// 		fmt.Println(nodeList)
// 		//叶子节点
// 		if currentNode.LeftChild == nil && currentNode.RightChild == nil {
// 			res = append(res, currentNode.Val)
// 			lastPopNode = currentNode
// 			continue
// 		}
// 		//左节点已出或没有左节点-时出当前节点
// 		if (currentNode.LeftChild != nil && currentNode.LeftChild == lastPopNode) || currentNode.LeftChild == nil {
// 			res = append(res, currentNode.Val)
// 			lastPopNode = currentNode
// 			continue
// 		}
// 		//根节点出
// 		if currentNode == root {
// 			res = append(res, currentNode.Val)
// 			lastPopNode = currentNode
// 			continue
// 		}

// 		if currentNode.LeftChild != nil {
// 			currentNodeList = append(currentNodeList, currentNode.LeftChild)
// 		}
// 		currentNodeList = append(currentNodeList, currentNode)
// 		if currentNode.RightChild != nil {
// 			currentNodeList = append(currentNodeList, currentNode.RightChild)
// 		}
// 		nodeList = append(currentNodeList, nodeList...)

// 	}
// 	return res
// }

func NewBSTree(data []int) *BinaryTreeNode {

	if len(data) <= 0 {
		return nil
	}

	root := NewBinaryTreeNode(data[0])
	for i := 1; i < len(data); i++ {
		root.InsertBSTreeNode(data[i])
	}

	return root
}

//1.左子树上的所有节点值均小于根节点值，2右子树上的所有节点值均不小于根节点值，3，左右子树也满足上述两个条件。
func (root *BinaryTreeNode) InsertBSTreeNode(val int) bool {

	hasVal, _, _ := root.SearchBSTreeVal(val)
	if hasVal {
		return true
	}

	currentNode := root
	targetNode := root
	for currentNode != nil {
		targetNode = currentNode
		if val > currentNode.Val {
			currentNode = currentNode.RightChild
		} else {
			currentNode = currentNode.LeftChild
		}
	}

	insertNode := NewBinaryTreeNode(val)
	if targetNode.Val > val {
		targetNode.LeftChild = insertNode
	} else {
		targetNode.RightChild = insertNode
	}

	return true
}

func (root *BinaryTreeNode) DeleteBSTreeVal(val int) (res bool, newRoot *BinaryTreeNode) {
	newRoot = root

	//无指定值得节点
	hasVal, node, parentNode := root.SearchBSTreeVal(val)
	if !hasVal {
		return false, root
	}

	//叶子节点
	if node.LeftChild == nil && node.RightChild == nil {

		//只有一个节点的树
		if node == root {
			return true, nil

		} else {
			if parentNode.LeftChild == node {
				parentNode.LeftChild = nil
			} else {
				parentNode.RightChild = nil
			}
			return true, newRoot
		}
	}

	//单分支节点
	if node.LeftChild == nil || node.RightChild == nil {

		//根节点
		if node == root {
			if node.LeftChild == nil {
				return true, node.RightChild
			} else {
				return true, node.LeftChild
			}
		}

		//左子树位空的非根节点
		if node.LeftChild == nil {

			//被删除元素位于左子树
			if parentNode.LeftChild == node {
				parentNode.LeftChild = node.RightChild
			} else {
				parentNode.RightChild = node.RightChild
			}

			return true, newRoot
		}
		//右子树为空的非根节点
		if node.RightChild == nil {

			//被删除的元素位于左子树
			if parentNode.LeftChild == node {

				parentNode.LeftChild = node.LeftChild

			} else {
				parentNode.RightChild = node.LeftChild
			}

			return true, newRoot
		}
	}

	//双分支节点
	if node.LeftChild != nil && node.RightChild != nil {

		//右子树的最小节点一定没有左节点
		minNode, pNode := node.RightChild.SearchMinBSTreeVal()

		//根节点
		if node == root {
			newRoot = minNode
			newRoot.LeftChild = root.LeftChild
			pNode.LeftChild = newRoot.RightChild
			newRoot.RightChild = root.RightChild
			return true, newRoot
		}

		//非根节点
		if node != root {

			minNode.LeftChild = node.LeftChild
			if pNode != nil {

				pNode.LeftChild = minNode.RightChild
			}

			if parentNode.RightChild == node {

				parentNode.RightChild = minNode
			}

			if parentNode.LeftChild == node {

				parentNode.LeftChild = minNode
			}
			return true, newRoot
		}

	}

	return true, newRoot
}

func (root *BinaryTreeNode) SearchMaxBSTreeVal() (node *BinaryTreeNode, parent *BinaryTreeNode) {

	var pNode *BinaryTreeNode

	currentNode := root

	for {

		if currentNode.RightChild == nil {
			return currentNode, pNode
		}

		pNode = currentNode
		currentNode = currentNode.RightChild

	}
}

func (root *BinaryTreeNode) SearchMinBSTreeVal() (node *BinaryTreeNode, parent *BinaryTreeNode) {

	var pNode *BinaryTreeNode

	currentNode := root

	for {

		if currentNode.LeftChild == nil {
			return currentNode, pNode
		}

		pNode = currentNode
		currentNode = currentNode.LeftChild

	}
}

func (root *BinaryTreeNode) SearchBSTreeVal(val int) (res bool, node *BinaryTreeNode, parentNode *BinaryTreeNode) {

	var pNode *BinaryTreeNode
	currentNode := root
	for currentNode != nil {

		if currentNode.Val == val {
			return true, currentNode, pNode
		} else {
			pNode = currentNode
			if val > currentNode.Val {
				currentNode = currentNode.RightChild
			} else {
				currentNode = currentNode.LeftChild
			}
		}
	}
	return false, nil, nil
}

//递归中序遍历-左根右
func (root *BinaryTreeNode) InOrderByRecursive() (slice []int, nodeList []*BinaryTreeNode) {
	var res []int
	var resNodeList []*BinaryTreeNode
	root.InOrderByRecursiveSub(root, &res, &resNodeList)
	return res, resNodeList
}

func (root *BinaryTreeNode) InOrderByRecursiveSub(node *BinaryTreeNode, res *[]int, resNodeList *[]*BinaryTreeNode) {
	if node.LeftChild != nil {
		root.InOrderByRecursiveSub(node.LeftChild, res, resNodeList)
	}
	*res = append(*res, node.Val)
	*resNodeList = append(*resNodeList, node)
	if node.RightChild != nil {
		root.InOrderByRecursiveSub(node.RightChild, res, resNodeList)
	}
}

//递归前序遍历-根左右
func (root *BinaryTreeNode) PreOrderByRecursive() (slice []int, nodeList []*BinaryTreeNode) {
	var res []int
	var resNodeList []*BinaryTreeNode
	root.PreOrderByRecursiveSub(root, &res, &resNodeList)
	return res, resNodeList
}

func (root *BinaryTreeNode) PreOrderByRecursiveSub(node *BinaryTreeNode, res *[]int, resNodeList *[]*BinaryTreeNode) {

	*res = append(*res, node.Val)
	*resNodeList = append(*resNodeList, node)

	if node.LeftChild != nil {
		root.PreOrderByRecursiveSub(node.LeftChild, res, resNodeList)
	}
	if node.RightChild != nil {
		root.PreOrderByRecursiveSub(node.RightChild, res, resNodeList)
	}
}

//前序遍历-中左右
func (root *BinaryTreeNode) PrevOrder() (resList []int, resNodeList []*BinaryTreeNode) {
	var list []int
	var nodeList []*BinaryTreeNode

	if root == nil {
		return list, nodeList
	}

	currentNode := root
	var nextNodeList []*BinaryTreeNode

	for currentNode != nil {
		list = append(list, currentNode.Val)
		nodeList = append(nodeList, currentNode)
		if currentNode.RightChild != nil {
			nextNodeList = append([]*BinaryTreeNode{currentNode.RightChild}, nextNodeList...)
		}
		if currentNode.LeftChild != nil {
			nextNodeList = append([]*BinaryTreeNode{currentNode.LeftChild}, nextNodeList...)
		}
		if len(nextNodeList) > 0 {
			currentNode = nextNodeList[0]
			nextNodeList = nextNodeList[1:]
		} else {
			break
		}
	}

	return list, nodeList
}
