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
	assert.Equal(t, 9, val, "Value should be 9, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	val, ok = list.RemoveAt(1)
	assert.Equal(t, 7, val, "Value should be 7, but is %d", val)
	assert.True(t, ok, "RemoveAt should return true, but returned false")
	assert.Equal(t, 2, list.Len(), "Length should be 2, but is %d", list.Len())

	// remove

	list.Append(11)
	val, ok = list.RemoveAt(1)
	assert.Equal(t, 9, val, "Value should be 9, but is %d", val)
	assert.True(t, ok, "RemoveAt should return true, but returned false")

	val, ok = list.Remove(9)
	assert.Equal(t, 0, val, "Value should be 0, but is %d", val)
	assert.False(t, ok, "Remove should return false, but returned true")

	val, ok = list.RemoveAt(0)
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)
	assert.True(t, ok, "RemoveAt should return true, but returned false")

	val, ok = list.RemoveAt(0)
	assert.Equal(t, 11, val, "Value should be 11, but is %d", val)
	assert.True(t, ok, "RemoveAt should return true, but returned false")
	assert.Equal(t, 0, list.Len(), "Length should be 0, but is %d", list.Len())

	// prepend

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	val, ok = list.Get(2)
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	val, ok = list.Get(0)
	assert.Equal(t, 9, val, "Value should be 9, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	val, ok = list.Remove(9)
	assert.Equal(t, 9, val, "Value should be 9, but is %d", val)
	assert.True(t, ok, "Remove should return true, but returned false")

	assert.Equal(t, 2, list.Len(), "Length should be 2, but is %d", list.Len())

	val, ok = list.Get(0)
	assert.Equal(t, 7, val, "Value should be 7, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	// insert

	list.InsertAt(10, 1)

	val, ok = list.Get(1)
	assert.Equal(t, 10, val, "Value should be 10, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	val, ok = list.Get(2)
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	list.InsertAt(20, 2)
	val, ok = list.Get(2)
	assert.Equal(t, 20, val, "Value should be 20, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")
	val, ok = list.Get(3)
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")

	list.InsertAt(30, 4)
	val, ok = list.Get(4)
	assert.Equal(t, 30, val, "Value should be 30, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")
	val, ok = list.Get(3)
	assert.Equal(t, 5, val, "Value should be 5, but is %d", val)
	assert.True(t, ok, "Get should return true, but returned false")
}
