package dijkstralist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestDijkstraList(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	list1 := make(dsa.WeightedAdjacencyList, 7)

	list1[0] = []dsa.GraphEdge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}
	list1[1] = []dsa.GraphEdge{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	}
	list1[2] = []dsa.GraphEdge{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	}
	list1[3] = []dsa.GraphEdge{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	}
	list1[4] = []dsa.GraphEdge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}
	list1[5] = []dsa.GraphEdge{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	}
	list1[6] = []dsa.GraphEdge{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	}

	expected := []int{
		0,
		1,
		4,
		5,
		6,
	}

	result := DijkstraList(0, 6, list1)

	if !dsa.ArraysEqual(expected, result) {
		t.Errorf("Expected result: %v, got: %v", expected, result)
	}
}
