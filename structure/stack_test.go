package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
