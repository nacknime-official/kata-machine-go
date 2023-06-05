package btpostorder

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestInOrderSearch(t *testing.T) {
	tree := dsa.Tree1

	result := PostOrderSearch(tree)
	expected := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}

	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
