package linearsearchlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	assert.Equal(t, LinearSearch(arr, 69), true, "Value 69 should be in the array")
	assert.Equal(t, LinearSearch(arr, 1336), false, "Value 1336 should not be in the array")
	assert.Equal(t, LinearSearch(arr, 69420), true, "Value 69420 should be in the array")
	assert.Equal(t, LinearSearch(arr, 69421), false, "Value 69421 should not be in the array")
	assert.Equal(t, LinearSearch(arr, 1), true, "Value 1 should be in the array")
	assert.Equal(t, LinearSearch(arr, 0), false, "Value 0 should not be in the array")
}
