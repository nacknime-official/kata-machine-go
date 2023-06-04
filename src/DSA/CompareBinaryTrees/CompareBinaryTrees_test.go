package comparebinarytrees

import "testing"

func TestCompare(t *testing.T) {
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

	tree2 := &BinaryNode{
		Value: 20,
		Right: &BinaryNode{
			Value: 50,
			Right: nil,
			Left: &BinaryNode{
				Value: 30,
				Right: &BinaryNode{
					Value: 45,
					Right: &BinaryNode{
						Value: 49,
						Left:  nil,
						Right: nil,
					},
					Left: nil,
				},
				Left: &BinaryNode{
					Value: 29,
					Right: nil,
					Left: &BinaryNode{
						Value: 21,
						Right: nil,
						Left:  nil,
					},
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

	result := Compare(tree, tree)
	if !result {
		t.Errorf("Expected result: true, got %v", result)
	}

	result = Compare(tree, tree2)
	if result {
		t.Errorf("Expected result: false, got %v", result)
	}
}
