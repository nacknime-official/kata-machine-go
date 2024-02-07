package primslist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/assert"
)

func TestPrims(t *testing.T) {
	list := dsa.AdjList1

	expected := dsa.WeightedAdjacencyList{
		{
			{To: 2, Weight: 1},
			{To: 1, Weight: 3},
		},
		{
			{To: 0, Weight: 3},
			{To: 4, Weight: 1},
		},
		{{To: 0, Weight: 1}},
		{{To: 6, Weight: 1}},
		{
			{To: 1, Weight: 1},
			{To: 5, Weight: 2},
		},
		{
			{To: 4, Weight: 2},
			{To: 6, Weight: 1},
		},
		{
			{To: 5, Weight: 1},
			{To: 3, Weight: 1},
		},
	}

	result := Prims(list)

	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
