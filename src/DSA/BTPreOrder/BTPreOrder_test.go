package btpreorder

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestPreOrder(t *testing.T) {
	tree := &dsa.BinaryNode[int]{
		Value: 20,
		Right: &dsa.BinaryNode[int]{
			Value: 50,
			Right: &dsa.BinaryNode[int]{
				Value: 100,
				Right: nil,
				Left:  nil,
			},
			Left: &dsa.BinaryNode[int]{
				Value: 30,
				Right: &dsa.BinaryNode[int]{
					Value: 45,
					Right: nil,
					Left:  nil,
				},
				Left: &dsa.BinaryNode[int]{
					Value: 29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &dsa.BinaryNode[int]{
			Value: 10,
			Right: &dsa.BinaryNode[int]{
				Value: 15,
				Right: nil,
				Left:  nil,
			},
			Left: &dsa.BinaryNode[int]{
				Value: 5,
				Right: &dsa.BinaryNode[int]{
					Value: 7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}

	expected := []int{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}

	result := PreOrderSearch(tree)

	if !dsa.ArraysEqual(expected, result) {
		t.Errorf("Expected result: %v, got: %v", expected, result)
	}
}
