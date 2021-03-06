package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPath(t *testing.T) {

	root := NewBinaryTreeNode(10)
	root.RightChild = NewBinaryTreeNode(12)
	root.LeftChild = NewBinaryTreeNode(5)
	root.LeftChild.LeftChild = NewBinaryTreeNode(4)
	root.LeftChild.RightChild = NewBinaryTreeNode(7)
	res := FindPath(root, 22)
	t.Log(res)

}

func TestVerifySquenceOfBST(t *testing.T) {

	testSlice := []int{4, 8, 6, 12, 16, 14, 10}
	assert.Equal(t, true, VerifySquenceOfBST(testSlice))

	testSlice = []int{4, 8, 6, 12, 16, 14, 15}
	assert.Equal(t, false, VerifySquenceOfBST(testSlice))

	testSlice = []int{4, 8, 6, 2, 16, 14, 10}
	assert.Equal(t, false, VerifySquenceOfBST(testSlice))

	testSlice = []int{4, 8, 6, 12, 16, 9, 10}
	assert.Equal(t, false, VerifySquenceOfBST(testSlice))
}

func TestIsPopOrder(t *testing.T) {

	pushV := []int{1, 2, 3, 4, 5}
	popV := []int{4, 5, 3, 2, 1}
	res := IsPopOrder(pushV, popV)
	assert.Equal(t, true, res)

	popV = []int{4, 3, 5, 1, 2}
	res = IsPopOrder(pushV, popV)
	assert.Equal(t, false, res)

	popV = []int{5, 4, 3, 2, 1}
	res = IsPopOrder(pushV, popV)
	assert.Equal(t, true, res)
}

func TestPrintFromTopToBottom(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestPrintMatrix(t *testing.T) {

	//剩余1个元素
	testSlice := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 2, 2, 2, 2, 2, 1},
		{1, 2, 3, 3, 3, 2, 1},
		{1, 2, 3, 4, 3, 2, 1},
		{1, 2, 3, 3, 3, 2, 1},
		{1, 2, 2, 2, 2, 2, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	res := PrintMatrix(testSlice)
	t.Log(res)
	assert.Equal(t, true, true)

	//剩余0
	testSlice = [][]int{
		{1, 1},
		{1, 1},
	}
	res = PrintMatrix(testSlice)
	assert.Equal(t, []int{1, 1, 1, 1}, res)

	testSlice = [][]int{
		{1, 1, 1, 1},
		{1, 2, 2, 1},
		{1, 2, 2, 1},
		{1, 1, 1, 1},
	}
	res = PrintMatrix(testSlice)
	assert.Equal(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2}, res)

	//剩余一行
	testSlice = [][]int{
		{1, 1, 1, 1},
		{1, 2, 2, 1},
		{1, 1, 1, 1},
	}
	res = PrintMatrix(testSlice)
	assert.Equal(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2}, res)

	//剩余一列
	testSlice = [][]int{
		{1, 2, 3},
		{10, 11, 4},
		{9, 12, 5},
		{8, 7, 6},
	}
	res = PrintMatrix(testSlice)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, res)
}

func TestMirror(t *testing.T) {

	testSlice := []int{8, 6, 10, 5, 7, 9, 11}
	root := NewBSTree(testSlice)
	mirror := Mirror(root)
	level := mirror.LevelOrder()
	expected := []int{8, 10, 6, 11, 9, 7, 5}
	assert.Equal(t, expected, level)

}

func TestHasSubtree(t *testing.T) {

	pSlice := []int{5, 3, 7, 1, 4, 6, 8, 0, 2, 9}
	pRoot := NewBSTree(pSlice)
	cSlice := []int{5, 3, 7, 1, 4, 6, 8, 0, 2, 9}
	cRoot := NewBSTree(cSlice)
	assert.Equal(t, true, HasSubtree(pRoot, cRoot))

	pSlice = []int{5, 3, 7, 1, 4, 6, 8, 0, 2, 9}
	pRoot = NewBSTree(pSlice)
	cSlice = []int{5}
	cRoot = NewBSTree(cSlice)
	assert.Equal(t, false, HasSubtree(pRoot, cRoot))

	pSlice = []int{5, 3, 7, 1, 4, 6, 8, 0, 2, 9}
	pRoot = NewBSTree(pSlice)
	cSlice = []int{3, 1, 4, 2}
	cRoot = NewBSTree(cSlice)
	assert.Equal(t, false, HasSubtree(pRoot, cRoot))

	pSlice = []int{5, 3, 7, 1, 4, 6, 8, 0, 2, 9}
	pRoot = NewBSTree(pSlice)
	cSlice = []int{3, 1, 4, 0, 2}
	cRoot = NewBSTree(cSlice)
	assert.Equal(t, false, HasSubtree(pRoot, cRoot))

}

