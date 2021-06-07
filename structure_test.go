package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSlice = []int{1, 2, 3, 4, 5, 6}

func TestGeneralLinkList(t *testing.T) {
	head := GeneralLinkList(testSlice)
	assert.Equal(t, testSlice[0], head.Val)
	assert.Equal(t, testSlice[1], head.Next.Val)
	assert.Equal(t, testSlice[2], head.Next.Next.Val)
	assert.Equal(t, testSlice[3], head.Next.Next.Next.Val)
	assert.Equal(t, testSlice[4], head.Next.Next.Next.Next.Val)
	assert.Equal(t, testSlice[5], head.Next.Next.Next.Next.Next.Val)

}

func TestInsertNode(t *testing.T) {
	head := GeneralLinkList(testSlice)
	_, _, err := head.InsertNode(0, -1)
	assert.NotEqual(t, err, nil)
	t.Logf(err.Error())

	head, insertRes, err := head.InsertNode(0, 0)
	assert.Equal(t, insertRes, true)
	assert.Equal(t, head.Val, 0)
	assert.Equal(t, head.Next.Val, testSlice[0])
	assert.Equal(t, head.Next.Next.Val, testSlice[1])
	t.Log("insert first success")

	head = GeneralLinkList(testSlice)
	head, insertRes, er := head.InsertNode(10, 1)
	assert.Equal(t, insertRes, true)
	assert.Equal(t, er, nil)
	assert.Equal(t, head.Val, testSlice[0])
	assert.Equal(t, head.Next.Val, 10)
	assert.Equal(t, head.Next.Next.Val, testSlice[1])
	t.Log("insert middle success")

	head = GeneralLinkList(testSlice)
	head, insertRes, e := head.InsertNode(100, 10)
	assert.Equal(t, nil, e)
	assert.Equal(t, true, insertRes)
	assert.Equal(t, 100, head.Next.Next.Next.Next.Next.Next.Val)
	t.Log("insert out success")

	head = GeneralLinkList(testSlice)
	head, insertRes, erro := head.InsertNode(200, 10)
	assert.Equal(t, nil, erro)
	assert.Equal(t, true, insertRes)
	assert.Equal(t, 200, head.Next.Next.Next.Next.Next.Next.Val)
	t.Log("insert last success")

}

func TestDeleteByVal(t *testing.T) {
	head := GeneralLinkList(testSlice)
	newHead, res := head.DeleteByVal(1)
	assert.Equal(t, true, res)
	assert.NotEqual(t, newHead, head)
	assert.Equal(t, testSlice[1], newHead.Val)
	assert.Equal(t, testSlice[2], newHead.Next.Val)
	t.Log("delete first success")

	//1, 2, 4, 5, 6
	head = GeneralLinkList(testSlice)
	newHead, re := head.DeleteByVal(3)
	assert.Equal(t, true, re)
	assert.Equal(t, newHead, head)
	assert.Equal(t, testSlice[1], newHead.Next.Val)
	assert.Equal(t, testSlice[3], newHead.Next.Next.Val)
	t.Log("delete middle success")

	//1,2,3,4,5,6
	head = GeneralLinkList(testSlice)
	newHead, r := head.DeleteByVal(6)
	assert.Equal(t, true, r)
	assert.Equal(t, 5, newHead.Next.Next.Next.Next.Val)
	assert.Equal(t, true, newHead.Next.Next.Next.Next.Next == nil)
	t.Log("delete last success")

}

func TestDeleteByIndex(t *testing.T) {
	head := GeneralLinkList(testSlice)
	_, _, zeroErr := head.DeleteByIndex(0)
	assert.NotEqual(t, nil, zeroErr)
	t.Log(zeroErr.Error())
	t.Log("dellete  0 success")

	newHead, res, _ := head.DeleteByIndex(1)
	assert.Equal(t, true, res)
	assert.NotEqual(t, newHead, head)
	assert.Equal(t, testSlice[1], newHead.Val)
	t.Log("dellete  1 success")

	head = GeneralLinkList(testSlice)
	head, r, err := head.DeleteByIndex(3)
	assert.Equal(t, true, r)
	assert.Equal(t, nil, err)
	assert.Equal(t, testSlice[1], head.Next.Val)
	assert.Equal(t, testSlice[3], head.Next.Next.Val)
	t.Log("dellete  middle success")

	head = GeneralLinkList(testSlice)
	head, result, err := head.DeleteByIndex(6)
	assert.Equal(t, true, result)
	assert.Equal(t, nil, err)
	assert.Equal(t, testSlice[4], head.Next.Next.Next.Next.Val)
	assert.Equal(t, true, head.Next.Next.Next.Next.Next == nil)
	t.Log("del last success")
}

