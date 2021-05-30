package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6, 9, 10}
	res, _ := BinarySearch(testSlice, 1)
	assert.Equal(t, true, res)

	res, _ = BinarySearch(testSlice, 10)
	assert.Equal(t, true, res)

	res, _ = BinarySearch(testSlice, 5)
	assert.Equal(t, true, res)

	res, _ = BinarySearch(testSlice, 7)
	assert.Equal(t, false, res)

}
