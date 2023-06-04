package dfsgraphlist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestDfs(t *testing.T) {
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	list2 := make(dsa.WeightedAdjacencyList, 7)

	list2[0] = []dsa.GraphEdge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}

	list2[1] = []dsa.GraphEdge{
		{To: 4, Weight: 1},
	}

	list2[2] = []dsa.GraphEdge{
		{To: 3, Weight: 7},
	}

	list2[3] = []dsa.GraphEdge{}

	list2[4] = []dsa.GraphEdge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}

	list2[5] = []dsa.GraphEdge{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
	}

	list2[6] = []dsa.GraphEdge{
		{To: 3, Weight: 1},
	}

	expected := []int{
		0,
		1,
		4,
		5,
		6,
	}

	result := Dfs(list2, 0, 6)
	if !dsa.ArraysEqual(result, expected) {
		t.Errorf("Expected result: %v, got %v", expected, result)
	}

	result = Dfs(list2, 6, 0)
	if result != nil {
		t.Errorf("Expected result: nil, got: %v", result)
	}
}
