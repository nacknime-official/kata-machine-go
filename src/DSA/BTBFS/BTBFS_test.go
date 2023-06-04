package btbfs

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestBfs(t *testing.T) {
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

	result := Bfs(tree, 45)
	if !result {
		t.Errorf("Expected result: true, got %v", result)
	}

	result = Bfs(tree, 7)
	if !result {
		t.Errorf("Expected result: true, got %v", result)
	}

	result = Bfs(tree, 69)
	if result {
		t.Errorf("Expected result: false, got %v", result)
	}
}