func TestReOrderArray(t *testing.T) {

	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := ReOrderArray(testSlice)
	expectRes := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	assert.Equal(t, expectRes, res)
}

func TestPowr(t *testing.T) {

	res := Power(2.0, 3)
	assert.Equal(t, 8.0, res)

	res = Power(2.1, 3)
	assert.Equal(t, 9.261000000000001, res)

	res = Power(2.0, -2)
	assert.Equal(t, 0.250, res)
}

func TestJumpFloorII(t *testing.T) {
	res := jumpFloorII(1)
	assert.Equal(t, 1, res)

	res = jumpFloorII(4)
	assert.Equal(t, 8, res)
}

func TestJumpFloor(t *testing.T) {
	res := JumpFloor(1)
	assert.Equal(t, 1, res)

	res = JumpFloor(2)
	assert.Equal(t, 2, res)

	res = JumpFloor(3)
	assert.Equal(t, 3, res)

	res = JumpFloor(6)
	assert.Equal(t, 13, res)
}

func TestFibonacci(t *testing.T) {
	res := Fibonacci(1)
	assert.Equal(t, 0, res)

	res = Fibonacci(2)
	assert.Equal(t, 1, res)

	res = Fibonacci(3)
	assert.Equal(t, 1, res)

	res = Fibonacci(8)
	assert.Equal(t, 13, res)
}

func TestPop(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	Push(testSlice[0])
	Push(testSlice[1])
	Push(testSlice[2])
	Push(testSlice[3])
	Push(testSlice[4])

	assert.Equal(t, 1, Pop())
	assert.Equal(t, 2, Pop())
	assert.Equal(t, 3, Pop())
	assert.Equal(t, 4, Pop())
	assert.Equal(t, 5, Pop())

}

func TestReConstructBinaryTree(t *testing.T) {

	// pre := []int{1, 2, 3, 4, 5, 6, 7}
	// vin := []int{3, 2, 4, 1, 6, 5, 7}
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	vin := []int{4, 7, 2, 1, 5, 3, 8, 6}
	res := ReConstructBinaryTree(pre, vin)
	assert.Equal(t, 1, res.Val)
}

//从尾到头打印链表
func TestPrintListFromTailToHead(t *testing.T) {
	testSlice := []int{67, 0, 24, 58}
	root := GeneralLinkList(testSlice)
	res := PrintListFromTailToHead(root)
	expectedSlice := []int{58, 24, 0, 67}
	assert.Equal(t, expectedSlice, res)
}

//替换空格
func TestReplaceSpace(t *testing.T) {
	test := "We Are Happy"

	res := ReplaceSpace(test)
	assert.Equal(t, "We%20Are%20Happy", res)
}

