package trie

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foo")
	trie.Insert("fool")
	trie.Insert("foolish")
	trie.Insert("bar")

	result := trie.Find("fo")
	sort.Strings(result)
	// If bumping to go 1.21
	// slices.Sort(result)

	expected := []string{
		"foo",
		"fool",
		"foolish",
	}
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)

	trie.Delete("fool")
	result = trie.Find("fo")
	sort.Strings(result)
	// If bumping to go 1.21
	// slices.Sort(result)

	expected = []string{
		"foo",
		"fool",
	}
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
