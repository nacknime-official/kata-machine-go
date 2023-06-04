package trie

type Trie struct{}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(item string) {}

func (t *Trie) Delete(item string) {}

func (t *Trie) Find(partial string) []string { return []string{} }
