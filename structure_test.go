package main

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
	_, _, err :=head.InsertNode(0, -1)
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
	head , r, err := head.DeleteByIndex(3)
	assert.Equal(t, true, r)
	assert.Equal(t, nil, err)
	assert.Equal(t, testSlice[1], head.Next.Val)
	assert.Equal(t, testSlice[3], head.Next.Next.Val)
	t.Log("dellete  middle success")

	head = GeneralLinkList(testSlice)
	head , result, err := head.DeleteByIndex(6)
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
	var head *Node
	head, v := PopStack(head)
	assert.Equal(t, -999, v)

	head = GeneralStack(testStackSlice)
	head, val := PopStack(head)
	assert.Equal(t, testStackSlice[6], val)

	head, va := PopStack(head)
	assert.Equal(t, testStackSlice[5], va)
}

func TestIsEmpty(t *testing.T) {
	var head *Node
	isEmpty := head.IsEmpty()
	assert.Equal(t, true, isEmpty)

	head = GeneralStack(testStackSlice)
	isEmpty = head.IsEmpty()
	assert.Equal(t, false, isEmpty)
}
