package btpreorder

import "testing"

func TestPostOrder(t *testing.T) {
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

	expected := []int{
		7,
		5,
		15,
		10,
		29,
		45,
		30,
		100,
		50,
		20,
	}

	result := PostOrderSearch(tree)

	if !arraysEqual(expected, result) {
		t.Errorf("Expected result: %v, got: %v", expected, result)
	}
}

func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
