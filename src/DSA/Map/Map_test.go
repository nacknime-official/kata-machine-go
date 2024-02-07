package mapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	mapp := NewMap[string, int]()

	mapp.Set("foo", 55)
	assert.Equal(t, 1, mapp.Len(), "Expected: 1, got: %v", mapp.Len())

	mapp.Set("fool", 75)
	assert.Equal(t, 2, mapp.Len(), "Expected: 2, got: %v", mapp.Len())

	mapp.Set("foolish", 105)
	assert.Equal(t, 3, mapp.Len(), "Expected: 3, got: %v", mapp.Len())

	mapp.Set("bar", 69)
	assert.Equal(t, 4, mapp.Len(), "Expected: 4, got: %v", mapp.Len())

	result, ok := mapp.Get("bar")
	assert.True(t, ok, "Get should return true, but returned false")
	assert.Equal(t, 69, result, "Expected: 69, got: %v", result)

	result, ok = mapp.Get("blaz")
	assert.False(t, ok, "Get should return false, but returned true")
	assert.Zero(t, result, "Value should be 0, but is %v", result)

	ok = mapp.Delete("barblarbr")
	assert.False(t, ok, "Delete should return false, but returned true")
	assert.Equal(t, 4, mapp.Len(), "Expected: 4, got: %v", mapp.Len())

	mapp.Delete("bar")
	assert.True(t, ok, "Delete should return true, but returned false")
	assert.Equal(t, 3, mapp.Len(), "Expected: 3, got: %v", mapp.Len())

	result, ok = mapp.Get("bar")
	assert.False(t, ok, "Get should return false, but returned true")
	assert.Zero(t, result, "Value should be 0, but is %v", result)
}
