package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[string, int](3)

	result := lru.Get("foo")
	assert.Nil(t, result, "Expected: nil, got: %v", result)

	lru.Update("foo", 69)
	result = lru.Get("foo")
	assert.Equal(t, 69, result, "Expected: 69, got: %v", result)

	lru.Update("bar", 420)
	result = lru.Get("bar")
	assert.Equal(t, 420, result, "Expected: 420, got: %v", result)

	lru.Update("baz", 1337)
	result = lru.Get("baz")
	assert.Equal(t, 1337, result, "Expected: 1337, got: %v", result)

	lru.Update("ball", 69420)
	result = lru.Get("ball")
	assert.Equal(t, 69420, result, "Expected: 69420, got: %v", result)

	result = lru.Get("foo")
	assert.Nil(t, result, "Expected: nil, got: %v", result)

	result = lru.Get("bar")
	assert.Equal(t, 420, result, "Expected: 420, got: %v", result)

	lru.Update("foo", 69)
	result = lru.Get("bar")
	assert.Equal(t, 420, result, "Expected: 420, got: %v", result)

	result = lru.Get("foo")
	assert.Equal(t, 69, result, "Expected: 69, got: %v", result)

	result = lru.Get("baz")
	assert.Nil(t, result, "Expected: nil, got: %v", result)
}
