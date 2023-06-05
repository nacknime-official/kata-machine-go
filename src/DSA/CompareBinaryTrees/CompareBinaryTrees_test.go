package comparebinarytrees

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	tree := dsa.Tree1
	tree2 := dsa.Tree2

	result := Compare(tree, tree)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = Compare(tree, tree2)
	assert.False(t, result, "Expected: false, got: %v", result)
}
