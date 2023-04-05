package ringbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingBuffer(t *testing.T) {
	buffer := NewRingBuffer[int]()

	buffer.Push(5)

	val, ok := buffer.Pop()
	assert.Equal(t, val, 5)
	assert.True(t, ok)

	val, ok = buffer.Pop()
	assert.Equal(t, val, 0)
	assert.False(t, ok)

	buffer.Push(42)
	buffer.Push(9)

	val, ok = buffer.Pop()
	assert.Equal(t, val, 42)
	assert.True(t, ok)

	val, ok = buffer.Pop()
	assert.Equal(t, val, 9)
	assert.True(t, ok)

	val, ok = buffer.Pop()
	assert.Equal(t, val, 0)
	assert.False(t, ok)

	buffer.Push(42)
	buffer.Push(9)
	buffer.Push(12)

	val, ok = buffer.Get(2)
	assert.Equal(t, val, 12)
	assert.True(t, ok)

	val, ok = buffer.Get(1)
	assert.Equal(t, val, 9)
	assert.True(t, ok)

	val, ok = buffer.Get(0)
	assert.Equal(t, val, 42)
	assert.True(t, ok)
}
