package ringbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingBuffer(t *testing.T) {
	buffer := NewRingBuffer[int]()

	buffer.Push(5)

	val, ok := buffer.Pop()
	assert.Equal(t, val, 5, "Value should be 5, but is %d", val)
	assert.True(t, ok, "Pop should return true, but returned false")

	val, ok = buffer.Pop()
	assert.Equal(t, val, 0, "Value should be 0, but is %d", val)
	assert.False(t, ok, "Pop should return false, but returned true")

	buffer.Push(42)
	buffer.Push(9)

	val, ok = buffer.Pop()
	assert.Equal(t, val, 42, "Value should be 42, but is %d", val)
	assert.True(t, ok, "Pop should return true, but returned false")

	val, ok = buffer.Pop()
	assert.Equal(t, val, 9, "Value should be 9, but is %d", val)
	assert.True(t, ok, "Pop should return true, but returned false")

	val, ok = buffer.Pop()
	assert.Equal(t, val, 0, "Value should be 0, but is %d", val)
	assert.False(t, ok, "Pop should return false, but returned true")

	buffer.Push(42)
	buffer.Push(9)
	buffer.Push(12)

	val, ok = buffer.Get(2)
	assert.Equal(t, val, 12, "Value should be 12, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	val, ok = buffer.Get(1)
	assert.Equal(t, val, 9, "Value should be 9, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	val, ok = buffer.Get(0)
	assert.Equal(t, val, 42, "Value should be 42, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")
}
