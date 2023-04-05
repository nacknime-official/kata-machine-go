package bubblesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	BubbleSort(arr)
	assert.Equal(t, []int{3, 4, 7, 9, 42, 69, 420}, arr)
}
