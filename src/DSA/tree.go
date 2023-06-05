package dsa

type BinaryNode[T comparable] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

var Tree1 BinaryNode[int] = BinaryNode[int]{
	Value: 20,
	Right: &BinaryNode[int]{
		Value: 50,
		Right: &BinaryNode[int]{
			Value: 100,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryNode[int]{
			Value: 30,
			Right: &BinaryNode[int]{
				Value: 45,
				Right: nil,
				Left:  nil,
			},
			Left: &BinaryNode[int]{
				Value: 29,
				Right: nil,
				Left:  nil,
			},
		},
	},
	Left: &BinaryNode[int]{
		Value: 10,
		Right: &BinaryNode[int]{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryNode[int]{
			Value: 5,
			Right: &BinaryNode[int]{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}

var Tree2 BinaryNode[int] = BinaryNode[int]{
	Value: 20,
	Right: &BinaryNode[int]{
		Value: 50,
		Right: nil,
		Left: &BinaryNode[int]{
			Value: 30,
			Right: &BinaryNode[int]{
				Value: 45,
				Right: &BinaryNode[int]{
					Value: 49,
					Left:  nil,
					Right: nil,
				},
				Left: nil,
			},
			Left: &BinaryNode[int]{
				Value: 29,
				Right: nil,
				Left: &BinaryNode[int]{
					Value: 21,
					Right: nil,
					Left:  nil,
				},
			},
		},
	},
	Left: &BinaryNode[int]{
		Value: 10,
		Right: &BinaryNode[int]{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryNode[int]{
			Value: 5,
			Right: &BinaryNode[int]{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}
