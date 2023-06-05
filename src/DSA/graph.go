package dsa

type GraphEdge struct {
	To     int
	Weight int
}

type (
	WeightedAdjacencyList   [][]GraphEdge
	WeightedAdjacencyMatrix [][]int
)

var AdjList1 WeightedAdjacencyList = WeightedAdjacencyList{
	{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	},
	{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	},
	{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	},
	{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	},
	{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	},
	{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	},
	{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	},
}

var AdjList2 WeightedAdjacencyList = WeightedAdjacencyList{
	{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	},
	{
		{To: 4, Weight: 1},
	},
	{
		{To: 3, Weight: 7},
	},
	{},
	{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	},
	{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
	},
	{
		{To: 3, Weight: 1},
	},
}

var AdjMatrix1 WeightedAdjacencyMatrix = WeightedAdjacencyMatrix{
	{0, 3, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0},
	{0, 0, 7, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 5, 0, 2, 0},
	{0, 0, 18, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 1},
}
