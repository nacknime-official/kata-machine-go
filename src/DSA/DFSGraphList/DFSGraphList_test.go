package dfsgraphlist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestDfs(t *testing.T) {
	list := dsa.AdjList2

	expected := []int{0, 1, 4, 5, 6}
	result := Dfs(list, 0, 6)
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)

	expected = []int{}
	result = Dfs(list, 6, 0)
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
