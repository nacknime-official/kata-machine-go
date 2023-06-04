package dfsgraphlist

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge

func Dfs(graph WeightedAdjacencyList, source int, needle int) []int {}
