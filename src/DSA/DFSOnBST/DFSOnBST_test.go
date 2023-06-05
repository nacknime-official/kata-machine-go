package dfsonbst

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestDfs(t *testing.T) {
	tree := dsa.Tree1

	result := Dfs(tree, 45)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = Dfs(tree, 7)
	assert.True(t, result, "Expected: true, got: %v", result)

	result = Dfs(tree, 69)
	assert.False(t, result, "Expected: false, got: %v", result)
}
