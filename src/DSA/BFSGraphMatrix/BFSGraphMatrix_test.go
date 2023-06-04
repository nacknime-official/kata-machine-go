package bfsgraphmatrix

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestBfs(t *testing.T) {
	matrix2 := dsa.WeightedAdjacencyMatrix{
		{0, 3, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}

	expected := []int{
		0,
		1,
		4,
		5,
		6,
	}

	result := Bfs(matrix2, 0, 6)

	if !dsa.ArraysEqual(expected, result) {
		t.Errorf("Expected result: %v, got: %v", expected, result)
	}

	result = Bfs(matrix2, 6, 0)
	if result != nil {
		t.Errorf("Expected result: nil, got: %v", result)
	}
}
