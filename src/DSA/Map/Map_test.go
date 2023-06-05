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

	result := mapp.Get("bar")
	assert.Equal(t, 69, result, "Expected: 69, got: %v", result)

	result = mapp.Get("blaz")
	assert.Nil(t, result, "Expected: nil, got: %v", result)

	mapp.Delete("barblarbr")
	assert.Equal(t, 4, mapp.Len(), "Expected: 4, got: %v", mapp.Len())

	mapp.Delete("bar")
	assert.Equal(t, 3, mapp.Len(), "Expected: 3, got: %v", mapp.Len())

	result = mapp.Get("bar")
	assert.Nil(t, result, "Expected: nil, got: %v", result)
}
