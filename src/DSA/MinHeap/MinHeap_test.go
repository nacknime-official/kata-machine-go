package minheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()

	assert.Equal(t, heap.Length, 0)

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	assert.Equal(t, heap.Length, 8)
	assert.Equal(t, heap.Delete(), 1)
	assert.Equal(t, heap.Delete(), 3)
	assert.Equal(t, heap.Delete(), 4)
	assert.Equal(t, heap.Delete(), 5)
	assert.Equal(t, heap.Length, 4)
	assert.Equal(t, heap.Delete(), 7)
	assert.Equal(t, heap.Delete(), 8)
	assert.Equal(t, heap.Delete(), 69)
	assert.Equal(t, heap.Delete(), 420)
	assert.Equal(t, heap.Length, 0)
}
