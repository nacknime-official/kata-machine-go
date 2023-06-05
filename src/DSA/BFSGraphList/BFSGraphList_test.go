package bfsgraphlist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestBfs(t *testing.T) {
	list := dsa.AdjList2

	result := Bfs(list, 6, 0)
	expected := []int{0, 1, 4, 5, 6}
	assert.Equal(t, expected, result, "Array should be: %v\nBut got: %v", expected, result)

	result = Bfs(list, 0, 6)
	expected = []int{}
	assert.Equal(t, expected, result, "Array should be: %v\nBut got: %v", expected, result)
}
