package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	list := NewQueue[int]()

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	val, ok := list.Deque()
	assert.True(t, ok, "Deque should return true, but returned false")
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)
	assert.Equal(t, 2, list.Length, "Length should be 2, but is %d", list.Length)

	list.Enqueue(11)

	val, ok = list.Deque()
	assert.True(t, ok, "Deque should return true, but returned false")
	assert.Equal(t, 7, val, "Value should be 7, but is %d", val)

	val, ok = list.Deque()
	assert.True(t, ok, "Deque should return true, but returned false")
	assert.Equal(t, 9, val, "Value should be 9, but is %d", val)

	val, ok = list.Peek()
	assert.True(t, ok, "Peek should return true, but returned false")
	assert.Equal(t, 11, val, "Value should be 11, but is %d", val)

	val, ok = list.Deque()
	assert.True(t, ok, "Deque should return true, but returned false")
	assert.Equal(t, 11, val, "Value should be 11, but is %d", val)

	val, ok = list.Deque()
	assert.False(t, ok, "Deque should return false, but returned true")
	assert.Zero(t, val, "Value should be 0, but is %d", val)
	assert.Equal(t, 0, list.Length, "Length should be 0, but is %d", list.Length)

	list.Enqueue(69)

	val, ok = list.Peek()
	assert.True(t, ok, "Peek should return true, but returned false")
	assert.Equal(t, 69, val, "Value should be 69, but is %d", val)
	assert.Equal(t, 1, list.Length, "Length should be 1, but is %d", list.Length)
}
