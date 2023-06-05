package trie

import (
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
	expected := []string{
		"foo",
		"fool",
		"foolish",
	}
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)

	trie.Delete("fool")
	expected = []string{
		"foo",
		"fool",
	}
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
}
