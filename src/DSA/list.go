package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type List[T comparable] interface {
	Len() int

	Prepend(item T)
	InsertAt(item T, idx int)
	Append(item T)
	Remove(item T) (T, bool)
	Get(idx int) (T, bool)
	RemoveAt(idx int) (T, bool)
}

func TestList(t *testing.T, list List[int]) {
	t.Helper()

	// append

	list.Append(5)
	list.Append(7)
	list.Append(9)

	val, ok := list.Get(2)
	assert.Equal(t, 9, val)
	assert.True(t, ok)

	val, ok = list.RemoveAt(1)
	assert.Equal(t, 7, val)
	assert.True(t, ok)
	assert.Equal(t, 2, list.Len())

	// remove

	list.Append(11)
	val, ok = list.RemoveAt(1)
	assert.Equal(t, 9, val)
	assert.True(t, ok)

	val, ok = list.Remove(9)
	assert.Equal(t, 0, val)
	assert.False(t, ok)

	val, ok = list.RemoveAt(0)
	assert.Equal(t, 5, val)
	assert.True(t, ok)

	val, ok = list.RemoveAt(0)
	assert.Equal(t, 11, val)
	assert.True(t, ok)
	assert.Equal(t, 0, list.Len())

	// prepend

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	val, ok = list.Get(2)
	assert.Equal(t, 5, val)
	assert.True(t, ok)

	val, ok = list.Get(0)
	assert.Equal(t, 9, val)
	assert.True(t, ok)

	val, ok = list.Remove(9)
	assert.Equal(t, 9, val)
	assert.True(t, ok)

	assert.Equal(t, 2, list.Len())

	val, ok = list.Get(0)
	assert.Equal(t, 7, val)
	assert.True(t, ok)

	// insert

	list.InsertAt(10, 1)

	val, ok = list.Get(1)
	assert.Equal(t, 10, val)
	assert.True(t, ok)

	val, ok = list.Get(2)
	assert.Equal(t, 5, val)
	assert.True(t, ok)

	list.InsertAt(20, 2)
	val, ok = list.Get(2)
	assert.Equal(t, 20, val)
	assert.True(t, ok)
	val, ok = list.Get(3)
	assert.Equal(t, 5, val)
	assert.True(t, ok)

	list.InsertAt(30, 4)
	val, ok = list.Get(4)
	assert.Equal(t, 30, val)
	assert.True(t, ok)
	val, ok = list.Get(3)
	assert.Equal(t, 5, val)
	assert.True(t, ok)
}