func TestFind(t *testing.T) {
	testSlice := [][]int{
		{1, 2, 8, 9},
	}

	res := Find(1, testSlice)
	assert.Equal(t, true, res)
	res = Find(8, testSlice)
	assert.Equal(t, true, res)
	res = Find(10, testSlice)
	assert.Equal(t, false, res)

	testSlice = [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	res = Find(1, testSlice)
	assert.Equal(t, true, res)

	res = Find(0, testSlice)
	assert.Equal(t, false, res)

	res = Find(17, testSlice)
	assert.Equal(t, false, res)

	res = Find(10, testSlice)
	assert.Equal(t, true, res)
	res = Find(8, testSlice)
	assert.Equal(t, true, res)
	res = Find(6, testSlice)
	assert.Equal(t, true, res)

}

func TestSortLinkListRepeat(t *testing.T) {

	testHeadRepeat := []int{1, 1, 1, 1, 1, 2, 3, 3, 3, 4, 4, 5, 6}

	repeatLink := GeneralLinkList(testHeadRepeat)

	newRepeatLink := SortLinkListRepeat(repeatLink)
	assert.NotEqual(t, newRepeatLink, repeatLink)
	assert.Equal(t, 2, newRepeatLink.Val)
	assert.Equal(t, 5, newRepeatLink.Next.Val)
	assert.Equal(t, 6, newRepeatLink.Next.Next.Val)

	testHeadRepeat = []int{1, 1, 1, 1, 1, 2, 3, 3, 3, 4, 4, 5, 6, 6}

	repeatLink = GeneralLinkList(testHeadRepeat)
	newRepeatLink = SortLinkListRepeat(repeatLink)
	assert.NotEqual(t, newRepeatLink, repeatLink)
	assert.Equal(t, 2, newRepeatLink.Val)
	assert.Equal(t, 5, newRepeatLink.Next.Val)
	assert.Equal(t, true, newRepeatLink.Next.Next == nil)

	testHeadRepeat = []int{1, 2, 3}

	repeatLink = GeneralLinkList(testHeadRepeat)
	newRepeatLink = SortLinkListRepeat(repeatLink)
	assert.Equal(t, 1, newRepeatLink.Val)
	assert.Equal(t, 2, newRepeatLink.Next.Val)
	assert.Equal(t, 3, newRepeatLink.Next.Next.Val)
}

func TestIncreasingLinkMerge(t *testing.T) {

	left := []int{1, 3, 5}
	right := []int{2, 4, 6}
	leftLink := GeneralLinkList(left)
	rightLink := GeneralLinkList(right)
	afterMerge := IncreasingLinkMerge(leftLink, rightLink)

	assert.Equal(t, 1, afterMerge.Val)
	assert.Equal(t, 2, afterMerge.Next.Val)
	assert.Equal(t, 3, afterMerge.Next.Next.Val)
	assert.Equal(t, 4, afterMerge.Next.Next.Next.Val)
	assert.Equal(t, 5, afterMerge.Next.Next.Next.Next.Val)
	assert.Equal(t, 6, afterMerge.Next.Next.Next.Next.Next.Val)

	left = []int{2, 11, 77, 79, 80}
	right = []int{3, 4, 5, 82, 83, 84}
	leftLink = GeneralLinkList(left)
	rightLink = GeneralLinkList(right)
	afterMerge = IncreasingLinkMerge(leftLink, rightLink)
	//{2,3,4,5,11,77,79,80,83,84}
	assert.Equal(t, 2, afterMerge.Val)
	assert.Equal(t, 3, afterMerge.Next.Val)
	assert.Equal(t, 4, afterMerge.Next.Next.Val)
	assert.Equal(t, 5, afterMerge.Next.Next.Next.Val)
	assert.Equal(t, 11, afterMerge.Next.Next.Next.Next.Val)
	assert.Equal(t, 77, afterMerge.Next.Next.Next.Next.Next.Val)
	assert.Equal(t, 79, afterMerge.Next.Next.Next.Next.Next.Next.Val)
	assert.Equal(t, 80, afterMerge.Next.Next.Next.Next.Next.Next.Next.Val)
	assert.Equal(t, 82, afterMerge.Next.Next.Next.Next.Next.Next.Next.Next.Val)
	assert.Equal(t, 83, afterMerge.Next.Next.Next.Next.Next.Next.Next.Next.Next.Val)
	assert.Equal(t, 84, afterMerge.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Val)
}

func TestReciprocalKNode(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	head := GeneralLinkList(testSlice)

	res, _ := ReciprocalKNode(head, 1)
	assert.Equal(t, 5, res)

	res, _ = ReciprocalKNode(head, 2)
	assert.Equal(t, 4, res)

	res, _ = ReciprocalKNode(head, 3)
	assert.Equal(t, 3, res)

	res, _ = ReciprocalKNode(head, 4)
	assert.Equal(t, 2, res)

	res, _ = ReciprocalKNode(head, 5)
	assert.Equal(t, 1, res)

}

func TestEntryNodeOfLoop(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	head := GeneralLinkList(testSlice)

	res := EntryNodeOfLoop(head)
	assert.Equal(t, true, res == nil)

	head.Next.Next.Next.Next.Next = head
	res = EntryNodeOfLoop(head)
	assert.Equal(t, head, res)

	head.Next.Next.Next.Next.Next = head.Next.Next
	res = EntryNodeOfLoop(head)
	assert.Equal(t, 3, res.Val)
}

func TestFindFirstCommonNode(t *testing.T) {
	leftSlice := []int{1, 2, 3, 4, 5}
	leftHead := GeneralLinkList(leftSlice)

	rightSlice := []int{1, 2, 3, 4, 5}
	rightHead := GeneralLinkList(rightSlice)

	res := FindFirstCommonNode(leftHead, rightHead)
	assert.Equal(t, true, res == nil)

	newNode := GeneralListNode(233)
	leftHead.Next.Next.Next.Next.Next = newNode
	rightHead.Next.Next.Next.Next.Next = newNode
	res = FindFirstCommonNode(leftHead, rightHead)
	assert.Equal(t, 233, res.Val)

	leftHead.Next.Next.Next.Next.Next = rightHead.Next.Next
	res = FindFirstCommonNode(leftHead, rightHead)
	assert.Equal(t, 3, res.Val)
}

func TestPrint(t *testing.T) {
	testSlice := []int{8, 6, 10, 5, 7, 9, 11}
	root := GeneralFullBinaryTree(testSlice)
	res := root.Print()

	assert.Equal(t, 8, res[0][0])

	assert.Equal(t, 6, res[1][0])
	assert.Equal(t, 10, res[1][1])

	assert.Equal(t, 5, res[2][0])
	assert.Equal(t, 7, res[2][1])
	assert.Equal(t, 9, res[2][2])
	assert.Equal(t, 11, res[2][3])
}

func TestKthNode(t *testing.T) {
	testSlice := []int{5, 3, 7, 2, 4, 6, 8}
	bst := NewBSTree(testSlice)

	node := KthNode(bst, 3)
	assert.Equal(t, 4, node.Val)
}
