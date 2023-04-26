package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()

	s.Push(5)
	s.Push(7)
	s.Push(9)

	val, ok := s.Pop()
	assert.True(t, ok, "Pop should return true, but returned false")
	assert.Equal(t, 9, val, "Value should be 9, but is %d", val)
	assert.Equal(t, 2, s.Length, "Length should be 2, but is %d", s.Length)

	s.Push(11)
	val, ok = s.Pop()
	assert.True(t, ok, "Pop should return true, but returned false")
	assert.Equal(t, 11, val, "Value should be 11, but is %d", val)

	val, ok = s.Pop()
	assert.True(t, ok, "Pop should return true, but returned false")
	assert.Equal(t, 7, val, "Value should be 7, but is %d", val)

	val, ok = s.Peek()
	assert.True(t, ok, "Peek should return true, but returned false")
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)

	val, ok = s.Pop()
	assert.True(t, ok, "Pop should return true, but returned false")
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)

	val, ok = s.Pop()
	assert.False(t, ok, "Pop should return false, but returned true")
	assert.Zero(t, val, "Value should be 0, but is %d", val)
	assert.Equal(t, 0, s.Length, "Length should be 0, but is %d", s.Length)

	// just wanted to make sure that I could not blow up myself when i remove
	// everything
	s.Push(69)
	val, ok = s.Peek()
	assert.True(t, ok, "Peek should return true, but returned false")
	assert.Equal(t, 69, val, "Value should be 69, but is %d", val)
	assert.Equal(t, 1, s.Length, "Length should be 1, but is %d", s.Length)

	//yayaya
}
