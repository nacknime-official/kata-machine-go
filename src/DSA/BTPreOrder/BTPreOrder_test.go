package btpreorder

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestInOrderSearch(t *testing.T) {
	tree := dsa.Tree1

	result := PreOrderSearch(tree)
	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
