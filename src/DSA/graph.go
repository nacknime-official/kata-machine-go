package dsa

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge

type WeightedAdjacencyMatrix [][]int

type Point struct {
	X int
	Y int
}

type BinaryNode[T comparable] struct {
	Value T
	Right *BinaryNode[T]
	Left  *BinaryNode[T]
}
