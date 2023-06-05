package btinorder

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestInOrderSearch(t *testing.T) {
	tree := dsa.Tree1

	result := InOrderSearch(tree)
	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}

	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