func TestFindNodeByVal(t *testing.T) {
	head := GeneralLinkList(testSlice)

	first := head.FindNodeByVal(1)
	assert.Equal(t, true, first)

	first = head.FindNodeByVal(3)
	assert.Equal(t, true, first)

	first = head.FindNodeByVal(6)
	assert.Equal(t, true, first)

	first = head.FindNodeByVal(100)
	assert.Equal(t, false, first)
}

func TestFindNodeByIndex(t *testing.T) {
	head := GeneralLinkList(testSlice)

	_, zeroErr := head.FindNodeByIndex(0)
	assert.NotEqual(t, nil, zeroErr)
	t.Log(zeroErr.Error())

	node, erro := head.FindNodeByIndex(1)
	assert.Equal(t, nil, erro)
	assert.Equal(t, testSlice[0], node.Val)

	node, er := head.FindNodeByIndex(3)
	assert.Equal(t, nil, er)
	assert.Equal(t, testSlice[2], node.Val)

	node, err := head.FindNodeByIndex(6)
	assert.Equal(t, nil, err)
	assert.Equal(t, testSlice[5], node.Val)

}

/**
stack
*/

var testStackSlice = []int{6, 5, 4, 3, 2, 1, 0}

func TestGeneralStack(t *testing.T) {
	head := GeneralStack(testStackSlice)
	assert.Equal(t, testStackSlice[6], head.Val)
	assert.Equal(t, testStackSlice[5], head.Next.Val)
	assert.Equal(t, testStackSlice[4], head.Next.Next.Val)
}

func TestPushStack(t *testing.T) {
	head := GeneralStack(testStackSlice)
	insert := PushStack(head, -1)
	assert.Equal(t, -1, insert.Val)
	assert.Equal(t, 0, insert.Next.Val)
	t.Log("push-stack ok")
}

func TestPopStack(t *testing.T) {
	var head *StackNode
	head, v := PopStack(head)
	assert.Equal(t, -999, v)

	head = GeneralStack(testStackSlice)
	head, val := PopStack(head)
	assert.Equal(t, testStackSlice[6], val)

	head, va := PopStack(head)
	assert.Equal(t, testStackSlice[5], va)
}

func TestIsEmpty(t *testing.T) {
	var head *StackNode
	isEmpty := head.IsEmpty()
	assert.Equal(t, true, isEmpty)

	head = GeneralStack(testStackSlice)
	isEmpty = head.IsEmpty()
	assert.Equal(t, false, isEmpty)
}

func TestGeneralFullBinaryTree(t *testing.T) {
	/***
			1
		2		3
	   4 5		6 7
	  8
		**/
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	head := GeneralFullBinaryTree(testSlice)
	assert.Equal(t, 1, head.Val)
	assert.Equal(t, 2, head.LeftChild.Val)
	assert.Equal(t, 4, head.LeftChild.LeftChild.Val)
	assert.Equal(t, 8, head.LeftChild.LeftChild.LeftChild.Val)

	assert.Equal(t, 3, head.RightChild.Val)
	assert.Equal(t, 7, head.RightChild.RightChild.Val)
}

func TestLevelOrder(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := GeneralFullBinaryTree(testSlice)
	res := root.LevelOrder()
	assert.Equal(t, testSlice[0], res[0])
	assert.Equal(t, testSlice[1], res[1])
	assert.Equal(t, testSlice[2], res[2])
	assert.Equal(t, testSlice[3], res[3])
	assert.Equal(t, testSlice[4], res[4])
	assert.Equal(t, testSlice[5], res[5])
	assert.Equal(t, testSlice[6], res[6])
	assert.Equal(t, testSlice[7], res[7])
}
func TestPrevOrderByIteration(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := GeneralFullBinaryTree(testSlice)
	res := root.PrevOrderByIteration()
	//1 2 4 8 5 3 6 7
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 2, res[1])
	assert.Equal(t, 4, res[2])
	assert.Equal(t, 8, res[3])
	assert.Equal(t, 5, res[4])
	assert.Equal(t, 3, res[5])
	assert.Equal(t, 6, res[6])
	assert.Equal(t, 7, res[7])
}

