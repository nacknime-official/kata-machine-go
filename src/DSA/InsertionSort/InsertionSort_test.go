package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	InsertionSort(arr)

	expected := []int{3, 4, 7, 9, 42, 69, 420}
	assert.Equal(t, expected, arr, "Expected: %v, got: %v", expected, arr)
}
