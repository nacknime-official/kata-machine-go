package dfsonbst

type BinaryNode struct {
	Value int
	Right *BinaryNode
	Left  *BinaryNode
}

func Dfs(head *BinaryNode, needle int) bool {}
