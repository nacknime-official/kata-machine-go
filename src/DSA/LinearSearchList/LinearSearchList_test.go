package linearsearchlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	assert.Equal(t, LinearSearch(arr, 69), true)
	assert.Equal(t, LinearSearch(arr, 1336), false)
	assert.Equal(t, LinearSearch(arr, 69420), true)
	assert.Equal(t, LinearSearch(arr, 69421), false)
	assert.Equal(t, LinearSearch(arr, 1), true)
	assert.Equal(t, LinearSearch(arr, 0), false)
}
