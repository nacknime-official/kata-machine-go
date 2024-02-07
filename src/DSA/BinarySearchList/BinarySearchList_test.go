package binarysearchlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	result := BinarySearch(foo, 69)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = BinarySearch(foo, 1336)
	assert.False(t, result, "Expected: false, got: %v", result)

	result = BinarySearch(foo, 69420)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = BinarySearch(foo, 69421)
	assert.False(t, result, "Expected: false, got: %v", result)

	result = BinarySearch(foo, 1)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = BinarySearch(foo, 0)
	assert.False(t, result, "Expected: false, got: %v", result)
}
