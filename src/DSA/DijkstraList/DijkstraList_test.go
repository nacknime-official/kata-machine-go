package dijkstralist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestDijkstraList(t *testing.T) {
	list := dsa.AdjList1

	expected := []int{0, 1, 4, 5, 6}
	result := DijkstraList(0, 6, list)
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
