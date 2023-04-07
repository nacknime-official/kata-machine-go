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
	assert.True(t, ok)
	assert.Equal(t, 5, val)
	assert.Equal(t, 2, list.Length)

	list.Enqueue(11)

	val, ok = list.Deque()
	assert.True(t, ok)
	assert.Equal(t, 7, val)

	val, ok = list.Deque()
	assert.True(t, ok)
	assert.Equal(t, 9, val)

	val, ok = list.Peek()
	assert.True(t, ok)
	assert.Equal(t, 11, val)

	val, ok = list.Deque()
	assert.True(t, ok)
	assert.Equal(t, 11, val)

	val, ok = list.Deque()
	assert.False(t, ok)
	assert.Zero(t, val)
	assert.Equal(t, 0, list.Length)

	list.Enqueue(69)

	val, ok = list.Peek()
	assert.True(t, ok)
	assert.Equal(t, 69, val)
	assert.Equal(t, 1, list.Length)
}
