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
	assert.True(t, ok)
	assert.Equal(t, 9, val)
	assert.Equal(t, 2, s.Length)

	s.Push(11)
	val, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 11, val)

	val, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 7, val)

	val, ok = s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 5, val)

	val, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 5, val)

	val, ok = s.Pop()
	assert.False(t, ok)
	assert.Zero(t, val)
	assert.Equal(t, 0, s.Length)

	// just wanted to make sure that I could not blow up myself when i remove
	// everything
	s.Push(69)
	val, ok = s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 69, val)
	assert.Equal(t, 1, s.Length)

	//yayaya
}
