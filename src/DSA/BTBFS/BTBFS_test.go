package btbfs

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestBfs(t *testing.T) {
	tree := dsa.Tree1

	result := Bfs(tree, 45)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = Bfs(tree, 7)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = Bfs(tree, 69)
	assert.False(t, result, "Expected: false, got: %v", result)
}
