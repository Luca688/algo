package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
