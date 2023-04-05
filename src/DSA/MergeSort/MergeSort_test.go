package mergesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	MergeSort(arr)
	assert.Equal(t, []int{3, 4, 7, 9, 42, 69, 420}, arr)
}