func TestPrevOrderByRecursive(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := GeneralFullBinaryTree(testSlice)
	res := root.PrevOrderByRecursive()
	//1 2 4 8 5 3 6 7
	assert.Equal(t, 1, res[0])
	// t.Log(res)
	assert.Equal(t, 2, res[1])
	assert.Equal(t, 4, res[2])
	assert.Equal(t, 8, res[3])
	assert.Equal(t, 5, res[4])
	assert.Equal(t, 3, res[5])
	assert.Equal(t, 6, res[6])
	assert.Equal(t, 7, res[7])
}

// func TestSequenceOrderByIteration(t *testing.T) {
// 	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	root := GeneralFullBinaryTree(testSlice)
// 	res := root.SequenceOrderByIteration()
// 	//1 2 4 8 5 3 6 7
// 	assert.Equal(t, 8, res[0])
// 	// t.Log(res)
// 	assert.Equal(t, 4, res[1])
// 	assert.Equal(t, 2, res[2])
// 	assert.Equal(t, 5, res[3])
// 	assert.Equal(t, 1, res[4])
// 	assert.Equal(t, 6, res[5])
// 	assert.Equal(t, 3, res[6])
// 	assert.Equal(t, 7, res[7])
// }

func TestNewBSTree(t *testing.T) {
	/**
			15
		5        17
	  3   12    16  20
	1   4
		**/
	testSlice := []int{15, 5, 17, 3, 12, 16, 20, 1, 4}
	root := NewBSTree(testSlice)

	assert.Equal(t, 15, root.Val)

	assert.Equal(t, 5, root.LeftChild.Val)
	assert.Equal(t, 17, root.RightChild.Val)

	assert.Equal(t, 3, root.LeftChild.LeftChild.Val)
	assert.Equal(t, 12, root.LeftChild.RightChild.Val)
	assert.Equal(t, 16, root.RightChild.LeftChild.Val)
	assert.Equal(t, 20, root.RightChild.RightChild.Val)

	assert.Equal(t, 1, root.LeftChild.LeftChild.LeftChild.Val)
	assert.Equal(t, 4, root.LeftChild.LeftChild.RightChild.Val)
}

func TestSearchBSTreeVal(t *testing.T) {
	/**
			15
		5        17
	  3   12    16  20
	1   4
		**/
	testSlice := []int{15, 5, 17, 3, 12, 16, 20, 1, 4}
	root := NewBSTree(testSlice)

	hasVal, _, _ := root.SearchBSTreeVal(15)
	assert.Equal(t, true, hasVal)

	hasVal, _, _ = root.SearchBSTreeVal(4)
	assert.Equal(t, true, hasVal)

	hasVal, _, _ = root.SearchBSTreeVal(3)
	assert.Equal(t, true, hasVal)

	hasVal, _, _ = root.SearchBSTreeVal(17)
	assert.Equal(t, true, hasVal)

	hasVal, _, _ = root.SearchBSTreeVal(12)
	assert.Equal(t, true, hasVal)

	hasVal, _, _ = root.SearchBSTreeVal(-999)
	assert.Equal(t, false, hasVal)
}

func TestSearchMaxBSTreeVal(t *testing.T) {
	/**
			15
		5        17
	  3   12    16  20
	1   4
		**/
	testSlice := []int{15, 5, 17, 3, 12, 16, 20, 1, 4}
	root := NewBSTree(testSlice)

	node, _ := root.SearchMaxBSTreeVal()
	assert.Equal(t, 20, node.Val)
}

