package dfsonbst

import "testing"

func TestDfs(t *testing.T) {
	tree := &BinaryNode{
		Value: 20,
		Right: &BinaryNode{
			Value: 50,
			Right: &BinaryNode{
				Value: 100,
				Right: nil,
				Left:  nil,
			},
			Left: &BinaryNode{
				Value: 30,
				Right: &BinaryNode{
					Value: 45,
					Right: nil,
					Left:  nil,
				},
				Left: &BinaryNode{
					Value: 29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &BinaryNode{
			Value: 10,
			Right: &BinaryNode{
				Value: 15,
				Right: nil,
				Left:  nil,
			},
			Left: &BinaryNode{
				Value: 5,
				Right: &BinaryNode{
					Value: 7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}

	result := Dfs(tree, 45)

	if !result {
		t.Errorf("Expected result: true, got: %d", result)
	}

	result = Dfs(tree, 7)

	if !result {
		t.Errorf("Expected result: true, got: %d", result)
	}

	result = Dfs(tree, 69)

	if !result {
		t.Errorf("Expected result: false, got: %d", result)
	}
}
