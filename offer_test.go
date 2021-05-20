package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortLinkListRepeat(t *testing.T)  {

	testHeadRepeat :=  []int{1,1,1,1,1,2,3,3,3,4,4,5,6}

	repeatLink := GeneralLinkList(testHeadRepeat)

	newRepeatLink := SortLinkListRepeat(repeatLink)
	assert.NotEqual(t, newRepeatLink, repeatLink)
	assert.Equal(t, 2, newRepeatLink.Val)
	assert.Equal(t, 5, newRepeatLink.Next.Val)
	assert.Equal(t, 6, newRepeatLink.Next.Next.Val)

	testHeadRepeat =  []int{1,1,1,1,1,2,3,3,3,4,4,5,6,6}

	repeatLink = GeneralLinkList(testHeadRepeat)
	newRepeatLink = SortLinkListRepeat(repeatLink)
	assert.NotEqual(t, newRepeatLink, repeatLink)
	assert.Equal(t, 2, newRepeatLink.Val)
	assert.Equal(t, 5, newRepeatLink.Next.Val)
	assert.Equal(t, true, newRepeatLink.Next.Next == nil)

	testHeadRepeat =  []int{1,2, 3}

	repeatLink = GeneralLinkList(testHeadRepeat)
	newRepeatLink = SortLinkListRepeat(repeatLink)
	assert.Equal(t, 1, newRepeatLink.Val)
	assert.Equal(t, 2, newRepeatLink.Next.Val)
	assert.Equal(t, 3, newRepeatLink.Next.Next.Val)
}