func TestDeleteBSTreeVal(t *testing.T) {

	//叶子节点
	testSlice := []int{15, 5, 3, 1}
	root := NewBSTree(testSlice)
	res, newRoot := root.DeleteBSTreeVal(1)
	assert.Equal(t, true, res)
	assert.Equal(t, true, newRoot == root)

	//无指定值得节点
	testSlice = []int{15, 5, 3, 1}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(666)
	assert.Equal(t, false, res)
	assert.Equal(t, true, newRoot == root)

	//只有一个节点的树
	testSlice = []int{15}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(15)
	assert.Equal(t, true, res)
	assert.Equal(t, true, newRoot == nil)

	//叶子节点
	testSlice = []int{15, 5, 17}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(5)
	assert.Equal(t, true, res)
	assert.Equal(t, true, newRoot.LeftChild == nil)
	assert.Equal(t, true, newRoot == root)
	res, newRoot = root.DeleteBSTreeVal(17)
	assert.Equal(t, true, res)
	assert.Equal(t, true, newRoot.RightChild == nil)

	//单分支节点

	//左斜树根节点
	testSlice = []int{15, 5, 3, 1}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(15)
	assert.Equal(t, true, res)
	assert.Equal(t, true, root.LeftChild == newRoot)
	assert.Equal(t, true, newRoot != root)

	//左斜树非根节点
	testSlice = []int{15, 5, 3, 1}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(3)
	assert.Equal(t, true, res)
	assert.Equal(t, true, newRoot == root)
	assert.Equal(t, 1, newRoot.LeftChild.LeftChild.Val)

	//右斜树根节点
	testSlice = []int{15, 17, 20}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(15)
	assert.Equal(t, true, res)
	assert.Equal(t, true, root.RightChild == newRoot)
	assert.Equal(t, true, newRoot != root)

	//右斜树非根节点
	testSlice = []int{15, 17, 20}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(17)
	assert.Equal(t, true, res)
	assert.Equal(t, true, newRoot == root)
	assert.Equal(t, 20, newRoot.RightChild.Val)

	testSlice = []int{15, 17, 20, 18}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(20)
	assert.Equal(t, true, res)
	assert.Equal(t, true, newRoot == root)
	assert.Equal(t, 18, newRoot.RightChild.RightChild.Val)

	//双节点
	/**
			15
		5        17
	  3   12    16  20
	1   4
	  2
		**/

	//非根节点左子树
	testSlice = []int{17, 14, 13, 15, 16}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(14)
	assert.Equal(t, true, res)
	assert.Equal(t, 17, newRoot.Val)
	assert.Equal(t, 15, newRoot.LeftChild.Val)
	assert.Equal(t, 13, newRoot.LeftChild.LeftChild.Val)
	assert.Equal(t, 16, newRoot.LeftChild.RightChild.Val)

	//非根节点又子树
	testSlice = []int{15, 5, 17, 3, 12, 16, 20, 1, 4, 2}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(17)
	assert.Equal(t, true, res)
	assert.Equal(t, 15, newRoot.Val)
	assert.Equal(t, 20, newRoot.RightChild.Val)
	assert.Equal(t, 16, newRoot.RightChild.LeftChild.Val)
	assert.Equal(t, 5, newRoot.LeftChild.Val)
	assert.Equal(t, 3, newRoot.LeftChild.LeftChild.Val)
	assert.Equal(t, 12, newRoot.LeftChild.RightChild.Val)

	//根节点
	testSlice = []int{15, 5, 17, 3, 12, 16, 20, 1, 4, 2}
	root = NewBSTree(testSlice)
	res, newRoot = root.DeleteBSTreeVal(15)
	assert.Equal(t, true, res)
	assert.Equal(t, 16, newRoot.Val)
	assert.Equal(t, 17, newRoot.RightChild.Val)
	assert.Equal(t, 20, newRoot.RightChild.RightChild.Val)
	assert.Equal(t, 5, newRoot.LeftChild.Val)
	assert.Equal(t, 3, newRoot.LeftChild.LeftChild.Val)
	assert.Equal(t, 12, newRoot.LeftChild.RightChild.Val)

}

func TestInOrderByRecursive(t *testing.T) {
	/**
			15
		5        17
	  3   12    16  20
	1   4
	  2
		**/

	//非根节点左子树
	testSlice := []int{15, 5, 17, 3, 12, 16, 20, 1, 4, 2}
	root := NewBSTree(testSlice)
	res, _ := root.InOrderByRecursive()
	resSlice := []int{1, 2, 3, 4, 5, 12, 15, 16, 17, 20}
	assert.Equal(t, resSlice, res)
}

//递归前序遍历-根左右
func TestPreOrderByRecursive(t *testing.T) {
	/**
			15
		5        17
	  3   12    16  20
	1   4
	  2
		**/

	//非根节点左子树
	testSlice := []int{15, 5, 17, 3, 12, 16, 20, 1, 4, 2}
	root := NewBSTree(testSlice)
	res, _ := root.PreOrderByRecursive()
	resSlice := []int{15, 5, 3, 1, 2, 4, 12, 17, 16, 20}
	assert.Equal(t, resSlice, res)
}

func TestPrevOrder(t *testing.T) {
	testSlice := []int{5, 3, 7, 1, 4, 6, 8, 0, 2, 9}
	root := NewBSTree(testSlice)
	res, _ := root.PrevOrder()
	resSlice := []int{5, 3, 1, 0, 2, 4, 7, 6, 8, 9}
	assert.Equal(t, resSlice, res)
}
